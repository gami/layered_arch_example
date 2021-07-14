FROM golang:1.16.6-alpine3.13

ENV GOOS linux
ENV GOARCH amd64
ENV PATH $PATH:/go/bin/linux_amd64
RUN \
  echo http://dl-cdn.alpinelinux.org/alpine/edge/community/ >> /etc/apk/repositories && \
  apk update --no-cache && \
  apk add --virtual .build \
  git \
  make \
  mysql-client \
  tzdata && \
  ln -f -s /usr/share/zoneinfo/Asia/Tokyo /etc/localtime

WORKDIR /app
