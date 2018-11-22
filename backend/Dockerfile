FROM golang:1.10 as builder

WORKDIR /go/src/github.com/goosetacob/asthtc/backend
COPY . .

# get dependencies and install binary
RUN go get -d -v ./...
RUN go install -v ./...

# copy binary into our base image.
FROM gcr.io/distroless/base
COPY --from=builder /go/bin/backend /

# start backend app
CMD ["/backend"]