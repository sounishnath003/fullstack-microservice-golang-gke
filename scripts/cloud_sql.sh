#!/bin/bash
# Refer blog: https://cloud.google.com/sql/docs/postgres/connect-kubernetes-engine#service-account-key

gcloud sql instances describe postgres-instance-demo

# Create Kubenetes secrets for cloud sql in GKE
kubectl create secret generic postgres-cloudsql-secret \
  --from-literal=username=root \
  --from-literal=password=password \
  --from-literal=authdb=auth \
  --from-literal=blogsdb=blogs \
  -n "fullstack-microservice-golang"

# Create a Kube SA account to connect to Cloud SQL via auth proxy
kubectl apply -f service-account.yaml

# -- to connect cloud sql from GKE
gcloud iam service-accounts add-iam-policy-binding \
--role="roles/iam.workloadIdentityUser" \
--member="serviceAccount:sounish-cloud-workstation.svc.id.goog[fullstack-microservice-golang/googlekubernetes-cloudsql-sa]" \
797087556919-compute@sounish-cloud-workstation.iam.gserviceaccount.com

# Generate the key.json for SA
gcloud iam service-accounts keys create ~/key.json \
--iam-account=797087556919-compute@developer.gserviceaccount.com@project-id.iam.gserviceaccount.com

# Turn your service account key into a k8s Secret:
kubectl create secret generic cloudsql-connect-ksa-gsa \
--from-file=service_account.json=~/key.json

# Mount the secret as a volume under the spec: for your k8s object:
volumes:
- name: cloudsql-connect-ksa-gsa-vol
  secret:
    secretName: cloudsql-connect-ksa-gsa
