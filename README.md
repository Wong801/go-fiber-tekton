### HOW TO RUN

## Install Minikube and Tekton

note: 
  you can use `kubectl <command>` or `minikube kubectl -- <command>`
  I am using `alias kubectl="minikube kubectl --"` since I am not installing kubectl

```bash
$ brew install minikube
$ brew install kubectl # Optional
$ kubectl apply --filename \
https://storage.googleapis.com/tekton-releases/pipeline/latest/release.yaml
```

## Set Kaniko Secret

Run this command to set your dockerhub token

```bash
$ echo -n 'username:pat_token' | base64
```

Replace the value into config.json

Run this command to set kaniko secret

```bash
$ kubectl create secret generic kaniko-secret --from-file=config.json=./config.json
```

## Preparing Pipeline Run

Go to tekton-pipeline/pipeline-run.yaml

Change the CONSUMER_IMAGE and SERVICE_IMAGE params to match your dockerhub repo

## Setup Build Environment

```bash
$ kubectl apply -f tekton-pipeline/persistent-volume.yaml
$ kubectl apply -f tekton-pipeline/persistent-volume-claim.yaml
```

Then, run this command to copy current dir into pvc

```bash
$ kubectl apply -f tekton-pipeline/temp-pod.yaml
$ kubectl cp . temp-pod:/mnt/data/source
$ kubectl delete pod temp-pod
```

## Running Deployment

```bash
$ kubectl apply -f tekton-pipeline/kaniko-task.yaml
$ kubectl apply -f tekton-pipeline/pipeline.yaml
$ kubectl apply -f tekton-pipeline/pipeline-run.yaml
```

Feel free to run this command after all completed

```bash
$ kubectl delete pipelinerun build-and-deploy-run
```

## Running Service

Edit the image source in deployment task to match your dockerhub image

```bash
$ kubectl apply -f tekton-pipeline/deployment-and-service.yaml
```

### TODO

- Connect to Postgres
- Connect to RabbitMQ
- Connect to Redis (optional)
