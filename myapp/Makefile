BINARY_NAME=PauloApp.exe

## build: builds all binaries
build:
	@go mod vendor
	@go build -o tmp/${BINARY_NAME} .
	@echo Paulo built!

run:
	@echo Starting Paulo...
	@start /min cmd /c tmp\${BINARY_NAME} &
	@echo Paulo started!

clean:
	@echo Cleaning...
	@DEL ${BINARY_NAME}
	@go clean
	@echo Cleaned!

test:
	@echo Testing...
	@go test ./...
	@echo Done!

start: run
    
stop:
	@echo "Starting the front end..."
	@taskkill /IM ${BINARY_NAME} /F
	@echo Stopped Paulo

restart: stop start