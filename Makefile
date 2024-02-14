build:
	go build cmd/service/app.go
run:
	go run cmd/service/app.go
compile:
	echo "Compiling for every OS and Platform"
	GOOS=linux GOARCH=arm go build -o bin/main-linux-arm cmd/service/app.go
	GOOS=linux GOARCH=arm64 go build -o bin/main-linux-arm64 cmd/service/app.go
	GOOS=freebsd GOARCH=386 go build -o bin/main-freebsd-386 cmd/service/app.go
generate-mock:
	go generate -x ./...
get-generator:
	go install github.com/golang/mock/mockgen
test:
	go test ./...
test-conver:
	go test -cover ./...