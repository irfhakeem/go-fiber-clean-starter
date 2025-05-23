tidy:
	go mod tidy
run:
	go run main.go
build:
	go build -o main main.go
test:
	go test -v ./tests | tee result.log
