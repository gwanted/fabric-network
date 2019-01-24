ORG=$1
export IMAGE_TAG=latest
cd peer
docker-compose -f peer0.org$ORG.ulaw.com.yaml down
docker-compose -f peer0.org$ORG.ulaw.com.yaml up -d
docker-compose -f peer1.org$ORG.ulaw.com.yaml down
docker-compose -f peer1.org$ORG.ulaw.com.yaml up -d

cd ../couchdb
if [ "$ORG" == "1" ]; then
docker-compose -f couchdb1.ulaw.com.yaml down
docker-compose -f couchdb1.ulaw.com.yaml up -d
docker-compose -f couchdb2.ulaw.com.yaml down
docker-compose -f couchdb2.ulaw.com.yaml up -d
elif [ "$ORG" == "2" ]; then
docker-compose -f couchdb3.ulaw.com.yaml down
docker-compose -f couchdb3.ulaw.com.yaml up -d
docker-compose -f couchdb4.ulaw.com.yaml down
docker-compose -f couchdb4.ulaw.com.yaml up -d
fi