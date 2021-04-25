TARGET=loctrack
PACKAGES=handler
STATUS= "\033[5m\033[32mDONE\033[0m \n"

.PHONY: all

all: build install clean

build:
	@echo -n "\n[\033[5m\033[32m+\033[0m] Creating a build..."
	@go build -ldflags="-s -w" -o $(TARGET)
	@echo $(STATUS)

clean:	
	@echo -n "\n[\033[5m\033[32m+\033[0m] Cleaning..."
	@go clean
	@go clean -modcache
	@echo $(STATUS)

# install: build
install:
	@echo -n "\n[\033[5m\033[32m+\033[0m] Installing..."
	@cp $(TARGET) /usr/local/bin
	@echo $(STATUS)

uninstall: 
	@echo -n "\n[\033[5m\033[32m+\033[0m] Uninstalling..."
	@rm /usr/local/bin/$(TARGET)
	@echo $(STATUS)
