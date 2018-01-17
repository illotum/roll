PACKAGE := roll
PROJECT := github.com/illotum/${PACKAGE}
VERSION := $(shell git describe --tags)
LDFLAGS := -ldflags "-w -s -X main.version=${VERSION}"
BINARY := ${PACKAGE}-${VERSION}
EXEC := docker run --rm -v "$(shell pwd)":/go/src/${PROJECT} -w /go/src/${PROJECT} golang:1.9 sh -c

.PHONY: clean install version test bench

install:
	@go generate ./...
	@go install ${LDFLAGS} ./cmd/roll

clean:
	@rm -rf out

release: out/${BINARY}.linux.gz out/${BINARY}.darwin.gz out/${BINARY}.exe.gz

out/${BINARY}.linux.gz: out/${BINARY}.linux
	@gzip -k out/${BINARY}.linux

out/${BINARY}.darwin.gz: out/${BINARY}.darwin
	@gzip -k out/${BINARY}.darwin

out/${BINARY}.exe.gz: out/${BINARY}.exe 
	@gzip -k out/${BINARY}.exe

out/${BINARY}.linux:
	@mkdir -p ./out
	@go generate ./...
	@GOOS=linux GOARCH=amd64 go build ${LDFLAGS} -o out/${BINARY}.linux ./cmd/roll 

out/${BINARY}.darwin:
	@mkdir -p ./out
	@go generate ./...
	@GOOS=darwin GOARCH=amd64 go build ${LDFLAGS} -o out/${BINARY}.darwin ./cmd/roll 

out/${BINARY}.exe:
	@mkdir -p ./out
	@go generate ./...
	@GOOS=windows GOARCH=amd64 go build ${LDFLAGS} -o out/${BINARY}.exe ./cmd/roll 


version:
	@echo "${VERSION}" > version

test:
	@${EXEC} "go test -v ./..."

bench:
	@${EXEC} "go test -bench=. -benchmem ./..."
