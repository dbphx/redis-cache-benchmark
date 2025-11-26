# Benchmark: MySQL vs MySQL + Redis Cache

This project benchmarks the performance difference between querying
MySQL directly and with Redis cache.

## Architecture

    wrk -> go-app -> MySQL
                \-> Redis

## How to run

### WITHOUT Redis

    ./scripts/wrk-no-redis.sh

### WITH Redis

    ./scripts/wrk-with-redis.sh

## Results

### WITHOUT Redis

-   Requests/sec: 9447.11
-   Latency avg: 63.97ms
-   Non-2xx: 53086

### WITH Redis

-   Requests/sec: 34099.31
-   Latency avg: 58.64ms

## Notes

Redis improves throughput \~3.6x and reduces MySQL load.
