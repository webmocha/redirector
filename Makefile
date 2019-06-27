
build:
	go build -o bin/redirector .

build-linux64:
	env GOOS=linux GOARCH=amd64 go build -o bin/redirector .

run:
	./bin/redirector

dev: build run

