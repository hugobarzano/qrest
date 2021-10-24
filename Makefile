
setup:
	echo "setup"
build:
	env GOOS=darwin GOARCH=amd64 go build -o ./bin/qrest.darwin server.go;
	env GOOS=linux GOARCH=amd64 go build -o ./bin/qrest server.go;
