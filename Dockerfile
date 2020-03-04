# import go image from hub
FROM golang:1.11
LABEL maintainer="yerassyl"
RUN export GO111MODULE=on
WORKDIR /marketplace
# firstly, copy .mod and .sum files (which nothing, but the list of external packages that are used in this project)
COPY go.mod go.sum ./
# install those packages
RUN go mod download
# copy everthing to the
COPY . .
# this creates a binary file
# which will be run by the docker-compose.yml file
# refer to command in docker-compose.yml file
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main .

