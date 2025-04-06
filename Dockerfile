ARG GO_IMAGE_NAME
ARG GO_IMAGE_VERSION

FROM ${GO_IMAGE_NAME}:${GO_IMAGE_VERSION}


RUN apk add --no-cache curl \
    && curl -L https://github.com/golang-migrate/migrate/releases/download/v4.16.2/migrate.linux-amd64.tar.gz | tar xvz -C /usr/local/bin \
    && chmod +x /usr/local/bin/migrate \
    && apk add --no-cache make

# Set working directory
WORKDIR /home/app

RUN go install github.com/air-verse/air@latest && \
    go install github.com/swaggo/swag/cmd/swag@latest
