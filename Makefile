GOOS=linux 
GOARCH=amd64 
CGO_ENABLED=0

build:
	go build -ldflags '-s -w' -o ./bin/imagedigest ./cmd/imagedigest/main.go

