MODE=$1
SERVER=$2
export IMAGE_TAG=latest
if [ "${MODE}" == "zookeeper" ]; then
cd zookeeper
docker-compose -f zookeeper$SERVER.ulaw.com.yaml down
docker-compose -f zookeeper$SERVER.ulaw.com.yaml up -d
elif [ "${MODE}" == "kafka" ]; then
cd kafka
docker-compose -f kafka$SERVER.ulaw.com.yaml down
docker-compose -f kafka$SERVER.ulaw.com.yaml up -d
elif [ "${MODE}" == "orderer" ]; then
cd orderer
docker-compose -f orderer$SERVER.ulaw.com.yaml down
docker-compose -f orderer$SERVER.ulaw.com.yaml up -d
fi