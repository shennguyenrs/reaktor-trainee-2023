docker-compose -f ./docker-compose.prod.yaml down
docker image rm birdnest-api
docker image rm birdnest-front
docker image rm crontasks
docker image rm nginx-proxy
docker volume prune
docker network prune
