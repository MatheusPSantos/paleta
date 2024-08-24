#bin/bash
docker compose -f docker/docker-compose.yml up -d\
&& go run main.go