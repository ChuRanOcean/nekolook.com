# Build stage
FROM golang:1.21-alpine AS builder

# Set China-friendly Go module proxy
ENV GOPROXY=https://goproxy.cn,direct \
    GOSUMDB=sum.golang.org

RUN apk add --no-cache ca-certificates git

WORKDIR /go/src/github.com/churanocean/nekolook

# Copy module files first
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build binary
RUN CGO_ENABLED=0 GOOS=linux go build -o /nekolook ./cmd/api

# Final stage
FROM alpine:3.19

WORKDIR /app

# Copy binary from builder
COPY --from=builder /nekolook .

EXPOSE 8080

CMD ["/app/nekolook"]
