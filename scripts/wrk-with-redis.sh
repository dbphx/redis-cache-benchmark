#!/bin/bash
echo "Benchmark WITH Redis..."
ENABLE_CACHE=true docker compose up -d --build
sleep 5
docker exec -it go-app sh -c 'echo $ENABLE_CACHE'
wrk -t8 -c200 -d30s "http://localhost:8080/user?id=1"
