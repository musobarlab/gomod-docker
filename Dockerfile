FROM golang:1.11.1-alpine3.8

WORKDIR /usr/app

# Add the source code
ENV SRC_DIR=/usr/app/

# build packages with cert
# ENV BUILD_PACKAGES="git curl ca-certificates"

# build packages
ENV BUILD_PACKAGES="git curl"

ADD . $SRC_DIR

RUN apk update && apk add --no-cache $BUILD_PACKAGES \
  && go mod download \
  && apk del $BUILD_PACKAGES \
  && CGO_ENABLED=0 GOOS=linux go build -ldflags '-w -s' -a -o app .

# Copy environment variable to source dir
COPY .env $SRC_DIR.env

EXPOSE 8080
EXPOSE 8081

ENTRYPOINT ["sh", "-c", "./app"]