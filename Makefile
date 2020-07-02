TIME=$(shell date +"%Y%m%d.%H%M%S")
VERSION=0.0.1-alpha-0.2
BINARY_NAME=html2pdf

BINARY_NAME_SERVER=html2pdf


BASE_FOLDER = $(shell pwd)
BUILD_FOLDER  = $(shell pwd)/build

FLAGS_LINUX   = CGO_LDFLAGS="-L./LIB -Wl,-rpath -Wl,\$ORIGIN/LIB" CGO_ENABLED=1 GOOS=linux GOARCH=amd64  
FLAGS_DARWIN  = OSXCROSS_NO_INCLUDE_PATH_WARNINGS=1 MACOSX_DEPLOYMENT_TARGET=10.6 CC=o64-clang CXX=o64-clang++ CGO_ENABLED=0
FLAGS_FREEBSD = GOOS=freebsd GOARCH=amd64 CGO_ENABLED=1
FLAGS_WINDOWS = GOOS=windows GOARCH=amd64 CC=i686-w64-mingw32-gcc CGO_ENABLED=1 
FLAGS_ARM = CGO_ENABLED=1 GOOS=linux GOARCH=arm GOARM=7 CGO_ENABLED=0 CC=arm-linux-gnueabi-gcc

GOFLAGS_WINDOWS = -ldflags -H=windowsgui

getdeps:
	./getDeps.sh


versioning:
	./version.sh ${VERSION} ${TIME}



## Linting
lint:
	@echo "[lint] Running linter on codebase"
	@golint ./...

## Linux Build

build/html2pdf-linux: 
	${FLAGS_LINUX} go build -o ${BUILD_FOLDER}/dist/linux/bin/html2pdf .


package:
	cp install.sh $(BUILD_FOLDER)/dist/linux/
