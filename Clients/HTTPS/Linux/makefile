compile:
	echo "Compiling for every OS and Platform"
	GOOS=freebsd GOARCH=amd64 go build -o bin/main-freebsd-386 main.go
	GOOS=linux GOARCH=amd64 go build -o bin/main-linux-386 main.go
	