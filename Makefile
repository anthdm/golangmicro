build:
	go build -o bin/micro

run: build
	./bin/micro
