FROM golang:1.21.5 as builder

WORKDIR /app
COPY . .

RUN go mod download && go mod verify

RUN CGO_ENABLED=0 GOOS=linux go build -o build/app ./cmd/app

# production stage
FROM alpine:latest
COPY --from=builder /app/build .
COPY --from=builder /app/config /config/
CMD ["./app"]