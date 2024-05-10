FROM golang:1.22-alpine AS builder
RUN set -eux; apk add --no-cache ca-certificates build-base;
RUN apk add git

WORKDIR /go/src/github.com/forbole/callisto
COPY . ./

RUN make docker-build

FROM alpine:latest
RUN apk update && apk add --no-cache ca-certificates build-base
WORKDIR /callisto
COPY --from=builder /go/src/github.com/forbole/callisto/build/callisto /usr/bin/callisto
CMD [ "callisto" ]
