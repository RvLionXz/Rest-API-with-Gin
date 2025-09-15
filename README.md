# Proyek Latihan: REST API Task Manager dengan Go & Gin

Dokumentasi ini adalah catatan dari proses belajar membangun sebuah REST API sederhana dari nol menggunakan bahasa Go dan framework Gin. Tujuannya bukan untuk membuat aplikasi production-ready, melainkan untuk memahami konsep-konsep dasar backend seperti arsitektur berlapis, routing, dan interaksi database.

## Filosofi & Struktur Proyek

Proyek ini dibangun menggunakan arsitektur berlapis sederhana untuk memisahkan tanggung jawab (*Separation of Concerns*). Analogi yang kita gunakan selama belajar adalah sebuah **restoran**:

- **Router (`router/`)**: Pelayan yang menerima pesanan dari pelanggan.
- **Controller (`internal/controller/`)**: Manajer Dapur yang memvalidasi pesanan dan mendelegasikannya.
- **Repository (`internal/repository/`)**: Koki Spesialis yang hanya tahu cara berinteraksi dengan gudang (database).
- **Model (`internal/model/`)**: Kartu Resep standar yang digunakan di seluruh dapur.
- **Database (MySQL)**: Gudang bahan makanan.

Struktur folder proyek ini adalah sebagai berikut:

```
.
├── cmd/api/main.go         # Pintu masuk utama aplikasi
├── config/database.go      # Konfigurasi koneksi database
├── internal/
│   ├── controller/         # Logika untuk handle request & response
│   ├── model/              # Definisi struct data (Task)
│   └── repository/         # Logika untuk query ke database
├── router/router.go        # Pengaturan semua rute/endpoint API
├── go.mod
├── go.sum
└── README.md
```

## Setup & Instalasi

1.  **Prasyarat**:
    *   Go (versi 1.18+).
    *   MySQL Server.

2.  **Database Setup**:
    Masuk ke MySQL dan jalankan perintah SQL berikut untuk membuat database dan tabel yang dibutuhkan.

    ```sql
    CREATE DATABASE task_manager_db;

    USE task_manager_db;

    CREATE TABLE tasks (
        id INT PRIMARY KEY AUTO_INCREMENT,
        title VARCHAR(255) NOT NULL,
        description TEXT,
        status VARCHAR(50) NOT NULL DEFAULT 'pending',
        created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
    );
    ```

3.  **Konfigurasi Aplikasi**:
    Buka file `config/database.go` dan sesuaikan `username` dan `password` MySQL Anda.

4.  **Instalasi Dependensi**:
    Buka terminal di root direktori proyek dan jalankan:
    ```bash
    go mod tidy
    ```

5.  **Menjalankan Server**:
    ```bash
    go run cmd/api/main.go
    ```
    Server akan berjalan di `http://localhost:8080`.

## Dokumentasi API Endpoint

### 1. Membuat Task Baru

- **Endpoint**: `POST /tasks`
- **Deskripsi**: Menambahkan sebuah task baru ke dalam database.
- **Request Body**:
  ```json
  {
      "title": "Belajar Dokumentasi API",
      "description": "Menulis file README.md yang baik.",
      "status": "pending"
  }
  ```
- **Success Response (`201 Created`)**:
  ```json
  {
      "id": 1,
      "title": "Belajar Dokumentasi API",
      "description": "Menulis file README.md yang baik.",
      "status": "pending",
      "created_at": "2025-09-16T15:00:00Z",
      "updated_at": "2025-09-16T15:00:00Z"
  }
  ```
- **Contoh `curl`**:
  ```bash
  curl -X POST http://localhost:8080/tasks \
  -H "Content-Type: application/json" \
  -d '{"title": "Belajar Dokumentasi API", "description": "Menulis file README.md"}'
  ```

### 2. Mengambil Semua Task

- **Endpoint**: `GET /tasks`
- **Deskripsi**: Mengambil daftar semua task yang ada di database.
- **Success Response (`200 OK`)**:
  ```json
  [
      {
          "id": 1,
          "title": "Belajar Dokumentasi API",
          ...
      },
      {
          "id": 2,
          "title": "Mencoba endpoint GET",
          ...
      }
  ]
  ```
- **Contoh `curl`**:
  ```bash
  curl http://localhost:8080/tasks
  ```

### 3. Mengambil Satu Task

- **Endpoint**: `GET /tasks/:id`
- **Deskripsi**: Mengambil satu task spesifik berdasarkan ID-nya.
- **Success Response (`200 OK`)**:
  ```json
  {
      "id": 1,
      "title": "Belajar Dokumentasi API",
      ...
  }
  ```
- **Error Response (`404 Not Found`)**:
  ```json
  {
      "error": "Task tidak ditemukan"
  }
  ```
- **Contoh `curl`**:
  ```bash
  curl http://localhost:8080/tasks/1
  ```

### 4. Memperbarui Task

- **Endpoint**: `PUT /tasks/:id`
- **Deskripsi**: Memperbarui data sebuah task berdasarkan ID.
- **Request Body**:
  ```json
  {
      "title": "Belajar Dokumentasi API (Updated)",
      "description": "Menulis file README.md yang baik dan lengkap.",
      "status": "completed"
  }
  ```
- **Success Response (`200 OK`)**:
  ```json
  {
      "id": 1,
      "title": "Belajar Dokumentasi API (Updated)",
      ...
  }
  ```
- **Contoh `curl`**:
  ```bash
  curl -X PUT http://localhost:8080/tasks/1 \
  -H "Content-Type: application/json" \
  -d '{"title": "Judul Baru", "description": "Deskripsi Baru", "status": "completed"}'
  ```

### 5. Menghapus Task

- **Endpoint**: `DELETE /tasks/:id`
- **Deskripsi**: Menghapus sebuah task berdasarkan ID-nya.
- **Success Response (`200 OK`)**:
  ```json
  {
      "message": "Task berhasil dihapus"
  }
  ```
- **Contoh `curl`**:
  ```bash
  curl -X DELETE http://localhost:8080/tasks/1
  ```

