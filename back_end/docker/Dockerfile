FROM golang:1.13.8-alpine as builder

ENV GO111MODULE on
ENV CGO_ENABLED 0
ENV GOARCH amd64
ENV GOOS linux
ENV HOME /go/src

WORKDIR $HOME

COPY ./back_end/go.mod $HOME/go.mod
COPY ./back_end/go.sum $HOME/go.sum
RUN set -ex && \
    apk --update --no-cache add git && \
    go mod download && \
    go get -u github.com/cosmtrek/air

COPY ./back_end/server.go $HOME/server.go
COPY ./back_end/api $HOME/api

RUN go build -v -o /app/api $HOME/server.go
CMD ["/app/api"]

FROM alpine:latest as prod

WORKDIR /app
COPY --from=builder /app/api /app/api

CMD ["/app/api"]
