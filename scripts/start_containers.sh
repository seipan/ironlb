#!/bin/bash

docker run -d --name web1 --network bridge --ip 172.17.0.2 my-web-image
docker run -d --name web2 --network bridge --ip 172.17.0.3 my-web-image