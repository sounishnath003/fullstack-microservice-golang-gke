AppName=authservice-$(shell date +"%Y%m%d")

.PHONY: install
install:
	go mod tidy
	go mod download
	go mod verify
	go run cmd/*.go

.PHONY: build
build:
	echo "building the auth service"
	go build -ldflags "-s -w" -o ./bin/$(AppName) cmd/*.go

.PHONY: run
run: build
	./bin/$(AppName)