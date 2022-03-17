FROM golang:1.17 AS builder
WORKDIR /go/src/github.com/Pylons-tech/bdjuno
COPY . ./
RUN make build

FROM alpine:latest
WORKDIR /bdjuno
COPY --from=builder /go/src/github.com/Pylons-tech/bdjuno/build/bdjuno /usr/bin/bdjuno
CMD [ "bdjuno", "parse" ]
