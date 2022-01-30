FROM golang:1.17 as builder

WORKDIR /build

COPY go.mod go.sum ./

RUN go mod download

COPY *.go ./

# `skaffold debug` sets SKAFFOLD_GO_GCFLAGS to disable compiler optimizations
ARG SKAFFOLD_GO_GCFLAGS
ENV GOARCH=amd64
ENV GOOS=linux
ENV CGO_ENABLED=0

RUN go build -gcflags="${SKAFFOLD_GO_GCFLAGS}" -o app

FROM alpine:3

# Define GOTRACEBACK to mark this container as using the Go language runtime
# for `skaffold debug` (https://skaffold.dev/docs/workflows/debug/).
ENV GOTRACEBACK=single

WORKDIR /app

COPY --from=builder /build/app run

ENTRYPOINT ["/app/run"]
