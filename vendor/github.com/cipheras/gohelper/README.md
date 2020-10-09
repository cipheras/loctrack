# GoHelper &nbsp; ![GitHub release (latest by date)](https://img.shields.io/github/v/release/cipheras/gohelper?style=flat&logo=superuser)

#### A GO module to help in projects in generating formatted logs in log files and colored messages on the terminal. 

![Lines of code](https://img.shields.io/tokei/lines/github/cipheras/gohelper?style=flat)
&nbsp;&nbsp;&nbsp;&nbsp;![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/cipheras/gohelper?style=flat)
&nbsp;&nbsp;&nbsp;&nbsp;![GitHub All Releases](https://img.shields.io/github/downloads/cipheras/gohelper/total?style=flat)

## Installation
You can import this module and start using it.
```
go get github.com/cipheras/gohelper
```

## Usage
**gohelper** basically has two features, creating formatted logs in a file and show formatted texts on console.

### How to create formatted logs:
***Note:** You have to configure `log` when calling log module to write logs to file instead of showing it on terminal.*

You don't have to write `if err!=nil{}` everytime, you can just do
```
Try(message, error, mode)
where, mode can be true or false
If mode=true, process will exit and if mode=false, process will generate a warning message.
```
This will write the same message to logs and also will show on terminal in that particular format.

### How to create formatted texts on console:
* To show colors on windows **cmd** also, call function `Cwindows()`.
```
Example:
err := Cwindows()
```

```
Cprint(mode, message)
where mode can be,
    N = "normal"
	E = "error"
	W = "warning"
	T = "text"
	I = "info"
	S = "shell"
```
You can also use available colors and formats inbetween any of these or your own console outputs. 
<br>Available ones are:
```
RESET | RED | GREEN | YELLOW | BLUE | PURPLE | CYAN | WHITE | BGBLACK | BOLD | UNDERLINE | BLINK | CLEAR
Example:     
fmt.Println(BLUE, "hello", BOLD, var1, BLINK, var2, "!!", RESET)
```
## To Do
- [x] Colored text on console for linux
- [x] Colored text on cmd for windows
- [ ] Add more text and background colors

## License
**gohelper** is made by **@cipheras** and is released under the terms of the &nbsp;![GitHub License](https://img.shields.io/github/license/cipheras/gohelper)

## Contact &nbsp; [![Twitter](https://img.shields.io/twitter/url?style=social&url=https%3A%2F%2Fgithub.com%2Fcipheras%2Fgohelper&label=Tweet)](https://twitter.com/intent/tweet?text=Hi:&url=https%3A%2F%2Fgithub.com%2Fcipheras%2Fgohelper)

> Feel free to submit a bug, add features or issue a pull request.

