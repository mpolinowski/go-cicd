FROM golang:alpine as build

WORKDIR /build

COPY . .

RUN go mod tidy

RUN go build -o hi

FROM alpine:latest

COPY --from=build /build/hi /usr/local/bin/hi

ENTRYPOINT ["/usr/local/bin/hi"]

# RUN ["/bin/ash", "-c", "/usr/local/bin/hi"]
