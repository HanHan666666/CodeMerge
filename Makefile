VERSION=$(shell git describe --tags `git rev-list --tags --max-count=1`)
BUILD_TIME=$(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
GITCOMMIT=$(shell git rev-parse HEAD)

GOVERSION=$(shell go version)

GO_FLAGS=-ldflags="-X 'main.VERSION=$(VERSION)' -X 'main.BUILDTIME=$(BUILD_TIME)' -X 'main.GITCOMMIT=$(GITCOMMIT)' -X 'main.GOVERSION=$(GOVERSION)' -extldflags -static -s -w" -trimpath


BINARY=CodeMerge
DIR=output

build:
		go build ${GO_FLAGS}

install:
		go install ${GO_FLAGS}

debug:
		go clean
		# build for linux_amd64
		GOOS=linux GOARCH=amd64 go build -o ${DIR}/${BINARY}_linux_amd64
		# build for linux_arm64
		GOOS=linux GOARCH=arm64 go build -o ${DIR}/${BINARY}_linux_arm64
		# build for windows_amd64
		GOOS=windows GOARCH=amd64 go build -o ${DIR}/${BINARY}_windows_amd64.exe
		# build for windows_arm64
		GOOS=windows GOARCH=arm64 go build -o ${DIR}/${BINARY}_windows_arm64.exe
		# build for darwin_amd64
		GOOS=darwin GOARCH=amd64 go build -o ${DIR}/${BINARY}_darwin_amd64
		# build for darwin_arm64
		GOOS=darwin GOARCH=arm64 go build -o ${DIR}/${BINARY}_darwin_arm64

release:
		go clean
		CGO_ENABLED=0

		# build for linux_amd64
		GOOS=linux GOARCH=amd64 go build ${GO_FLAGS} -o ${DIR}/${BINARY}_linux_amd64
		# build for linux_arm64
		GOOS=linux GOARCH=arm64 go build ${GO_FLAGS} -o ${DIR}/${BINARY}_linux_arm64
		# build for windows_amd64
		GOOS=windows GOARCH=amd64 go build ${GO_FLAGS} -o ${DIR}/${BINARY}_windows_amd64.exe
		# build for windows_arm64
		GOOS=windows GOARCH=arm64 go build ${GO_FLAGS} -o ${DIR}/${BINARY}_windows_arm64.exe
		# build for darwin_amd64
		GOOS=darwin GOARCH=amd64 go build ${GO_FLAGS} -o ${DIR}/${BINARY}_darwin_amd64
		# build for darwin_arm64
		GOOS=darwin GOARCH=arm64 go build ${GO_FLAGS} -o ${DIR}/${BINARY}_darwin_arm64

clean:
		go clean

.PHONY:  clean build