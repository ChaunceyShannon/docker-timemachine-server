FROM golang:1.17.3-buster as golang-builder-base

WORKDIR /src

COPY go.* ./
RUN go mod download 

# --- 

FROM golang-builder-base as golang-builder

WORKDIR /app

ENV CGO_ENABLED=1

COPY go.* ./
COPY *.go ./

RUN --mount=type=cache,target=/root/.cache/go-build go build -o run -ldflags " -a -s -w -extldflags '-static'"

# --- 

FROM chaunceyshannon/cicd-tools:1.0.0 as upx-builder

ARG BIN_NAME=run

WORKDIR /app

COPY --from=golang-builder /app/${BIN_NAME} ./

RUN upx -9 ${BIN_NAME}

# --- 

FROM ubuntu:20.04

COPY --from=upx-builder /app/run /bin/run

RUN apt update;\
    DEBIAN_FRONTEND=noninteractive apt install netatalk -y;\
    apt clean all

ENTRYPOINT ["/bin/run"]
