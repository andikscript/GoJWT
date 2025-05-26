# GoJWT 🛡️

**GoJWT** adalah proyek contoh sederhana penerapan JSON Web Token (JWT) menggunakan bahasa pemrograman **Golang**. Proyek ini dibuat sebagai referensi dasar bagaimana mengimplementasikan autentikasi berbasis token di aplikasi berbasis REST API.

## 🔧 Fitur

- Implementasi JWT sederhana menggunakan Go.
- 3 endpoint utama:
  - **`/login`**: Menggunakan **Basic Auth** (username & password) untuk login dan mendapatkan token JWT.
  - **`/index`**: Endpoint yang dilindungi, memerlukan **Bearer Token** di header `Authorization`.
  - **`/noauth`**: Endpoint publik, tidak membutuhkan token.

## 📦 Endpoint Overview

| Endpoint | Method | Autentikasi | Deskripsi |
|----------|--------|-------------|-----------|
| `/login` | POST   | Basic Auth  | Login dan mendapatkan JWT |
| `/index` | ALL    | Bearer Token | Endpoint aman, hanya bisa diakses dengan token valid |
| `/noauth`| ALL    | Tidak perlu | Endpoint bebas tanpa autentikasi |

## 🔐 Contoh Penggunaan

### 1. Login untuk mendapatkan Token

```bash
curl -X POST http://localhost:8080/login -u your_username:your_password
