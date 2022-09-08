SHELL=/bin/sh

build:
	GOOS=windows GOARCH=amd64 go build -o handler main.go

buildMac:
	go build -o handler main.go

run: buildMac
	func start

deploy: build
	func azure functionapp publish winefamilies