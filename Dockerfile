FROM golang:1.22-alpine AS builder
RUN apk update && apk add --no-cache make git
WORKDIR /go/src/github.com/stalwart-algoritmiclab/callisto
COPY . ./
RUN go mod download
RUN make build

FROM alpine:latest
WORKDIR /callisto
COPY --from=builder /go/src/github.com/stalwart-algoritmiclab/callisto/build/callisto /usr/bin/callisto
CMD [ "callisto" ]