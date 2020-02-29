GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

VERSION=$(shell git describe --exact-match --tags 2>/dev/null)

BUILD_DIR=build

PACKAGE_UVRDUMP=uvrdump-$(VERSION)_linux_armhf
PACKAGE_UVRINFLUX=uvrinflux-$(VERSION)_linux_armhf
PACKAGE_UVRJSON=uvrjson-$(VERSION)_linux_armhf

unexport GOPATH

all: test build
build:
	$(GOBUILD) -o $(BUILD_DIR)/$(BINARY_NAME) -i $(BUILD_SRC)

test:
	$(GOTEST) -v ./...

clean:
	$(GOCLEAN)
	rm -rf $(BUILD_DIR)

package-uvrdump: build-uvrdump
	tar -cvzf $(PACKAGE_UVRDUMP).tar.gz -C $(BUILD_DIR) $(PACKAGE_UVRDUMP)

package-uvrinflux: build-uvrinflux
	tar -cvzf $(PACKAGE_UVRINFLUX).tar.gz -C $(BUILD_DIR) $(PACKAGE_UVRINFLUX)

package-uvrjson: build-uvrjson
	tar -cvzf $(PACKAGE_UVRJSON).tar.gz -C $(BUILD_DIR) $(PACKAGE_UVRJSON)

build-uvrjson:
	GOOS=linux GOARCH=arm GOARM=6 $(GOBUILD) -o $(BUILD_DIR)/uvrjson-$(VERSION)_linux_armhf/usr/bin/uvrjson -i cmd/uvrjson/main.go
