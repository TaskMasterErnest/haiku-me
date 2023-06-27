#! /usr/bin/bash

# sleep for 65 seconds for database to initialize
sleep 65s

# replace the image name with a new one pushed to DockerHub
sed -i "s#imageName#ernestklu/haikubox:app-v1#g" kubernetes-yaml/app-deploy.yaml

# deploy the application and its service
kubectl apply -f kubernetes-yaml/app-deploy.yaml -n haiku-default