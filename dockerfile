# =======================================================================
# Builder
# =======================================================================
FROM golang:1.13.7-alpine3.11 AS builder

RUN apk update && apk upgrade && \
    mkdir -p /go/src/github.com/bagustyo92/auth && \
    apk add git

WORKDIR /go/src/github.com/bagustyo92/auth
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o auth

# =======================================================================
# Distribution
# =======================================================================
FROM alpine:latest
RUN apk update && apk upgrade
# This log file just add because of error that it caused
RUN mkdir log 
COPY --from=builder /go/src/github.com/bagustyo92/auth .
CMD [ "./auth" ]