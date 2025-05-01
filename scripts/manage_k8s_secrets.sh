#!/bin/bash
# Script to manage Kubernetes secrets for logistics marketplace

set -e

NAMESPACE=${1:-logistics}
SECRET_NAME=${2:-logistics-marketplace-secrets}

echo "Creating/updating Kubernetes secret '$SECRET_NAME' in namespace '$NAMESPACE'..."

kubectl create namespace $NAMESPACE --dry-run=client -o yaml | kubectl apply -f -

kubectl delete secret $SECRET_NAME -n $NAMESPACE --ignore-not-found

kubectl create secret generic $SECRET_NAME \
  --from-literal=DATABASE_URL="$DATABASE_URL" \
  --from-literal=AWS_ACCESS_KEY_ID="$AWS_ACCESS_KEY_ID" \
  --from-literal=AWS_SECRET_ACCESS_KEY="$AWS_SECRET_ACCESS_KEY" \
  -n $NAMESPACE

echo "Secret '$SECRET_NAME' created/updated successfully in namespace '$NAMESPACE'."
