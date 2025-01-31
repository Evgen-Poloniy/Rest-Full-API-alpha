FROM golang:1.22.2 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY src/ ./src/
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o Go-API ./src

FROM scratch

COPY --from=builder /app/Go-API /Go-API

EXPOSE 8080

CMD ["./Go-API"]
