#!/bin/bash

# -- to access the GKE from instance gcloud auth
gcloud container clusters get-credentials "sounish-autopilot-cluster-01" \
    --region="asia-south1"