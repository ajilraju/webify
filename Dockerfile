FROM golang:alpine AS builder
WORKDIR /go/src/github.com/ajilraju/webify
COPY . .
RUN go install .

FROM alpine:latest
WORKDIR /app
COPY --from=builder /go/bin/webify .
ENTRYPOINT [ "./webify" ]
