#!/bin/bash
# Simple load testing script using k6

set -e

TARGET_URL=${1:-"http://localhost:8080/api/endpoint"}
DURATION=${2:-"30s"}
VU=${3:-10}

echo "Starting load test against $TARGET_URL with $VU virtual users for $DURATION..."

k6 run --vus $VU --duration $DURATION - <(cat <<EOF
import http from 'k6/http';
import { check, sleep } from 'k6';

export default function () {
  let res = http.get('$TARGET_URL');
  check(res, {
    'status is 200': (r) => r.status === 200,
  });
  sleep(1);
}
EOF
)

echo "Load test completed."
