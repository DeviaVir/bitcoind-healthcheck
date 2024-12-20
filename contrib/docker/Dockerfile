# Start by building the application.
FROM golang:1.23 AS build

ENV CGO_ENABLED=0

WORKDIR /go/src/app
ADD . /go/src/app

RUN go mod tidy \
 && go build -o /go/bin/app

# Now copy it into our base image.
FROM gcr.io/distroless/base-debian12
COPY --from=build /go/bin/app /
CMD ["/app"]

