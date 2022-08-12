SHELL=CMD
DB= "mongodb+srv://j:rootroot@cluster0.rj0tg.mongodb.net/test?authSource=admin&replicaSet=atlas-cjefh4-shard-0&readPreference=primary&ssl=true"
ENV = DEVELOPMENT

# build: builds all the binaries
build: clean build_api
	@echo All binaries built!

# clean: cleans all the binaries and runs go clean
clean: 
	@echo Cleaning binaries...
	@echo y | DEL /S dist
	@go clean 
	@echo Cleaned and deleted binaries 

#build_api: build the api for the backend
build_api: 
	@echo Building API...
	@go build -o dist/jobby_api.exe ./cmd 
	@echo API built!!!

#start: will start to build and run
start: start_api

#start_api: starts the api
start_api: build_api 
	@echo Starting the API
	set DB=${DB}&& set ENV=${ENV}&& start /B ./dist/jobby_api.exe
	@echo API runnning in the background!!!

# stop: will stop all the api
stop: stop_api

# stop_api: will stop the api
stop_api: 
	@echo Stopping the API...
	@taskkill /IM jobby_api.exe /F 
	@echo API stopped!!