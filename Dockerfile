FROM golang:1.17-alpine AS builder
RUN mkdir /build
ADD controllers /build/controllers
ADD handlers /build/handlers
ADD models /build/models
ADD go.mod go.sum main.go tree-wal.db /build/
WORKDIR /build
RUN apk add --update gcc musl-dev
RUN CGO_ENABLED=1 GOOS=linux go build -o tree-web-server

FROM alpine:3.14
COPY --from=builder /build/tree-web-server /app/
COPY tree-wal.db /app/tree-wal.db
WORKDIR /app
CMD ["./tree-web-server"]