# syntax=docker/dockerfile:1

FROM golang:1.20.2-alpine3.17

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./
COPY .env ./

COPY api/ ./api/
COPY pkg/ ./pkg/
COPY settings/ ./settings/

ENV DOCKERIZE_VERSION v0.6.1
RUN wget https://github.com/jwilder/dockerize/releases/download/$DOCKERIZE_VERSION/dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && tar -C /usr/local/bin -xzvf dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && rm dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz

RUN go build -o /financial-template-api

EXPOSE 8081

ENTRYPOINT ["financial-template-api"]