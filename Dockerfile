FROM golang:alpine3.12 as builder

RUN apk add git
RUN mkdir -p /go/src/github.com/masato25
ENV GOPATH=/go
WORKDIR /go/src/github.com/masato25
RUN git clone https://github.com/masato25/go-wild-dns
RUN cd go-wild-dns && go build

FROM alpine:3.7
RUN mkdir -p /opt/go-wild-dns
COPY --from=builder /go/src/github.com/masato25/go-wild-dns/go-wild-dns /opt/go-wild-dns
WORKDIR /opt/go-wild-dns
COPY go-wild-dns .
CMD ./go-wild-dns
