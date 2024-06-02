# On Golang Will build the projects 
FROM golang:alpine as build

WORKDIR /consumer 

COPY . ./

RUN go mod tidy 

RUN go build -o web-consumer

# Stage Deployer
# OS Alpine Will Running The app with copy all data from build stage
FROM alpine as main

WORKDIR /consumer

COPY --from=build /consumer/web-consumer /consumer 

ENTRYPOINT ["/consumer/web-consumer"]

EXPOSE 8000