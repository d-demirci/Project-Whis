info:
	echo "Windows Agent Application"

build:
	GOOS=windows GOARCH=amd64 go build -o bin/app-amd64.exe main.go
	# GOOS=windows GOARCH=386 go build -o bin/app-windows-386.exe main.go

run:
	go run main.go