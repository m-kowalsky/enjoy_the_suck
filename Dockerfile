FROM golang:1.24.1 AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo enjoy_the_suck .

FROM alpine
WORKDIR /app
COPY --from=builder /app/enjoy_the_suck .
COPY templates ./templates
RUN chmod +x /app/enjoy_the_suck
EXPOSE 80
ENTRYPOINT ["/app/enjoy_the_suck"]

