FROM golang:1.13.8-alpine as builder

WORKDIR /app
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux

COPY go.mod .
# COPY go.sum .
# RUN set -ex && \
#     apk add --no-cache git && \
#     go mod download

COPY server.go .
COPY ./api ./api
RUN set -ex && \
    go build -a -o /go/bin/api .


FROM alpine:latest as prod

WORKDIR /app
COPY --from=builder /go/bin/api .
CMD ["/app/api"]

FROM golang:1.13.8-alpine as dev

WORKDIR /app
ENV GO111MODULE=on
COPY go.mod .
COPY ./api ./api
COPY server.go .
COPY .realize.yaml .

RUN set -ex && \
    apk update && \
    apk add --no-cache git && \
    : "for realize hot reloader" && \
    go get github.com/oxequa/realize

EXPOSE 1998
CMD [ "realize", "start" ]