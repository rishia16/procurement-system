# 🚀 Procurement System (Sistem Pengadaan Barang)
Aplikasi web sederhana untuk mencatat pembelian barang (Procurement) dari Supplier. Terdiri dari Backend (Go + Fiber + MySQL) dan Frontend (HTML + jQuery + Bootstrap).

## 📌Fitur Utama Backend

| Fitur | Keterangan |
|-------|-----------|
| User Authentication | Register & Login (JWT Token) |
| Master Data | CRUD Items & CRUD Suppliers |
| Purchasing Transaction | Buat pembelian baru, hitung subtotal & grand total, update stok otomatis |
| Bonus | Transaction ACID (rollback jika gagal), Webhook async notifikasi |

## 📌Fitur Utama Frontend

- Login Page
- Dashboard & Inventory
- Create Purchase Page (keranjang + submit order)
- Reusable AJAX & Event Delegation
- SweetAlert2 untuk notifikasi/error handling

## 📂 PROCUREMENT-SYSTEM
- 📁 config
  - 📄 database.go — Koneksi DB & migrasi otomatis
- 📁 handlers
  - 📄 auth_handler.go — Register & Login user, JWT token
  - 📄 item_handler.go — CRUD Item
  - 📄 purchasing_handler.go — Proses transaksi pembelian
  - 📄 supplier_handler.go — CRUD Supplier
- 📁 middleware
  - 📄 jwt_middleware.go — Proteksi endpoint dengan JWT
- 📁 models
  - 📄 item.go — Struktur Item
  - 📄 purchasing_detail.go — Detail transaksi
  - 📄 purchasing.go — Header transaksi
  - 📄 supplier.go — Struktur Supplier
  - 📄 user.go — Struktur User
- 📁 routes
  - 📄 routes.go — Definisi route API & middleware
- 📁 utils
  - 📄 jwt.go — Fungsi generate JWT
- 📁 procurement-frontend
  - 📁 css
    - 📄 style.css — Styling halaman
  - 📁 js
    - 📄 api.js — Wrapper AJAX + token handling
    - 📄 auth.js — Login page JS
    - 📄 dashboard.js — Dashboard page JS
    - 📄 purchase.js — Halaman create purchase
  - 📄 dashboard.html — Tabel inventory & link create purchase
  - 📄 login.html — Halaman login
  - 📄 purchase.html — Halaman create purchase
- 📄 .env — Konfigurasi environment (DB, JWT secret)
- 📄 go.mod — Modul Go
- 📄 go.sum — Modul Go
- 📄 main.go — Entry point aplikasi
- 📄 Local_Procurement_Env.postman_environment.json
- 📄 Simple_Procurement_System.postman_collection.json

# ⚙️ Persiapan Backend
  1. Pastikan sudah menginstall:
     - Go >= 1.20
     - MySQL / PostgreSQL
     - Git
  2. Clone repository backend:
     ```bash
     git clone <repo-backend-url>
  3. masuk ke directory procurement-system:
     ```bash
     cd procurement-system
  4. Install dependencies Go:
     ```bash
     go mod tidy
  5. Buat database MySQL:
     ```bash
     CREATE DATABASE procurement;
  6. Buat /sesuaikan file .env di folder backend:
     ```bash
     DB_USER=root
     DB_PASS=
     DB_HOST=localhost
     DB_PORT=3306
     DB_NAME=procurement
     JWT_SECRET=super_rahasia
     #- Sesuaikan DB_USER, DB_PASS, dll. dengan setup MySQL masing-masing -#
  7. Jalankan backend:
     ```bash
     go run main.go
     #- Server akan berjalan di http://localhost:3000 -#
     
# ⚙️ Persiapan Frontend
  1. didalam directory procurement arahkan masuk menggunakan Visual Studio Code kedalam directory procurement-frontend
  2. Tidak perlu install apapun (semua library via CDN).
  3. Buka login.html di VS Code menggunakan Live Server / browser untuk memulai.

# 📝 Cara Menggunakan Aplikasi Frontend
- Register & Login
  - Akses login.html
  - Masukkan username & password
  - Klik Login → token disimpan otomatis di localStorage
- Dashboard
  - Tampilkan daftar items & stok
  - Link ke halaman Create Purchase
- Create Purchase
  - Pilih Supplier
  - Pilih Item & input Qty, klik Tambah
  - Item masuk ke keranjang tabel
  - Klik Submit Order → backend otomatis hitung subtotal & grand total, update stok
  - Notifikasi sukses/error muncul via SweetAlert2

# 🛠️ Endpoint API Backend
| Method | Endpoint         | Keterangan                   | Protected |
| ------ | ---------------- | ---------------------------- | --------- |
| POST   | /register        | Register user                | ❌         |
| POST   | /login           | Login user, return JWT token | ❌         |
| GET    | /api/items       | List items                   | ✅         |
| POST   | /api/items       | Tambah item                  | ✅         |
| GET    | /api/suppliers   | List suppliers               | ✅         |
| POST   | /api/suppliers   | Tambah supplier              | ✅         |
| POST   | /api/purchasings | Buat purchase baru           | ✅         |

> Catatan: Semua endpoint `/api/...` protected → membutuhkan token JWT.

# 📊 Rubrik Penilaian
  Backend
  - Setup, Auth, Schema, Logic, Security
  - Bonus: DB Transaction, Webhook
  
  Frontend
  - Tech (jQuery & CDN)
  - Auth (Login + Token)
  - Fetch Data, DOM Manipulation, Submit JSON
  - Bonus: Event Delegation, Reusable AJAX

# 📝 Catatan Teknis Tambahan
  > Cara reset database: hapus tabel & buat ulang sesuai schema
  > Test API menggunakan Postman (import Simple_Procurement_System.postman_collection.json)
  > Port default backend: http://localhost:3000 (atau menggunakan import environtment Local_Procurement_Env.postman_environment.json)
