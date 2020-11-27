FROM golang:1.15.2-alpine AS builder

WORKDIR /go/src/github.com/Asuha-a/URLShortener/api/

COPY ./go.mod ./services/user/
COPY ./go.sum ./services/user/
RUN mkdir -p ./services/user/pb/user/
COPY ./pb/user/. ./services/user/pb/user/

WORKDIR /go/src/github.com/Asuha-a/URLShortener/api/services/user/
COPY ./services/user/. .

RUN go build -o /bin/user

FROM alpine:3.12.1

COPY --from=builder /bin/user /bin/user

ENTRYPOINT [ "/bin/user" ]