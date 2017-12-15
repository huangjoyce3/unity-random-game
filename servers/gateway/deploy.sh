#!/usr/bin/env bash
set -e
echo "building go server for Linux..."
GOOS=linux go build
docker build -t huangjoyce3/unity-api .
docker push huangjoyce3/unity-api
go clean 

echo "now running script on droplet SERVER"
ssh -oStrictHostKeyChecking=no root@138.197.2.19 'bash -s' < run.sh

