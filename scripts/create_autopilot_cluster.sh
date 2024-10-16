#!/bin/bash

gcloud beta container --project "sounish-cloud-workstation" clusters create-auto "sounish-autopilot-cluster-01" --region "asia-south1" --release-channel "regular" --network "projects/sounish-cloud-workstation/global/networks/default" --subnetwork "projects/sounish-cloud-workstation/regions/asia-south1/subnetworks/default" --cluster-ipv4-cidr "/17" --binauthz-evaluation-mode=DISABLED