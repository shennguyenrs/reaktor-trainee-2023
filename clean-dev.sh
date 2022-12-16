docker-compose -f ./docker-compose.dev.yaml down
docker image rm birdnest-api
docker image rm crontasks
docker volume prune
docker network prune
