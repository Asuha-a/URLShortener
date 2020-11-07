FROM golang:1.15.2-alpine AS builder

WORKDIR /go/src/github.com/Asuha-a/URLShortener/api

COPY . .

RUN go build -o /bin/gateway

FROM alpine:latest

COPY --from=builder /bin/gateway /bin/gateway

ENTRYPOINT [ "/bin/gateway" ]