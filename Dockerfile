FROM golang:1.21-alpine AS builder
RUN apk update && apk add --no-cache make git
WORKDIR /go/src/github.com/forbole/callisto/v4
COPY . ./
RUN go mod download
RUN make build

FROM alpine:latest
WORKDIR /callisto
COPY --from=builder /go/src/github.com/forbole/callisto/v4/build/callisto /usr/bin/callisto
CMD [ "callisto" ]