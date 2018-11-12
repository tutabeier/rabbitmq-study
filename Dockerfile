FROM golang:1.11-alpine
ARG type
ENV TYPE=${type}
ENV DEP_VERSION="0.5.0"

WORKDIR $GOPATH/src/github.com/tutabeier/rabbitmq/${type}

RUN apk update; \
    apk add --no-cache \
        ca-certificates \
        curl \
        git \
        make \
        openssl; \
    curl -L -s https://github.com/golang/dep/releases/download/v${DEP_VERSION}/dep-linux-amd64 -o /bin/dep; \
    chmod +x /bin/dep; \
    rm -rf /var/cache/apk/*; \
    rm -rf /tmp/*;
COPY Gopkg.toml Gopkg.lock ./
RUN dep ensure --vendor-only
COPY ${type} ${type}

CMD go run $TYPE/cmd/main.go
