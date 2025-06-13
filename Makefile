.PHONY: build clean run vet

build:
	go build -o notes_app

run: build 
	./notes_app

clean:
	rm -f notes_app

vet: 
	go vet ./...