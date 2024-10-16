DockerImageName=asia-south1-docker.pkg.dev/sounish-cloud-workstation/fullstack-microservices-golang-gke/blogs-app-frontend:v$(shell date +'%y%m%H.%M%S')

.PHONY: docker-build
docker-build-push:
	docker build --platform=linux/amd64 \
	--build-arg NG_APP_AUTH_SERVICE_ENDPOINT=http://0.0.0.0:3000 \
	--build-arg NG_APP_BLOGS_SERVICE_ENDPOINT=http://0.0.0.0:3001 \
	-t $(DockerImageName) -f Dockerfile .
	docker images

# Note: ENDPOINT IPs will be the K8s Ingress LB IP of the ClusterIP services 
# Which will help in communication from External LB ---> Ingress (kube-proxy-lb) ----> (k8s cluster node) [srv1, srv2,...] 
.PHONY: docker-build-push
docker-build-push:
	docker build --platform=linux/amd64 \
	--build-arg NG_APP_AUTH_SERVICE_ENDPOINT=http://34.160.24.38:80 \
	--build-arg NG_APP_BLOGS_SERVICE_ENDPOINT=http://34.54.83.193:80 \
	-t $(DockerImageName) -f Dockerfile .
	docker images
	docker push $(DockerImageName)

# docker run -p 4200:80 -ti 2e7ac4b932fc