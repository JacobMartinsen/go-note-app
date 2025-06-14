.PHONY: build clean run vet

build: clean
	@echo "\n\033[32mBuilding the application...\033[0m"
	@echo "\n\033[32mBuilding notes_app...\033[0m"
	go build -o notes_app
	@echo "\n\033[32mCleaning build directory...\033[0m"
	rm -rf ./build
	mkdir -p ./build/staging
	mkdir ./build/app
	@echo "\n\033[32mBuilding frontend appliction...\033[0m"
	cd ./build/staging && git clone https://github.com/JacobMartinsen/notes_app_react.git .
	cd ./build/staging
	cd ./build/staging && npm install 
	cd ./build/staging && npm run build  
	@echo "\n\033[32mRemoving temporary files...\033[0m"
	cp -R ./build/staging/dist/ ./build/app/
	rm -rf ./build/staging/
	mv ./notes_app ./build/app/notes_app
	cp ./app_config.env ./build/app/

run: build
	@echo "\n\033[32mRunning the application...\033[0m"
	cd ./build/app/ && ./notes_app
	@echo "Build completed" 
clean:
	rm -rf ./build
	@echo "\n\033[32mCleaned build directory...\033[0m" 

vet: 
	go vet

