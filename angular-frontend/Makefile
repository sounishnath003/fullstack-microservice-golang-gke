DockerImageName=asia-south1-docker.pkg.dev/sounish-cloud-workstation/fullstack-microservices-golang-gke/blogs-app-frontend:v$(shell date +'%y%m%H.%M%S')

.PHONY: docker-build-push
docker-build-push:
	docker build \
	--build-arg NG_APP_AUTH_SERVICE_ENDPOINT=http://0.0.0.0:3000 \
	--build-arg NG_APP_BLOGS_SERVICE_ENDPOINT=http://0.0.0.0:3001 \
	-t $(DockerImageName) -f Dockerfile .
	docker images
	docker push $(DockerImageName)

# docker run -p 4200:80 -ti 2e7ac4b932fc