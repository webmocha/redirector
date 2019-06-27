
build:
	go build -o bin/redirector .

build-linux64:
	env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -installsuffix 'static' -o bin/redirector .

run:
	./bin/redirector

dev: build run

