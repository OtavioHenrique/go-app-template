run:
	go run main.go

run-server:
	go run main.go run

build:
	go build -o bin/main main.go

test:
	go test ./...

build-docker:
	docker build . -t go-app-template

run-docker:
	docker run -p 8081:8081 go-app-template
