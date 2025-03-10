# Gateway Services API

API ini adalah aplikasi layanan multifinance yang dikembangkan menggunakan Golang dan PostgreSQL.

## Persiapan Sebelum Memulai

1. **Pastikan Database Sudah Tersedia**
   - Pastikan Anda sudah membuat database PostgreSQL dengan nama **`multifinance_db`**.

2. **Konfigurasi Aplikasi**
   - Buat file `config.yaml`.
   - Anda bisa melihat contoh configurasinya pada file `sample.config.yaml`.
   - Sesuaikan pengaturan koneksi database dan konfigurasi lainnya sesuai dengan lingkungan Anda.
 

## Langkah-Langkah Menjalankan Aplikasi

### 1. Menjalankan Migrations
   - Migration digunakan untuk mengatur struktur database yang diperlukan oleh aplikasi ini.
   - Masuk ke direktori `migrations/` dengan perintah berikut:
     ```bash
     cd migrations/
     ```
   - Jalankan perintah berikut untuk menjalankan migration:
     ```bash
     go run migration.go ./sql "host=localhost port=5432 user=root dbname=db_users password=password sslmode=disable" up
     ```
   - Pastikan detail koneksi (seperti host, port, user, dan password) sesuai dengan konfigurasi database PostgreSQL Anda.

### 2. Menjalankan Aplikasi
   - Setelah konfigurasi selesai, Anda dapat menjalankan aplikasi dengan perintah:
     ```bash
     go run .
     ```
   - Aplikasi sekarang akan berjalan dan terhubung ke database `db_users`.

## Struktur Direktori

- **config/**: Menyimpan konfigurasi yang berkaitan dengan layanan pihak ketiga.
- **handlers/**: Lapisan yang menangani permintaan dari pengguna, baik dari aplikasi mobile maupun web.
- **migrations/**: Menyimpan file migration SQL dan kode untuk mengatur dan memperbarui struktur database.
- **models/**: Berisi struktur data (constructs di Golang) yang memudahkan dalam membuat kontrak untuk request dan response.
- **respository/**: Lapisan yang berfungsi khusus untuk berinteraksi dengan database, termasuk operasi pencatatan dan pengambilan data.
- **routes/**: Menyimpan definisi endpoint utama yang mengarahkan permintaan ke handler terkait.
- **usecases/**: Lapisan yang menangani logika bisnis, termasuk pengolahan data dari input pengguna atau hasil dari database.
- **util/**: Menyimpan middleware dan fungsi utilitas lainnya.
- **config.yaml**: File konfigurasi utama untuk mengatur koneksi database dan parameter aplikasi lainnya.

## Endpoint
`SignUp User`\
`POST :localhost:8080/multifinance/signup`
```curl --location ':8080/multifinance/signup' \
--header 'Content-Type: application/json' \
--data '{
    "nik":"1111111111111121",
    "fullname":"rama",
    "legal_name":"ramadan rangkuti",
    "place_of_birth":"Jakarta",
    "date_of_birth":"1998-01-22",
    "salary":5000000,
    "password":"12345678",
    "role":"User"
}'

```
`Signin User`\
`POST :localhost:8080/multifinance/signup`
```curl --location ':8080/multifinance/signup' \
--header 'Content-Type: application/json' \
--data '{
    "nik":"1111111111111121",
    "password":"12345678"
}'
```

`Create transaction`\
`POST :localhost:8080/multifinance/transactions`
```curl --location ':8080/multifinance/transactions' \
--header 'Content-Type: application/json' \
--data '{
    "nik":"1111111111111121",
    "password":"12345678"
}'
```

`Create transaction`\
`Headers :`\
``` Authorization: Bearer your-jwt-token ```\
`POST :localhost:8080/multifinance/transactions`
```curl --location ':8080/multifinance/transactions' \
--header 'Content-Type: application/json' \
--data '{
    "nik":"1111111111111121",
    "password":"12345678"
}'
```

`Get transaction`\
`Headers :`\
``` Authorization: Bearer your-jwt-token ```\
`GET :localhost:8080/multifinance/transactions`
```curl --location ':8080/multifinance/transactions' \
--header 'Content-Type: application/json' 
```

## System Architecture
structure

![alt text ](https://github.com/RamadanRangkuti/multifinance-api/development/structure.png?raw=true)

<br />

<br />
Database Diagram


## Teknologi yang Digunakan

- **Golang**: Backend aplikasi utama.
- **PostgreSQL**: Database untuk menyimpan data pengguna.

#### Dokumentasi ini akan terus diperbarui sesuai dengan perkembangan aplikasi. Jika ada pertanyaan, silakan hubungi tim pengembang.