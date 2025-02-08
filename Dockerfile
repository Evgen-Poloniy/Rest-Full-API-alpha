FROM golang:1.22.2 AS builder

WORKDIR /app

RUN apt-get update && apt-get install -y ca-certificates

COPY libs/docker ./
RUN go mod download

COPY src/docker ./src/
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o Go-API ./src

FROM scratch

COPY --from=builder /app/Go-API /Go-API

EXPOSE 8080

CMD ["./Go-API"]