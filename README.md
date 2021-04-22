# LocTrack &nbsp; ![GitHub release (latest by date)](https://img.shields.io/github/v/release/cipheras/loctrack?style=plastic&logo=superuser)
#### A tool to locate people using social engineering. 

![Lines of code](https://img.shields.io/tokei/lines/github/cipheras/loctrack?style=plastic)
&nbsp;&nbsp;&nbsp;&nbsp;![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/cipheras/loctrack?style=plastic)
&nbsp;&nbsp;&nbsp;&nbsp;![GitHub All Releases](https://img.shields.io/github/downloads/cipheras/loctrack/total?style=plastic)

![Code Quality](https://img.shields.io/badge/dynamic/json?url=https://jsonkeeper.com/b/KNO7&label=code%20quality&query=codequality&style=plastic&labelColor=grey&color=yellowgreen)
&nbsp;&nbsp;&nbsp;&nbsp;![dependencies](https://img.shields.io/badge/dynamic/json?url=https://jsonkeeper.com/b/KNO7&label=dependencies&query=dependencies&style=plastic&labelColor=grey&color=darkgreen)
&nbsp;&nbsp;&nbsp;&nbsp;![build](https://img.shields.io/badge/dynamic/json?url=https://jsonkeeper.com/b/KNO7&label=build&query=build&style=plastic&labelColor=grey&color=darkgreen)
&nbsp;&nbsp;&nbsp;&nbsp;![platform](https://img.shields.io/badge/dynamic/json?url=https://jsonkeeper.com/b/KNO7&label=platform&query=platform&style=plastic&labelColor=grey&color=purple)


## Installation
You can either use a precompiled binary package for your architecture or you can compile **loctrack** from source.
<br>Grab the package you want from here:

Windows | Linux
--------|-------
[win-x64](https://github.com/cipheras/loctrack/releases/download/v1.4.2/loctrack-win-x64.exe) | [linux-x64](https://github.com/cipheras/loctrack/releases/download/v1.4.2/loctrack-linux-x64)

For other versions or releases go to release page.

***NOTE:** In windows, installtion is not needed. You can directly execute the **exe** file.*
*In cmd write `loctrack.exe -h`*

### Installing precompiled binary in Linux
* In order to install precompiled binary, make sure you have installed `make`.
* Download **Makefile** from [here](https://github.com/cipheras/loctrack/releases/download/v1.4.2/Makefile) and keep it and your binary in the same directory.
* Now open terminal in the same dir and run commands:

To install:
```
sudo make install
```
To uninstall:
```
sudo make uninstall
```


### Installing from source in Linux
In order to compile from source, make sure you have installed **GO** of version at least **1.15.0** (get it from [here](https://golang.org/doc/install)).

To install:
```
sudo make
```
To uninstall:
```
sudo make uninstall
```
To build:
```
make build
```


## Usage
For help type `loctrack -h`.
```
-c  For your own certificates located in cert folder
-d  Subdomain (optional)
-m  Manual Tunnel
-p  Port Number (optional) (default 8080)

```
If you want to use your own **ssl/tls** certificates put them in folder **cert** and choose option `-c`.
<br> If you want to use your own **SSH keys**, put your **ssh key** in ssh-key folder.
<br> And run command `chmod 700 ssh-key/rsa`.

## To Do
- [x] cmd color support
- [ ] More templates  

## Disclaimer
*This tools is merely a POC of what attackers can do. Author is not responsible for any use of this tool in any nefarious activity.*

## License
**loctrack** is made by **@cipheras** and is released under the terms of the &nbsp;![GitHub License](https://img.shields.io/github/license/cipheras/loctrack?color=darkgreen)

## Contact &nbsp; [![Twitter](https://img.shields.io/twitter/url?style=social&url=https%3A%2F%2Fgithub.com%2Fcipheras%2Floctrack&label=Tweet)](https://twitter.com/intent/tweet?text=Hi:&url=https%3A%2F%2Fgithub.com%2Fcipheras%2Floctrack)
> Feel free to submit a bug, add features or issue a pull request.


