.PHONY: build clean run vet

build:
	go build -o notes_app
	rm -rf ./static/staging
	rm -rf ./static/app
	mkdir ./static/staging
	mkdir ./static/app
	pwd
	cd ./static/staging && pwd 
	cd ./static/staging && git clone https://github.com/JacobMartinsen/notes_app_react.git .
	cd ./static/staging && pwd
	cd ./static/staging && npm install 
	cd ./static/staging && ls 
	cd ./static/staging && npm run build  
	cp -R ./static/staging/dist/ ./static/app/
	rm -rf ./static/staging/

run: build 
	./notes_app


clean:
	rm -f notes_app

vet: 
	go vet ./...

