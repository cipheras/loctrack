TARGET=loctrack
PACKAGES=handler
STATUS= -e "\033[5m\033[32mcompleted\033[0m \n"

.PHONY: all

all: build install clean

build:
	@echo -ne "\n[+] Creating a build..."
	@go build -o ./bin/$(TARGET) -mod=vendor
	@echo $(STATUS)

clean:	
	@echo -ne "\n[+] Cleaning..."
	@go clean
	@sudo rm -rf ~/go/pkg/mod/github.com/cipheras
	@rm -f ~/go/bin/$(TARGET)
	@rm -rf ./bin
	@echo $(STATUS)

install: build
	@echo -ne "\n[+] Installing..."
	@cp ./bin/$(TARGET) /usr/local/bin
	@echo $(STATUS)

uninstall: 
	@echo -ne "\n[+] Uninstalling..."
	@rm /usr/local/bin/$(TARGET)
	@echo $(STATUS)
