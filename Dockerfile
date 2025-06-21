# Gunakan base image Go terbaru berbasis Alpine (ringan)
FROM golang:1.21-alpine

# Set working directory di dalam container
WORKDIR /app

# Copy semua file ke dalam container
COPY . .

# Download semua dependency dari go.mod
RUN go mod download

# Build file utama aplikasi
RUN go build -o main ./cmd/main.go

# Jalankan binary hasil build
CMD ["./main"]
