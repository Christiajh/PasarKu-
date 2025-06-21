# Gunakan image Go ringan berbasis Alpine
FROM golang:1.21-alpine

# Buat direktori kerja di dalam container
WORKDIR /app

# Salin dependency terlebih dahulu
COPY go.mod go.sum ./
RUN go mod download

# Salin semua source code
COPY . .

# Build aplikasi
RUN go build -o main ./cmd/main.go

# Jalankan aplikasi
CMD ["./main"]
