AppName=authservice-$(shell date +"%Y%m%d")

.PHONY: install
install:
	go mod tidy
	go mod download
	go mod verify

.PHONY: build
build:
	echo "building the auth service"
	go build -ldflags "-s -w" -o ./bin/$(AppName) main.go

.PHONY: run
run: build
	export DSN="postgres://root:password@127.0.0.1:5432/auth?sslmode=disable"
	./bin/$(AppName)
