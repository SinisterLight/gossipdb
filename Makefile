.PHONY: build
.SILENT: run

SVC_ROOT = `pwd`
BINARY   = gossipdb
PACKAGE  = github.com/SinisterLight/gossipdb

all: init deps test

init:
	mkdir -p build

deps:
	glide install

test:
	go test

clean:
	if [ -f build/${BINARY} ] ; then rm build/${BINARY} ; fi


build:
	go build -v -o ${SVC_ROOT}/build/${BINARY} ${SVC_ROOT}/gossipdb.go
