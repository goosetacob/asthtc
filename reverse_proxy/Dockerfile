FROM golang:1.10 as builder

WORKDIR /go/src/github.com/goosetacob/asthtc
COPY . .

# get dependencies and install binary
RUN go get -d -v ./reverse_proxy
RUN go install -v ./reverse_proxy

# copy binary into our base image.
FROM gcr.io/distroless/base
COPY --from=builder /go/bin/reverse_proxy /

# assume running as an individual container (localhost => host.docker.internal)
ENV RPC_ENDPOINT "host.docker.internal:80"

# start backend app
CMD ["/reverse_proxy"]