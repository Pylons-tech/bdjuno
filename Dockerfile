FROM golang:1.17
WORKDIR /go/src/github.com/Pylons-tech/bdjuno
COPY . ./
COPY ./deploy/config /bdjuno/
RUN make build


WORKDIR /bdjuno
COPY ./build/bdjuno /usr/bin/bdjuno
#COPY /go/src/github.com/Pylons-tech/bdjuno/deploy/config /bdjuno/
CMD [ "/usr/bin/bdjuno", "parse", "--home", "/bdjuno/" ]
