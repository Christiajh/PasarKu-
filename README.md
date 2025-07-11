<div align="center">

# **🧠 SkillShare API**

*Platform backend sederhana namun powerful untuk berbagi keterampilan dan kursus singkat.  
Dirancang agar pengembang dapat dengan mudah membangun sistem pelatihan online atau aplikasi komunitas pembelajaran interaktif.*

[![last commit](https://img.shields.io/badge/last%20commit-today-brightgreen)](#)
[![Go](https://img.shields.io/badge/Go-1.20-blue?logo=go)](#)
[![PostgreSQL](https://img.shields.io/badge/PostgreSQL-Relational%20DB-blue?logo=postgresql)](#)

*Built with the tools and technologies:*

![Go](https://img.shields.io/badge/Go-00ADD8?logo=go&logoColor=white)
![Echo](https://img.shields.io/badge/Echo-132947?logo=go&logoColor=white)
![GORM](https://img.shields.io/badge/GORM-FF6B6B?logo=go&logoColor=white)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-336791?logo=postgresql&logoColor=white)
![JWT](https://img.shields.io/badge/JWT-000000?logo=json-web-tokens&logoColor=white)
![Godotenv](https://img.shields.io/badge/Godotenv-9cf?logo=go&logoColor=black)
</div>
---

## 🎯 Fitur Unggulan

- 🔐 **Registrasi & Login** dengan otentikasi JWT.
- 📚 **Manajemen Kelas (CRUD):** Buat, ubah, hapus, dan tampilkan kelas keterampilan.
- 👤 **Pendaftaran Kelas:** Pengguna dapat mendaftar atau membatalkan kelas.
- 🔒 **JWT Protected Routes** dengan Bearer Token.
- 🗃️ **Database Relasional** dengan GORM + PostgreSQL.
- 📽️ **Video Presentasi:**  
  👉 Klik gambar di bawah untuk menonton demo lengkap aplikasi SkillShare API:  
  [![Tonton Video Presentasi](https://img.shields.io/badge/Tonton%20Video%20Presentasi-FF0000?logo=youtube&logoColor=white)](#)

---

## 📽️ Video Presentasi

Klik gambar di bawah untuk menonton demo lengkap aplikasi SkillShare API:

[![Tonton Video Presentasi](https://img.youtube.com/vi/KsjL2NAwx3U/0.jpg)](https://youtu.be/KsjL2NAwx3U)

---


---

## 🧭 Struktur API (Endpoint Utama)

| Metode | Endpoint                  | Deskripsi                             | Proteksi |
|--------|---------------------------|---------------------------------------|----------|
| POST   | `/register`               | Daftarkan pengguna baru               | ❌       |
| POST   | `/login`                  | Login dan dapatkan token JWT          | ❌       |
| GET    | `/skills`                 | Ambil daftar kelas                    | ✅       |
| POST   | `/skills`                 | Buat kelas baru                       | ✅       |
| PUT    | `/skills/:id`             | Perbarui kelas                        | ✅       |
| DELETE | `/skills/:id`             | Hapus kelas tertentu                  | ✅       |
| POST   | `/skills/:id/enroll`      | Daftar ke kelas                       | ✅       |
| POST   | `/skills/:id/unenroll`    | Batalkan pendaftaran                  | ✅       |

---

## 📸 Galeri API & Routing

### 🔐 Otentikasi

**POST /register**
![Register](./images/post%20register.png)

**POST /login**
![Login](./images/post%20login.png)

---

### 📚 Manajemen Kelas

**POST /skills**  
![Post Kelas](./images/post%20classes.png)

**PUT /skills/:id**  
![Put Kelas](./images/put%20classes.png)

**DELETE /skills/:id**  
![Delete Kelas](./images/delete%20classes_3.png)

---

### 🛡️ Otentikasi Bearer Token

Gunakan token JWT sebagai Bearer Token di header authorization.
![Bearer Token](./images/bearer%20token%20put.png)

---

### 📌 Struktur Routing & Main

**Routes File**
![Routes](./images/routes.png)

**main.go**
![Main](./images/main%20nya.png)

---
