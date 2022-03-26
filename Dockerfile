ARG IMG_GO=golang:1.17-alpine
ARG IMG_ALPINE=alpine:3.15

FROM $IMG_GO AS builder

RUN apk add --no-cache --update git \
    build-base

WORKDIR /app

COPY ./go.mod ./go.sum ./Makefile ./

RUN make install-deps

COPY ./ .

RUN make vendor && make build-auth && make build-clear && make build-migrate

FROM $IMG_ALPINE as image-auth

WORKDIR /app

COPY --from=builder /app/out/bin/cryptomath-auth ./

VOLUME ["/app/configs"]

RUN chmod +x /app/cryptomath-auth

EXPOSE 5002

ENTRYPOINT ["./cryptomath-auth"]

FROM $IMG_ALPINE as image-auth-clear

WORKDIR /app

COPY --from=builder /app/out/bin/cryptomath-auth-clear ./

VOLUME ["/app/configs"]

RUN chmod +x /app/cryptomath-auth-clear

ENTRYPOINT ["./cryptomath-auth-clear"]

FROM $IMG_ALPINE as image-auth-migrate

WORKDIR /app

COPY --from=builder /app/out/bin/cryptomath-auth-migrate ./
COPY --from=builder /app/migrations ./

RUN chmod +x /app/cryptomath-auth-migrate

ENTRYPOINT ["./cryptomath-auth-migrate"]

