# Build Geth in a stock Go builder container
FROM golang:1.11-alpine as builder

RUN apk add --no-cache make gcc musl-dev linux-headers

ADD . /go/src/Didux-blackbox
RUN cd /go/src/Didux-blackbox && make build

# Pull Geth into a second stage deploy alpine container
FROM alpine:latest

RUN apk add --no-cache ca-certificates
COPY --from=builder /go/src/Didux-blackbox/bin/blackbox /usr/local/bin/

EXPOSE 9000
ENTRYPOINT ["blackbox"]
