# Procurement System (Sistem Pengadaan Barang)
Aplikasi web sederhana untuk mencatat pembelian barang (Procurement) dari Supplier. Terdiri dari Backend (Go + Fiber + MySQL) dan Frontend (HTML + jQuery + Bootstrap).

## Fitur Utama Backend

| Fitur | Keterangan |
|-------|-----------|
| User Authentication | Register & Login (JWT Token) |
| Master Data | CRUD Items & CRUD Suppliers |
| Purchasing Transaction | Buat pembelian baru, hitung subtotal & grand total, update stok otomatis |
| Bonus | Transaction ACID (rollback jika gagal), Webhook async notifikasi |

## Frontend

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

⚙️ Persiapan Backend
1. Pastikan sudah menginstall:
   - Go >= 1.20
   - MySQL / PostgreSQL
   - Git

2. Clone repository backend:
   
