# stage 1

FROM golang:alpine AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o Void-Go .

# Stage 2

FROM alpine:latest

LABEL org.opencontainers.image.source="https://github.com/Arjuna-Ragil/Void-Go"

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/Void-Go .

EXPOSE 53/udp
EXPOSE 53/tcp

CMD ["./Void-Go"]