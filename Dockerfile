FROM golang:1.11-alpine
ARG type
ADD ${type} .
CMD ["go", "run", "/go/cmd/main.go"]
