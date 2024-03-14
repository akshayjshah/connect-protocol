#!/usr/bin/env bash
set -euo pipefail

buf curl \
  http://localhost:8080/ping.v1.PingService/Ping \
  --http2-prior-knowledge \
  --protocol grpc \
  --data '{"text": "hello"}'
