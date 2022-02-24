FROM golang:1.16-alpine AS builder
RUN apk update && apk add --no-cache make
WORKDIR /go/src/github.com/pylons-tech/bdjuno
COPY . ./
RUN make build

FROM alpine:latest
WORKDIR /bdjuno
COPY --from=builder /go/src/github.com/pylons-tech/bdjuno/build/bdjuno /usr/bin/bdjuno
CMD [ "bdjuno" ]
