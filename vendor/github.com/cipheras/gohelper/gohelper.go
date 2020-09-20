package gohelper

import (
	"fmt"
	"log"
	"os"
)

/*
RESET
RED
GREEN
YELLOW
BLUE
PURPLE
CYAN
WHITE
BGBLACK
BOLD
UNDERLINE
BLINK
*/
const (
	// ANSI color codes
	RESET     = "\033[0m"
	RED       = "\033[31m"
	GREEN     = "\033[32m"
	YELLOW    = "\033[33m"
	BLUE      = "\033[34m"
	PURPLE    = "\033[35m"
	CYAN      = "\033[36m"
	WHITE     = "\033[37m"
	BGBLACK   = "\033[40m"
	BOLD      = "\033[1m"
	UNDERLINE = "\033[4m"
	BLINK     = "\033[5m"
	CLEAR     = "\033[2J\033[H"
)

const (
	// N :
	N = "normal"
	// E :
	E = "error"
	// W :
	W = "warning"
	// T :
	T = "text"
	// I :
	I = "info"
	// S :
	S = "shell"
)

// Try ...
func Try(msg string, err error, mode bool) { //msg,err,exit/noexit
	if err != nil {
		if msg != "" {
			if mode == true {
				Cprint(E, msg)
				log.Println(msg)
				os.Exit(0)
			}
			Cprint(W, msg)
			log.Println(msg)
			return
		}
		if mode == true {
			Cprint(E, err)
			log.Println(err)
			os.Exit(0)
		}
		Cprint(W, err)
		log.Println(err)
	}
}

// Cprint ...
func Cprint(mode string, msg ...interface{}) {
	var msgs string
	for _, v := range msg {
		msgs = msgs + fmt.Sprintf("%v ", v)
	}
	switch mode {
	case N: //normal
		fmt.Println("\n" + CYAN + "[" + GREEN + "+" + CYAN + "] " + GREEN + msgs + RESET)
	case E: //error
		fmt.Println("\n" + CYAN + "[" + RED + BLINK + "-" + RESET + CYAN + "] " + RED + BGBLACK + BOLD + "ERROR" + RESET + " " + RED + msgs + RESET)
	case W: //warning
		fmt.Println("\n" + CYAN + "[" + YELLOW + BLINK + "!" + RESET + CYAN + "] " + YELLOW + BGBLACK + BOLD + "WARN" + RESET + " " + YELLOW + msgs + RESET)
	case T: //text
		fmt.Println("\n" + CYAN + "[" + PURPLE + "*" + CYAN + "] " + PURPLE + msgs + RESET)
	case I: //info
		fmt.Println("\n" + CYAN + "[" + BLUE + "i" + CYAN + "] " + BLUE + msgs + RESET)
	case S: //shell
		fmt.Print("\n" + CYAN + "[" + PURPLE + "*" + CYAN + "] " + PURPLE + msgs + "\n" + GREEN + ">> " + RESET)
	}
}
