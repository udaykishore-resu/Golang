FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o ds-tool ./cmd/ds-tool

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/ds-tool .
COPY --from=builder /app/configs ./configs
EXPOSE 8080
ENV OTEL_SERVICE_NAME=data-structures-go
CMD ["./ds-tool"]