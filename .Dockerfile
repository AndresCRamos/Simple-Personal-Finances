# syntax=docker/dockerfile:1

##
## Build
##
FROM golang:1.18-buster AS build

WORKDIR /app

COPY ./app/go.mod ./
COPY ./app/go.sum ./
RUN go mod download

COPY ./app ./

RUN go build -o /app_bin

##
## Deploy
##
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /app_bin /app_bin

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/app_bin"]