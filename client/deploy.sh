#!/usr/bin/env bash
echo "building and pushing docker container image"
set -e
docker build -t huangjoyce3/unity-client .
docker push huangjoyce3/unity-client
go clean 

echo "now running script on droplet"
ssh -oStrictHostKeyChecking=no root@104.236.215.90 'bash -s' < run.sh

