#!/bin/bash

## based on the postgres.docker.compose.yaml file
# it has the kubernetes deployment manifests generation command line

# -- Auth-service deployment
kubectl create deployment auth-service --image=asia-south1-docker.pkg.dev/sounish-cloud-workstation/fullstack-microservices-golang-gke/go-auth-service:v241000.3127 --port=3000 --replicas=2 --output=yaml --show-managed-fields=true --namespace=fullstack-microservice-golang --dry-run=client

# -- Create a NodePort service for auth service 
kubectl create service clusterip auth-service --tcp=3000 --output=yaml --namespace=fullstack-microservice-golang --dry-run=client

kubectl create deployment <deployment-name> --image=<image-name> --env="ENV_VAR1=value1" --env="ENV_VAR2=value2" --namespace=fullstack-microservice-golang --dry-run=client