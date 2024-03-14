#!/usr/bin/env bash
set -euo pipefail

curl -s -v \
  'http://localhost:8080/ping.v1.PingService/Ping?encoding=json&message=%7B%22text%22%3A%22hello%22%7D' | jq .
