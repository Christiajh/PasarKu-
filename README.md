# PasarKu – Platform Digital UMKM Lokal Indonesia

PasarKu adalah RESTful API berbasis Golang yang dirancang untuk mendukung digitalisasi UMKM lokal di Indonesia. Platform ini memudahkan penjual mikro dari berbagai daerah untuk menjual, mempromosikan, dan melacak produk secara digital, sekaligus menyediakan sistem ulasan dan pelacakan distribusi produk berbasis lokasi.

## 🎯 Latar Belakang

UMKM merupakan tulang punggung ekonomi Indonesia, namun masih banyak yang belum terdigitalisasi. Beberapa permasalahan utama:
- Minim platform yang mendukung pasar mikro/pedesaan.
- Sulitnya distribusi dan promosi produk lokal.
- Tidak adanya sistem review yang memadai.
- Tidak bisa melacak atau memonitor sebaran produk secara efektif.

## 🧠 Fitur Utama

- 🔐 **User Management**  
  Registrasi, login, otentikasi JWT.

- 🛍️ **Product Management**  
  CRUD produk UMKM, upload gambar.

- 📦 **Order Management**  
  Transaksi produk, status pemesanan, dan pembayaran.

- ⭐ **Review System**  
  Pembeli dapat memberikan ulasan pada produk.

- 🎯 **Kategori UMKM**  
  Mendukung kategori seperti makanan, kerajinan, pakaian, dll.

- 🏷️ **Tag Lokal**  
  Produk dapat ditandai dengan wilayah tertentu (contoh: `#Bandung`, `#Medan`).

- 📍 **Geo Check-in Produk**  
  Pelacakan distribusi produk berdasarkan lokasi.

- 🧑‍🤝‍🧑 **Role Management**  
  Mendukung peran: User, Seller, dan Admin.

---

## 🧱 Struktur Direktori

pasarku/
├── controller/ # Handler endpoint (User, Product, Order, Review)
├── database/ # Setup dan migrasi DB
├── helper/ # JWT, hash password, validasi, response handler
├── middleware/ # Auth, role check
├── model/ # Model & logic bisnis (repository)
├── public/ # File upload (foto produk)
├── routes/ # Definisi semua route
├── structs/ # DTO request dan response
├── docs/ # Swagger documentation
├── main.go # Entry point aplikasi
├── go.mod # File dependensi Go
└── dbconfig.yml # Konfigurasi database

## 🛠️ Teknologi & Library

| Teknologi        | Deskripsi                             |
|------------------|----------------------------------------|
| [Golang](https://golang.org)         | Bahasa utama untuk backend API         |
| [Gin](https://github.com/gin-gonic/gin) | Web framework ringan dan cepat          |
| [GORM](https://gorm.io)              | ORM untuk PostgreSQL                    |
| [PostgreSQL](https://www.postgresql.org) | Basis data relasional                   |
| [JWT-Go](https://github.com/golang-jwt/jwt) | Otentikasi token berbasis JWT         |
| [Bcrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt) | Hashing password                       |
| [Viper](https://github.com/spf13/viper) | Manajemen konfigurasi environment      |
| [sql-migrate](https://github.com/rubenv/sql-migrate) | Migrasi skema database                 |
| [Swag](https://github.com/swaggo/swag) | Dokumentasi Swagger API                |


🧑‍💻 Developer
Christian Johannes Hutahaean
