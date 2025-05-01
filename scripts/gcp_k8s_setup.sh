#!/bin/bash
# Script to create GCP service account, generate key, and configure Kubernetes cluster access

set -e

PROJECT_ID=$1
SERVICE_ACCOUNT_NAME="terraform-deployer"
K8S_CLUSTER_NAME=$2
K8S_CLUSTER_ZONE=$3

if [ -z "$PROJECT_ID" ] || [ -z "$K8S_CLUSTER_NAME" ] || [ -z "$K8S_CLUSTER_ZONE" ]; then
  echo "Usage: $0 <gcp-project-id> <k8s-cluster-name> <k8s-cluster-zone>"
  exit 1
fi

echo "Creating service account: $SERVICE_ACCOUNT_NAME in project $PROJECT_ID"
gcloud iam service-accounts create $SERVICE_ACCOUNT_NAME --project $PROJECT_ID

echo "Assigning roles to service account"
gcloud projects add-iam-policy-binding $PROJECT_ID \
  --member="serviceAccount:$SERVICE_ACCOUNT_NAME@$PROJECT_ID.iam.gserviceaccount.com" \
  --role="roles/container.admin"

gcloud projects add-iam-policy-binding $PROJECT_ID \
  --member="serviceAccount:$SERVICE_ACCOUNT_NAME@$PROJECT_ID.iam.gserviceaccount.com" \
  --role="roles/storage.admin"

gcloud projects add-iam-policy-binding $PROJECT_ID \
  --member="serviceAccount:$SERVICE_ACCOUNT_NAME@$PROJECT_ID.iam.gserviceaccount.com" \
  --role="roles/iam.serviceAccountUser"

echo "Creating service account key"
gcloud iam service-accounts keys create ./gcp-key.json \
  --iam-account=$SERVICE_ACCOUNT_NAME@$PROJECT_ID.iam.gserviceaccount.com \
  --project $PROJECT_ID

echo "Setting up kubectl context for cluster $K8S_CLUSTER_NAME"
gcloud container clusters get-credentials $K8S_CLUSTER_NAME --zone $K8S_CLUSTER_ZONE --project $PROJECT_ID

echo "Extracting Kubernetes API server endpoint"
K8S_HOST=$(kubectl config view --minify -o jsonpath='{.clusters[0].cluster.server}')
echo "K8S_HOST=$K8S_HOST"

echo "Extracting Kubernetes cluster CA certificate"
K8S_CA_CERT=$(kubectl config view --minify -o jsonpath='{.clusters[0].cluster.certificate-authority-data}')
echo "K8S_CLUSTER_CA_CERTIFICATE=$K8S_CA_CERT"

echo "Extracting Kubernetes user token"
USER_NAME=$(kubectl config view --minify -o jsonpath='{.users[0].name}')
K8S_TOKEN=$(kubectl config view --minify -o jsonpath="{.users[0].user.token}")
echo "K8S_TOKEN=$K8S_TOKEN"

echo "Export the following environment variables or add them as GitHub secrets:"
echo "GCP_CREDENTIALS=$(cat ./gcp-key.json | base64 -w 0)"
echo "K8S_HOST=$K8S_HOST"
echo "K8S_CLUSTER_CA_CERTIFICATE=$K8S_CA_CERT"
echo "K8S_TOKEN=$K8S_TOKEN"

echo "Cleanup: You may want to securely store or delete the gcp-key.json file after use."
