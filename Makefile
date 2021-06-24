build:
	go build -o bin/netcomp-cli main.go

run:
	go run main.go

test: 
	go test ./... -cover

compile:
	echo "Compiling for every OS and Platform"
	GOOS=linux GOARCH=arm go build -o bin/netcomp-cli-linux-arm main.go
	GOOS=linux GOARCH=arm64 go build -o bin/netcomp-cli-linux-arm64 main.go
	GOOS=darwin GOARCH=amd64 go build -o bin/netcomp-cli-darwin-amd64 main.go
	GOOS=windows GOARCH=amd64 go build -o bin/netcomp-cli-windows.exe main.go

all: test build