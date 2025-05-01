#!/bin/bash
# Script to automate Grafana dashboard import using Grafana API

set -e

GRAFANA_URL=${GRAFANA_URL:-http://localhost:3000}
GRAFANA_API_KEY=${GRAFANA_API_KEY:-your_api_key_here}
DASHBOARD_FILE="monitoring/grafana-dashboard.json"

if [ ! -f "$DASHBOARD_FILE" ]; then
  echo "Dashboard file $DASHBOARD_FILE not found!"
  exit 1
fi

DASHBOARD_JSON=$(jq -c '.' "$DASHBOARD_FILE")

curl -X POST "$GRAFANA_URL/api/dashboards/db" \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $GRAFANA_API_KEY" \
  -d "{\"dashboard\": $DASHBOARD_JSON, \"overwrite\": true}"

echo "Dashboard imported successfully."
