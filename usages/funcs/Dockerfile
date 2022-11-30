#docker build -t optional_params:v3 .
#docker run -it --rm optional_params:v3
#docker images | grep optional_params
# All above after ":". they are the tag name of docker, the tag name can none.

FROM golang:alpine AS builder
WORKDIR /build
COPY . .
RUN go build -o optional_params optional_params.go
FROM alpine
WORKDIR /build
COPY --from=builder /build/optional_params /build/optional_params
CMD ["./optional_params"]