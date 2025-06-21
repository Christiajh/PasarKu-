# Gunakan image Go berbasis Alpine
FROM golang:1.21-alpine

# Buat direktori kerja dalam container
WORKDIR /app

# Copy file dependency terlebih dahulu
COPY go.mod go.sum ./
RUN go mod download

# Copy semua source code ke container
COPY . .

# Build binary aplikasi
RUN go build -o main ./cmd/main.go

# Jalankan aplikasi
CMD ["./main"]
