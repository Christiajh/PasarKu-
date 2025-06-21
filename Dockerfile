# Gunakan image ringan Go berbasis Alpine
FROM golang:1.21-alpine

# Buat working directory di dalam container
WORKDIR /app

# Copy file go.mod dan go.sum dulu, agar cache download efektif
COPY go.mod go.sum ./

# Download dependency dulu (lebih cepat, terpisah dari source code)
RUN go mod download

# Copy semua source code setelah download selesai
COPY . .

# Build aplikasi
RUN go build -o main ./cmd/main.go

# Jalankan aplikasi
CMD ["./main"]
