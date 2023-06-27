#! /usr/bin/bash

# create a namespace in Kubernetes
kubectl create --namespace haiku-default

# replace the image name with a new one pushed to DockerHub
sed -i "s#imageName#ernestklu/haikubox:db-v1#g" kubernetes-yaml/db-deploy.yaml

# deploy the database and service in the created namespace
kubectl apply -f kubernetes-yaml/db-deploy.yaml -n haiku-default

