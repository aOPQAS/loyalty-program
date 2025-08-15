FROM golang:1.23-alpine AS builder
ENV GO111MODULE=on
WORKDIR /app
COPY go.mod .
COPY go.sum .
COPY . .
ENV GOOS=linux
RUN go mod download
RUN go build -o back ./cmd/main.go

FROM alpine
WORKDIR /app
COPY --from=builder /app/back /app/back
COPY --from=builder /app/migrations /app/migrations
EXPOSE 8080
RUN chmod +x /app/back
CMD ["./back"]
