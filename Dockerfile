# build
FROM golang:1.22.5-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build main.go

# run
FROM alpine:latest AS runner

WORKDIR /

COPY --from=builder /app/main /main

EXPOSE 8000

CMD ["/main"]
