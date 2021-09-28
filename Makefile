dev:
	nodemon --exec "go fmt && go build -o bin/bare" main.go

build:
	go build -o bin/bare main.go

release:
	@GOOS=linux go build -o bin/linux/bare main.go
	@GOOS=darwin go build -o bin/darwin/bare main.go
	@GOOS=windows go build -o bin/windows/bare.exe main.go
	# tar -czvf bare.tar.gz ./bin
	# zip bare.zip ./bin

nightly:
	go build -o bin/baren.go main.go
