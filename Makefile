TARGET=loctrack
PACKAGES=handler
STATUS= -e "\033[5m\033[32mDONE\033[0m \n"

.PHONY: all

all: build install clean

build:
	@echo -ne "\n[+] Creating a build..."
	@go build -o $(TARGET) -mod=vendor
	@echo $(STATUS)

clean:	
	@echo -ne "\n[+] Cleaning..."
	@go clean
	@sudo rm -rf ~/go/pkg/mod/github.com/cipheras
	@rm -f ~/go/bin/$(TARGET)
	# @rm -rf ./bin
	rm $(TARGET)
	@echo $(STATUS)

# install: build
install:
	@echo -ne "\n[+] Installing..."
	@cp $(TARGET) /usr/local/bin
	@echo $(STATUS)

uninstall: 
	@echo -ne "\n[+] Uninstalling..."
	@rm /usr/local/bin/$(TARGET)
	@echo $(STATUS)
