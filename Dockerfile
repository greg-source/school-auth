FROM golang:1.21 as builder
WORKDIR /app
COPY . .
RUN go build -o school-auth ./cmd/main.go

FROM busybox
WORKDIR /app
COPY --from=builder /app/school-auth .
COPY --from=builder /app/config/config.env ./config/
COPY --from=builder /app/migrations ./migrations

ENTRYPOINT ["/app/school-auth"]
