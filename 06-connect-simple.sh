#!/usr/bin/env bash
set -euo pipefail

curl -s -v \
  --json '{"text": "hello"}' \
  http://localhost:8080/ping.v1.PingService/Ping | jq .
