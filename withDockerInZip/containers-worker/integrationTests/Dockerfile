# Use an official Golang runtime as a parent image
FROM golang:1.20-alpine

ARG GIT_USER
ARG GIT_TOKEN

WORKDIR /app
ENV GOPRIVATE=github.com/checkmarxDev/*,github.com/CheckmarxDev/*

RUN env
RUN apk add --no-cache git && git config --global url."https://${GIT_USER}:${GIT_TOKEN}@github.com".insteadOf "https://github.com"


COPY go.mod .
COPY go.sum .

## Get dependencies - will also be cached if we won't change mod/sum
RUN go mod download -x

RUN go install github.com/onsi/ginkgo/v2/ginkgo
RUN go install github.com/onsi/gomega/...

# COPY the source code as the last step
COPY .. .

RUN ginkgo version

# Run the tests
CMD ["ginkgo", "./tests"]