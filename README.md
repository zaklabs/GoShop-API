# Evermos API - RESTful API with Golang
Backend API untuk platform marketplace Evermos, dibangun dengan Golang menggunakan Clean Architecture pattern.

## 🚀 Fitur
- **Authentication & Authorization**: JWT-based authentication dengan role-based access control (Admin/User)
- **User Management**: Register, login, profile management
- **Toko Management**: CRUD toko dengan file upload untuk foto
- **Product Management**: CRUD produk dengan multiple foto upload, filtering, dan pagination
- **Category Management**: CRUD kategori (Admin only)
- **Address Management**: CRUD alamat pengiriman
- **Transaction System**: 
  - Create transaksi dengan multiple items
  - Auto-generate invoice code
  - Product snapshot (log_produk) untuk historical data
  - Stock management
- **Smart Delete System (Soft Delete)**: 
  - Intelligent product deletion dengan validasi transaksi
  - Soft delete untuk produk yang sudah memiliki riwayat transaksi (data preservation)
  - Hard delete untuk produk tanpa transaksi (data cleanup)
  - Automatic filtering untuk produk yang di-delete dari semua query
  - Validasi transaksi untuk mencegah order produk yang sudah dihapus
  - Menjaga integritas data historis untuk keperluan audit dan pelaporan
- **Security**:
  - Password hashing dengan bcrypt
  - JWT token authentication
  - Ownership validation
  - Admin-only endpoints

## 📁 Struktur Project
```
evermos-api/
├── cmd/
│   └── api/
│       └── main.go                 # Entry point aplikasi
├── internal/
│   ├── config/
│   │   ├── config.go              # Configuration loader
│   │   └── database.go            # Database connection & migration
│   ├── delivery/
│   │   ├── http/
│   │   │   ├── handler/           # HTTP handlers/controllers
│   │   │   └── router.go          # Route definitions
│   │   └── middleware/            # Middleware (auth, logger, CORS)
│   ├── model/                     # Models & DTOs
│   ├── repository/                # Data access layer
│   ├── usecase/                   # Business logic layer
│   └── utils/                     # Helper utilities
├── uploads/                       # File upload directory
├── .env.example                   # Environment variables template
├── go.mod                         # Go modules
└── README.md
```

## 🛠️ Tech Stack
- **Language**: Go 1.25+
- **Framework**: Gin (HTTP router)
- **ORM**: GORM
- **Database**: MySQL 8.0+
- **Authentication**: JWT (golang-jwt/jwt v5)
- **Validation**: go-playground/validator v10
- **Password**: bcrypt
- **Testing**: testify

## � Prerequisites
- Go 1.21 atau lebih tinggi
- MySQL 8.0 atau lebih tinggi
- Git

## 🔧 Installation & Setup

### 1. Clone Repository
```bash
cd c:\Users\NOC-01\rekamin\final-projek
```

### 2. Setup Database
Buat database MySQL:

```sql
CREATE DATABASE evermos CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;
```

### 3. Configure Environment
Copy `.env.example` ke `.env` dan sesuaikan konfigurasi:

```bash
copy .env.example .env
```

Edit file `.env`:

```env
# Database Configuration
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=your_password_here
DB_NAME=evermos

# JWT Configuration
JWT_SECRET=your-super-secret-jwt-key-change-this-in-production
JWT_EXPIRE_HOURS=24

# Server Configuration
SERVER_PORT=8000

# Upload Configuration
UPLOAD_PATH=./uploads
MAX_UPLOAD_SIZE=5242880
```

### 4. Install Dependencies
```bash
go mod download
go mod tidy
```

### 5. Run Database Migration
Migration akan otomatis dijalankan saat aplikasi start. Atau jalankan manual:

```bash
go run cmd/api/main.go
```

### 6. Start Server
```bash
go run cmd/api/main.go
```

Server akan berjalan di `http://localhost:8000`

## 📚 API Documentation
Public URL: `[https://documenter.getpostman.com/view/14392918/2sB3Wk14Lh](https://documenter.getpostman.com/view/14392918/2sB3Wk14Lh)`

## 🔒 Security Features

### Ownership Validation
Sistem memastikan user hanya bisa mengakses/mengubah data milik sendiri:
- **User Profile**: User hanya bisa lihat/update profile sendiri
- **Alamat**: User hanya bisa CRUD alamat milik sendiri
- **Toko**: User hanya bisa update toko milik sendiri
- **Produk**: User hanya bisa CRUD produk dari toko milik sendiri
- **Transaksi**: User hanya bisa lihat transaksi sendiri
- **Kategori**: Hanya role admin yang dapat menambah kategori

## 🧪 Testing
Run unit tests:
```bash
go test ./tests/integration/... -v
```

## 📝 Response Format
Semua response menggunakan format unified:

**Success Response:**
```json
{
  "status": true,
  "message": "Succeed to GET data",
  "errors": null,
  "data": { ... }
}
```

**Error Response:**
```json
{
  "status": false,
  "message": "Failed to POST data",
  "errors": ["error message here"],
  "data": null
}
```

## 🗄️ Database Schema
Database schema mengikuti file draw.io yang diberikan, dengan tabel:
- `users` - User accounts
- `toko` - Stores
- `alamat` - Shipping addresses
- `category` - Product categories
- `produk` - Products
- `foto_produk` - Product photos
- `log_produk` - Product snapshots (transaction history)
- `trx` - Transactions
- `detail_trx` - Transaction details

## 🚦 Development
### Build for production
```bash
go build -o evermos-api.exe cmd/api/main.go
```

### Run in production
```bash
./evermos-api.exe
```

## 📄 License
MIT License

## 👤 Author
Zaki Fuadi
Rakamin x Evermos Virtual Internship Project

---

**Happy Coding! 🎉**

