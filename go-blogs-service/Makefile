AppName=blogsservice-$(shell date +"%Y%m%d")
DockerImageName=asia-south1-docker.pkg.dev/sounish-cloud-workstation/fullstack-microservices-golang-gke/go-blogs-service:v$(shell date +'%y%m%H.%M%S')

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


.PHONY: docker-build-push
docker-build-push:
	docker build -t $(DockerImageName) -f Dockerfile .
	docker images
	docker push $(DockerImageName)