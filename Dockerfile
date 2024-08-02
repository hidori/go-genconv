FROM golang:1.22-alpine3.20 AS builder

RUN apk update \
    && apk add --no-cache \
    ca-certificates \
    && update-ca-certificates

COPY . /src
WORKDIR /src

RUN go build -o /genmapper ./cmd/genmapper/main.go

FROM alpine:3.20 AS runner

COPY --from=builder /genmapper /usr/local/bin

ENTRYPOINT [ "genmapper" ]
