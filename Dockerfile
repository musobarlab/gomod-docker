FROM golang:1.11.1-alpine3.8

# Add args
ARG APP_NAME=gomod-docker
ARG LOG_DIR=/${APP_NAME}/logs

# env build packages
ARG BUILD_PACKAGES="git curl"
# build packages with cert
# ARG BUILD_PACKAGES="git curl ca-certificates"

# Create folder logs 
RUN mkdir -p ${LOG_DIR}

# Set the Current Working Directory inside the container
WORKDIR /usr/app

# Add the source code
ENV SRC_DIR=/usr/app/

# Environment Variables
ENV LOG_FILE_LOCATION=${LOG_DIR}/app.log

COPY . $SRC_DIR

RUN apk update && apk add --no-cache $BUILD_PACKAGES \
  && go mod download \
  && apk del $BUILD_PACKAGES \
  && CGO_ENABLED=0 GOOS=linux go build -ldflags '-w -s' -a -o app .

# Copy environment variable to source dir
COPY .env $SRC_DIR.env

EXPOSE 8080
EXPOSE 8081

ENTRYPOINT ["sh", "-c", "./app"]