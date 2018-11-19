#run this locally so that you do not need to restore all the time the external dependencies (go get)
#of course, in a CI/CD environment, you need to change this approach 
docker rm $(docker ps -aqf "name=microservice-stockinfo")
docker build -t local/microservice-stockinfo .
docker tag local/microservice-stockinfo gcr.io/itdays-201118/microservice-stockinfo
docker run \
    --name microservice-stockinfo \
    -e STOCKINFO_LISTEN_ADDR=':3001' \
    -p 3001:3001 \
    local/microservice-stockinfo
