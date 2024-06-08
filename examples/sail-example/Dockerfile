# syntax=docker/dockerfile:1

FROM golang:1.22-alpine3.18 AS builder
ENV GOPROXY=https://goproxy.io,direct

WORKDIR /app/examples
ADD . .
RUN --mount=type=cache,target=$GOPATH/pkg,id=examples-gomodcache \
    --mount=type=cache,target=/root/.cache/go-build,id=examples-gocache \
    go mod download && go build -trimpath -tags 'go_json' -ldflags '-w -s' -o examples

FROM alpine:3.18
WORKDIR /app/examples
COPY --from=builder /app/examples/examples .
COPY --from=builder /app/examples/config.yaml .

VOLUME ["/app/paopao-ce/custom"]
EXPOSE 8080 8081 8082
HEALTHCHECK --interval=5s --timeout=3s  --retries=3  CMD ps -ef | grep examples || exit 1

ENTRYPOINT ["/app/examples"]
CMD ["serve"]
