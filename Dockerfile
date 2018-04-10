FROM golang:1.10-alpine as builder

ARG REPO

RUN apk add --no-cache \
  build-base \
  git \
  make \
  curl

RUN go get github.com/Masterminds/glide

WORKDIR /go/src/github.com/$REPO
COPY . .
RUN glide install
RUN go build -v -a -o /usr/local/bin/catmd ./cmd/catmd

FROM alpine:latest
COPY --from=builder /usr/local/bin/catmd /usr/local/bin/catmd
ENTRYPOINT ["/usr/local/bin/catmd"]