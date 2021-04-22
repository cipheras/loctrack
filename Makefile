TARGET=loctrack
PACKAGES=handler
STATUS= "\033[5m\033[32mDONE\033[0m \n"

.PHONY: all

all: build install clean

build:
	@echo -n "\n[+] Creating a build..."
	@go build -o $(TARGET)
	@echo $(STATUS)

clean:	
	@echo -n "\n[+] Cleaning..."
	@go clean
	@go clean -modcache
	@echo $(STATUS)

# install: build
install:
	@echo -n "\n[+] Installing..."
	@cp $(TARGET) /usr/local/bin
	@echo $(STATUS)

uninstall: 
	@echo -n "\n[+] Uninstalling..."
	@rm /usr/local/bin/$(TARGET)
	@echo $(STATUS)
