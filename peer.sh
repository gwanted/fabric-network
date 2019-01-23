SERVER=$1
ORG=$2
cd peer
docker-compose -f peer{$SERVER}.org{$ORG}.ulaw.com.yaml down
docker-compose -f peer{$SERVER}.org{$ORG}.ulaw.com.yaml up -d

if [ ${SERVER} == 0 && ${ORG} == 1]; then
cd couchdb
docker-compose -f couchdb1.ulaw.com.yaml down
docker-compose -f couchdb1.ulaw.com.yaml up -d
elif [ ${SERVER} == 1 && ${ORG} == 1]; then
cd couchdb
docker-compose -f couchdb2.ulaw.com.yaml down
docker-compose -f couchdb2.ulaw.com.yaml up -d
elif [ ${SERVER} == 0 && ${ORG} == 2]; then
cd couchdb
docker-compose -f couchdb3.ulaw.com.yaml down
docker-compose -f couchdb3.ulaw.com.yaml up -d
elif [ ${SERVER} == 1 && ${ORG} == 2]; then
cd couchdb
docker-compose -f couchdb4.ulaw.com.yaml down
docker-compose -f couchdb4.ulaw.com.yaml up -d
fi