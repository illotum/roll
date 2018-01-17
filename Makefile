PACKAGE := roll
PROJECT := github.com/illotum/${PACKAGE}
VERSION := $(shell git describe --tags)
LDFLAGS := -ldflags "-w -s -X main.version=${VERSION}"
BINARY := ${PACKAGE}-${VERSION}
EXEC := docker run --rm -v "$(shell pwd)":/go/src/${PROJECT} -w /go/src/${PROJECT} golang:1.9 sh -c

.PHONY: clean install version test bench run

install:
	@go install ${LDFLAGS}

clean:
	@rm -rf out

release: out/${BINARY}.gz

out/${BINARY}.gz: out/${BINARY}
	@gzip -k out/${BINARY}

out/${BINARY}:
	@mkdir -p ./out
	@GOOS=linux GOARCH=amd64 go build ${LDFLAGS} -o out/${BINARY}

version:
	@echo "${VERSION}" > version

test:
	@${EXEC} "go test -v ./..."

bench:
	@${EXEC} "go test -bench=. -benchmem ./..."

run:
	@${EXEC} "go install && triumph"