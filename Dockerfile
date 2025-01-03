# build
FROM golang:1.22.5-alpine AS builder

WORKDIR /app

COPY . .

RUN go mod download
RUN go build main.go

# run
FROM alpine:latest AS runner

WORKDIR /

COPY --from=builder /app/main /main
COPY --from=builder /app/config /config

EXPOSE 8000

CMD ["/main"]
