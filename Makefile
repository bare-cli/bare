dev:
	nodemon --exec "go fmt && go build -o bin/test/baren" main.go

build:
	go build -o bin/bare main.go

release:
	@GOOS=linux go build -o bin/linux/bare main.go
	@GOOS=darwin go build -o bin/darwin/bare main.go
	@GOOS=windows go build -o bin/windows/bare.exe main.go

nightly:
	go build -o bin/baren main.go

setup:
	echo "Setting up"
	go mod download
