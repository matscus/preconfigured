# Builder
FROM golang:1.15.2-alpine3.12 as builder

WORKDIR /application

RUN apk update && apk upgrade && \
    apk --update add git make bash

COPY . .

RUN make engine

# Distribution
FROM alpine:latest

RUN apk update && apk upgrade && \
    apk --update --no-cache add tzdata && \
    mkdir /application 

WORKDIR /application 

COPY --from=builder /application/engine /application

CMD ["/application/engine", "--dbhost", "${POSTHRESQL_NODE}", "--dbname", "${PUSHER_DB_NAME}"]