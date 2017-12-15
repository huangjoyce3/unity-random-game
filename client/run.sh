#!/usr/bin/env bash
echo "running run.sh"
docker pull huangjoyce3/unity-client
docker rm -f unityclient

docker run -d \
--name unityclient \
-p 80:80 -p 443:443 \
-v /etc/letsencrypt:/etc/letsencrypt:ro \
huangjoyce3/unity-client