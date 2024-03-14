#!/usr/bin/env bash
set -euo pipefail

curl -s -v \
  -X 'POST' \
  -H "Content-Type: application/json" \
  -H "Content-Encoding: identity" \
  -H "Accept-Encoding: identity" \
  --data '{"text": "hello"}' \
  http://localhost:8080/ping.v1.PingService/Ping | jq .
