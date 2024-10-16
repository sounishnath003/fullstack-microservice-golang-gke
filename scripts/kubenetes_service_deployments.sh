#!/bin/bash

## based on the postgres.docker.compose.yaml file
# it has the kubernetes deployment manifests generation command line

# -- Auth-service deployment
kubectl create deployment auth-service --image=asia-south1-docker.pkg.dev/sounish-cloud-workstation/fullstack-microservices-golang-gke/go-auth-service:v241000.3127 --port=3000 --replicas=2 --output=yaml --show-managed-fields=true --namespace=fullstack-microservice-golang --dry-run=client

# -- Create a ClusterIP service for auth service 
kubectl create service clusterip auth-service --tcp=3000 --output=yaml --namespace=fullstack-microservice-golang --dry-run=client

# -- Blogs-service deployment
kubectl create deployment blog-service --image=asia-south1-docker.pkg.dev/sounish-cloud-workstation/fullstack-microservices-golang-gke/go-blogs-service:v241015.4818 --port=3001 --replicas=2 --output=yaml --show-managed-fields=true --namespace=fullstack-microservice-golang --dry-run=client

# -- Create a ClusterIP service for blog service 
kubectl create service clusterip blog-service --tcp=3001 --output=yaml --namespace=fullstack-microservice-golang --dry-run=client

# -- Blogs Frontend App deployment
kubectl create deployment blogs-app-frontend --image=asia-south1-docker.pkg.dev/sounish-cloud-workstation/fullstack-microservices-golang-gke/blogs-app-frontend:v241000.5219 --port=80 --replicas=2 --output=yaml --namespace=fullstack-microservice-golang --dry-run=client

# Blogs Frontend App Service
kubectl create service loadbalancer blogs-app-frontend --tcp=4200:80 --output=yaml --namespace=fullstack-microservice-golang --dry-run=client 