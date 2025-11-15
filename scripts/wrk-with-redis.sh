#!/bin/bash
echo "Benchmark WITH Redis..."
ENABLE_CACHE=true docker compose up -d --build
sleep 5
wrk -t8 -c200 -d30s "http://localhost:8080/user?id=1"
