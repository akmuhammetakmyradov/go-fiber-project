# Build stage
FROM golang:1.19-buster AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o testmain ./cmd/main.go


# Deploy stage
FROM debian:buster-slim

WORKDIR /app

COPY --from=builder /app/testmain /app/config.yml ./

EXPOSE 3000

CMD ["./testmain"]