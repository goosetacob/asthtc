FROM golang:1.10 as builder

WORKDIR /go/src/github.com/goosetacob/asthtc
COPY . .

# get dependencies and install binary
RUN go get -d -v ./backend
RUN go install -v ./backend

# copy binary into our base image.
FROM gcr.io/distroless/base
COPY --from=builder /go/bin/backend /

# start backend app
CMD ["/backend"]