FROM golang:1.21.5-alpine3.18 as build-env
RUN apk add build-base

ARG GIT_USER
ARG GIT_TOKEN

RUN test -n "$GIT_TOKEN" || (echo "Error: GIT_TOKEN not set" && false)
RUN test -n "$GIT_USER" || (echo "Error: GIT_USER not set" && false)

WORKDIR /app
ENV GOPRIVATE=github.com/checkmarxDev/*,github.com/CheckmarxDev/*

RUN env

RUN apk add --no-cache git && git config --global url."https://${GIT_USER}:${GIT_TOKEN}@github.com".insteadOf "https://github.com"

#Copy go mod and sum files
COPY go.mod .
COPY go.sum .
#
## Get dependencies - will also be cached if we won't change mod/sum
RUN go mod download -x
# COPY the source code as the last step
COPY . .
#run unit tests
RUN go test -v ./...

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -a -installsuffix cgo -o containers-worker cmd/main.go cmd/config.go cmd/health.go

#RUN cd pkg/scanner && go test
#RUN cd internal/sca && go test
WORKDIR /app

#runtime image
FROM alpine:3.14.0

ARG COMMIT_SHA
ARG RELEASE_TAG


LABEL cx.commit-sha ${COMMIT_SHA}
LABEL cx.release-tag ${RELEASE_TAG}

COPY --from=build-env /app/containers-worker /app/containers-worker

RUN apk add libcap && setcap 'cap_net_bind_service=+ep' /app/containers-worker

EXPOSE 80

USER 1001

CMD ["/app/containers-worker"]
