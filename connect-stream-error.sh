#!/usr/bin/env bash
set -euo pipefail

echo -en '\x00\x00\x00\x00\x13{"text": "explode"}' | curl -s -v \
  -X 'POST' \
  -H "Content-Type: application/connect+json" \
  -H "Connect-Encoding: identity" \
  -H "Connect-Accept-Encoding: identity" \
  --data-binary @- \
  http://localhost:8080/ping.v1.PingService/Pings | od -c
