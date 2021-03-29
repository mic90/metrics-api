ARG GOLANG_VERSION=1.16
ARG IMAGE_TAG=buster

# Build app
FROM golang:${GOLANG_VERSION}-${IMAGE_TAG} as builder

WORKDIR /go/src/app
ADD . /go/src/app

RUN go get -d -v ./...

RUN go build -o /go/bin/app

# Copy binary to final image and run
FROM gcr.io/distroless/base-debian10 as runner

COPY --from=builder /go/bin/app /

EXPOSE 8080

CMD ["/app"]