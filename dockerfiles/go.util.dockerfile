FROM golang:1.18beta2-alpine3.15

WORKDIR /go

ENTRYPOINT [ "go" ]