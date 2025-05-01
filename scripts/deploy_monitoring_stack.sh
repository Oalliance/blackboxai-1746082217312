#!/bin/bash
# Script to deploy Prometheus monitoring stack with Helm and setup Grafana dashboards

set -e

NAMESPACE="monitoring"
RELEASE_NAME="prometheus"
VALUES_FILE="monitoring/prometheus-values.yml"

# Create namespace if not exists
kubectl get namespace $NAMESPACE || kubectl create namespace $NAMESPACE

# Add Prometheus community Helm repo if not added
helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
helm repo update

# Install or upgrade Prometheus stack
helm upgrade --install $RELEASE_NAME prometheus-community/kube-prometheus-stack \
  --namespace $NAMESPACE \
  -f $VALUES_FILE

# Wait for Prometheus server pod to be ready
kubectl rollout status deployment/${RELEASE_NAME}-kube-prometheus-prometheus -n $NAMESPACE

# Setup Grafana admin password secret (optional)
kubectl create secret generic grafana-admin-credentials \
  --from-literal=admin-user=admin \
  --from-literal=admin-password=admin123 \
  -n $NAMESPACE --dry-run=client -o yaml | kubectl apply -f -

echo "Prometheus monitoring stack deployed successfully in namespace $NAMESPACE."
echo "Access Grafana via the LoadBalancer service or port-forward:"
echo "kubectl -n $NAMESPACE port-forward svc/${RELEASE_NAME}-grafana 3000:80"
echo "Login with admin/admin123 or your configured credentials."
