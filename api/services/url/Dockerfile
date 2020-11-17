FROM golang:1.15.2-alpine AS builder

WORKDIR /go/src/github.com/Asuha-a/URLShortener/api/services/url

COPY . .

RUN go build -o /bin/url

FROM alpine:3.12.1

COPY --from=builder /bin/url /bin/url

ENTRYPOINT [ "/bin/url" ]