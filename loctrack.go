package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
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
	subdomain = flag.String("d", "", "Subdomain (optional)")
	port      = flag.Int("p", 8080, "Port Number (optional)")
)
var url string

const (
	// VERSION ...
	VERSION = "1.4.1"
)

func main() {
	flag.Usage = func() {
		Cprint(I, "Choose options. By default a tunnel will be created itself")
		Cprint(I, "Run your own tunnel by using "+GREEN+"'-manual'"+BLUE+" flag")
		Cprint(I, "Manual TLS certificate using "+GREEN+"'-c'"+BLUE+" flag. Keep your own certs in "+GREEN+"'cert'"+BLUE+" folder")
		fmt.Println("\n" + GREEN + "##################################" + BLUE + "LocTrack" + GREEN + "##################################" + RESET)
		flag.PrintDefaults()
		fmt.Println(GREEN + "##################################" + BLUE + "LocTrack" + GREEN + "##################################\n" + RESET)
	}
	flag.Parse()
	// // Make log file
	// if _, err := os.Stat("logs"); os.IsNotExist(err) {
	// 	os.Mkdir("logs", os.ModePerm)
	// }
	f, err := os.OpenFile("log.txt", os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0644)
	Try("", err, false)
	log.SetOutput(f)
	defer f.Close()

	interrupt()
	banner()
	Cwindows() //for colors on cmd in window
	Cprint(I, "Try"+GREEN+" loctrack -h "+BLUE+"for help and other options")
	if *mantunnel {
		Cprint(T, "You have chosen manual mode. Run your own tunnel.")
	} else {
		urlCreation()
	}
	server(templateSel())
}

func banner() {
	bnr := `
	██▓     ▒█████   ▄████▄  ▄▄▄█████▓ ██▀███   ▄▄▄       ▄████▄   ██ ▄█▀
	▓██▒    ▒██▒  ██▒▒██▀ ▀█  ▓  ██▒ ▓▒▓██ ▒ ██▒▒████▄    ▒██▀ ▀█   ██▄█▒ 
	▒██░    ▒██░  ██▒▒▓█    ▄ ▒ ▓██░ ▒░▓██ ░▄█ ▒▒██  ▀█▄  ▒▓█    ▄ ▓███▄░ 
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
	fmt.Printf("/n%s%s%s", GREEN, bnr, RESET)
	fmt.Print(CYAN + "Created by:" + GREEN + crtr + RESET)
	time.Sleep(1500 * time.Millisecond)
}

func urlCreation() {
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
		log.Println("Timeout for service 1")
	} else {
		if sResp.StatusCode == 200 {
			log.Println("Service 1 Online")
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
			Try("SSH command error", err, false)
			// defer cmd.Process.Kill()
			re, err := regexp.Compile(`https:\/\/.+serveo.*\.com`)
			Try("Regexp failed", err, false)
			for i := 0; i < 8; i++ {
				url = re.FindString(stdout.String())
				if url != "" {
					Cprint(N, WHITE+"URL: "+GREEN+url)
					return
				}
				time.Sleep(1 * time.Second)
			}
			cmd.Process.Kill()
			Cprint(E, "Failed to generate URL")
			log.Println("failed to generate URL")
		} else {
			fmt.Println(RED + BGBLACK + BLINK + BOLD + "Offline" + RESET)
			log.Println("Offline...service 1 down")
		}
	}
	time.Sleep(1500 * time.Millisecond)

	// Checking localhost.run status
	fmt.Print(YELLOW + "\n" + CYAN + "[" + YELLOW + "*" + CYAN + "] " + YELLOW + "Checking service 2 status..." + RESET)
	lrResp, err := client.Get("http://localhost.run")
	// time.Sleep(1 * time.Second)
	if err != nil {
		fmt.Println(RED + BLINK + BOLD + "Offline" + RESET)
		fmt.Println(PURPLE + "Try again later...or report to the creator.\n" + RESET)
		log.Println("Timeout for service 2")
		os.Exit(0)
	}
	if lrResp.StatusCode == 200 {
		lrResp.Body.Close()
		fmt.Println(GREEN + BGBLACK + BLINK + BOLD + "Online" + RESET)
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
		Try("SSH command error", err, true)
		// defer cmd.Process.Kill()
		re, err := regexp.Compile(`http:\/\/.+localhost\.run`)
		Try("Regexp failed", err, true)
		for i := 0; i < 8; i++ {
			url = re.FindString(stdout.String())
			if url != "" {
				Cprint(N, WHITE+"URL: "+RESET+GREEN+url)
				return
			}
			time.Sleep(1 * time.Second)
		}
		cmd.Process.Kill()
		Cprint(E, "Failed to generate URL")
		log.Println("failed to generate URL")
		os.Exit(0)
	}
	fmt.Println(RED + BLINK + BOLD + "Offline" + RESET)
	log.Println("Offline...service 2 is also down")
	os.Exit(0)
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
	Try("", err, true)
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
		time.Sleep(1800 * time.Millisecond)
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
		Try("", err, true)
		return
	}
	err := http.ListenAndServe(":"+strconv.Itoa(*port), nil) //with default cert
	Try("", err, true)
}

func interrupt() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Print("\n" + CYAN + "[" + PURPLE + "*" + CYAN + "] " + PURPLE + "Aborting " + RESET)
		for i := 1; i <= 6; i++ {
			fmt.Print(PURPLE + "# " + RESET)
			time.Sleep(time.Millisecond * 200)
		}
		fmt.Print(CLEAR)
		// fmt.Print("\n\n")
		// time.Sleep(1500 * time.Millisecond)
		os.Exit(0)
	}()
}
