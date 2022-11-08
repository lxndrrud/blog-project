docker exec -it backend-blog go run ./db/script.go -m=down
docker exec -it cacher-blog redis-cli -a 'redis_password' FLUSHALL