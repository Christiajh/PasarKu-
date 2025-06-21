# Gunakan image Go ringan
FROM golang:1.21-alpine

# Set working directory
WORKDIR /app

# Install git (dibutuhkan untuk go mod jika pakai private repo)
RUN apk add --no-cache git

# Copy go mod dan sum, lalu download dependency
COPY go.mod go.sum ./
RUN go mod download

# Copy semua kode ke dalam container
COPY . .

# Build binary dari cmd/main.go dan hasilkan ke root (/app/main)
RUN go build -o main ./cmd/main.go

# Berikan permission execute ke binary
RUN chmod +x ./main

# Jalankan aplikasi
CMD ["./main"]
