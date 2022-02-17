FROM golang:1.18beta2-alpine3.15 AS build_stage

RUN apk add --no-cache git

WORKDIR /tmp/go-api

COPY go.mod .
# COPY go.sum .

RUN go mod download

COPY . .

# RUN CGO_ENABLED=0 go test -v

RUN go build -o ./out/executable .

FROM alpine:latest
RUN apk add ca-certificates

COPY --from=build_stage /tmp/go-api/out/executable /app/executable

EXPOSE 8080

CMD ["/app/executable"]