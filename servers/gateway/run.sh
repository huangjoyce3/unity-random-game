#!/usr/bin/env bash
docker pull huangjoyce3/unity-api
docker rm -f unityapi
docker network rm authNet

export TLSCERT=/etc/letsencrypt/live/unityapi.joycehuang.me/fullchain.pem
export TLSKEY=/etc/letsencrypt/live/unityapi.joycehuang.me/privkey.pem

# create new private network named authNet
echo "creating private network"
docker network create authNet

echo "running unityapi image in network"
docker run -d \
--name unityapi \
-p 443:443 \
-v /etc/letsencrypt:/etc/letsencrypt:ro \
-e TLSCERT=$TLSCERT \
-e TLSKEY=$TLSKEY \
--network authNet \
huangjoyce3/unity-api
