ORG=$1
PEER=$2
export IMAGE_TAG=latest
cd peer
docker-compose -f peer$PEER.org$ORG.ulaw.com.yaml down
docker-compose -f peer$PEER.org$ORG.ulaw.com.yaml up -d

cd ../couchdb
if [ "$PEER" == "0" ]; then
docker-compose -f couchdb1.ulaw.com.yaml down
docker-compose -f couchdb1.ulaw.com.yaml up -d
elif [ "$PEER" == "1" ]; then
docker-compose -f couchdb2.ulaw.com.yaml down
docker-compose -f couchdb2.ulaw.com.yaml up -d
fi