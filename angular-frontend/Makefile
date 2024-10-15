DockerImageName=asia-south1-docker.pkg.dev/sounish-cloud-workstation/fullstack-microservices-golang-gke/blogs-app-frontend:v$(shell date +'%y%m%H.%M%S')

.PHONY: docker-build-push
docker-build-push:
	docker build -t $(DockerImageName) -f Dockerfile .
	docker images
	docker push $(DockerImageName)