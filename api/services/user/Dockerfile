FROM golang:1.15.2-alpine AS builder

WORKDIR /go/src/github.com/Asuha-a/URLShortener/api/services/user

COPY . .

RUN go build -o /bin/user

FROM alpine:latest

COPY --from=builder /bin/user /bin/user

ENTRYPOINT [ "/bin/user" ]