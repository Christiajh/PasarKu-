# Gunakan image Go ringan
FROM golang:1.21-alpine

# Buat direktori kerja
WORKDIR /app

# Copy dependency dan download
COPY go.mod go.sum ./
RUN go mod download

# Copy seluruh project
COPY . .

# Build binary dengan nama main
RUN go build -o main ./cmd/main.go

# Jalankan aplikasi
CMD ["./main"]
