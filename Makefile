start_dev:
	nodemon --exec go run main.go --signal SIGTERM

start:
	./lenslocked.com

build:
	go build -o ./bin .
