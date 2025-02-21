# Etapa de build
FROM golang:1.23.5 AS builder
WORKDIR /app
COPY go.mod .

RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -a -installsuffix cgo -o app .


FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/app .
ENTRYPOINT ["./app"]