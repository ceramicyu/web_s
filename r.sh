#!/bin/sh

sudo docker build -t bee-sport:1.0  .

sudo docker stop bee-sports && sudo docker rm bee-sports && sudo docker run -p 15555:8080 -d --name bee-sports  bee-sport:1.0
