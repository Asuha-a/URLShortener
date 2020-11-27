FROM golang:1.15.2-alpine AS builder

WORKDIR /go/src/github.com/Asuha-a/URLShortener/api/

COPY ./go.mod ./services/url/
COPY ./go.sum ./services/url/
RUN mkdir -p ./services/url/pb/url/
COPY ./pb/url/. ./services/url/pb/url/
RUN mkdir ./services/url/utility/
COPY ./utility/. ./services/url/utility/

WORKDIR /go/src/github.com/Asuha-a/URLShortener/api/services/url/
COPY ./services/url/. .

RUN go build -o /bin/url

FROM alpine:3.12.1

COPY --from=builder /bin/url /bin/url

ENTRYPOINT [ "/bin/url" ]