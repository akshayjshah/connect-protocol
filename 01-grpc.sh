#!/usr/bin/env bash
set -euo pipefail

echo -en '\x00\x00\x00\x00\x11{"text": "hello"}' | curl -s -v --http2 \
  -X 'POST' \
  -H "Content-Type: application/grpc+json" \
  -H "Grpc-Timeout: 1S" \
  -H "Grpc-Encoding: identity" \
  -H "Grpc-Accept-Encoding: identity" \
  -H "TE: trailers" \
  --data-binary @- \
  http://localhost:8080/ping.v1.PingService/Ping | od -c
