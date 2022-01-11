FROM golang:alpine AS builder

RUN apk update && apk add --no-cache git
WORKDIR /app

COPY go.mod ./
RUN go mod download
COPY main.go ./

RUN go build -o nocodb-wrapper

FROM alpine

WORKDIR /
COPY --from=builder /app/nocodb-wrapper /go/bin/nocodb-wrapper
EXPOSE 8787

ENTRYPOINT ["/go/bin/nocodb-wrapper"]