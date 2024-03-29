#!/usr/bin/env bash
set -euo pipefail

echo -en '\x00\x00\x00\x00\x13{"text": "explode"}' | curl -s -v \
  -X 'POST' \
  -H "Content-Type: application/grpc-web+json" \
  -H "Grpc-Encoding: identity" \
  -H "Grpc-Accept-Encoding: identity" \
  --data-binary @- \
  http://localhost:8080/ping.v1.PingService/Pings | od -c
