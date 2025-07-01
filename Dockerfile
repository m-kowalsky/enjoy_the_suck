FROM golang:1.24.1 AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo enjoy_the_suck .

FROM alpine
WORKDIR /app
COPY --from=builder /app/enjoy_the_suck .
COPY templates ./templates
RUN chmod +x /app/enjoy_the_suck
EXPOSE 8080
HEALTHCHECK --interval=10s --timeout=3s --start-period=30s --retries=3 \
  CMD wget --spider -q http://localhost:8080/ || exit 1
ENTRYPOINT ["/app/enjoy_the_suck"]

