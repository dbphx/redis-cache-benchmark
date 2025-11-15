#!/bin/bash
echo "Benchmark WITHOUT Redis..."
ENABLE_CACHE=false docker compose up -d --build
sleep 5
docker exec -it go-app sh -c 'echo $ENABLE_CACHE'
wrk -t8 -c200 -d30s "http://localhost:8080/user?id=10000"
