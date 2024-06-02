# On Golang Will build the projects 
FROM golang:alpine as build

WORKDIR /service 

COPY . ./

RUN go mod tidy 

RUN go build -o web-service

# Stage Deployer
# OS Alpine Will Running The app with copy all data from build stage
FROM alpine as main

WORKDIR /service

COPY --from=build /service/web-service /service 

ENTRYPOINT ["/service/web-service", "serve"]

EXPOSE 8000