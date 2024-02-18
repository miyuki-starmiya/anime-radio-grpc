FROM golang:1.21 as builder

# Set an environment variable to ensure the binary is built with Go modules enabled.
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

# Build the Go app as a static binary.
RUN go build -a -installsuffix cgo -o main .

FROM alpine:latest

ENV ENV=production

WORKDIR /app

COPY --from=builder /app/main .

CMD ["./main"]
