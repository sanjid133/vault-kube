FROM golang:1.12 AS builder

ENV GO111MODULE=on \
  CGO_ENABLED=0 \
  GOOS=linux \
  GOARCH=amd64

WORKDIR /src
COPY . .

RUN ./build.sh

RUN which vkube
FROM alpine:latest

RUN apk add --no-cache --update ca-certificates


# Copy App binary to image
COPY --from=builder /go/bin/vkube /usr/local/bin/vkube
ENTRYPOINT ["vkube"]
