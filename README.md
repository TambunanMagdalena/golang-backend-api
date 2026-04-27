# Golang Backend Paket Data API

REST API sederhana untuk studi kasus **pembelian Paket Data (Telkomsel-like)**.
Dibangun menggunakan **Golang**, **Echo**, **GORM**, dan **PostgreSQL** dengan pendekatan **Clean Architecture**.

---

##  Tech Stack

* Golang
* Echo
* GORM
* PostgreSQL
* Swagger (API Documentation)

---

##  Fitur

### User

* Create User
* Get All Users
* Get User by ID
* Update User
* Delete User

###  Paket Data

* Create Paket Data
* Get All Paket Data
* Get Paket Data by ID
* Update Paket Data
* Delete Paket Data

### Transaksi

* Create Transaksi (User membeli paket data)
* Get All Transaksi
* Get Transaksi by ID

---

##  Relasi Data

* 1 User ➝ banyak Transaksi
* 1 Paket Data ➝ banyak Transaksi

 Harga paket disimpan sebagai **snapshot saat transaksi**, sehingga tidak terpengaruh perubahan harga di masa depan.

---

## Cara Menjalankan Project

### 1. Clone Repository

```bash
git clone https://github.com/TambunanMagdalena/golang-backend-api.git
cd golang-backend-api
```

---

### 2. Setup Environment

Copy file:

```bash
cp cmd/.env.example .env
```

Lalu sesuaikan dengan database lokal:

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=yourpassword
DB_NAME=paket_data
```

---

### 3. Setup Database PostgreSQL

Pastikan PostgreSQL sudah berjalan, lalu buat database:

```sql
CREATE DATABASE paket_data;
```

---

### 4. Jalankan Aplikasi

```bash
go run cmd/main.go
```

Server akan berjalan di:

```text
http://localhost:3000
```

---

## API Documentation (Swagger)

Swagger tersedia di:

```text
http://localhost:3000/v1/swagger/index.html
```
 Bisa langsung digunakan untuk testing endpoint tanpa Postman.

---

## Contoh Endpoint

### Create User

```http
POST /v1/users
```

Body:

```json
{
  "name": "Magdalena",
  "phone_number": "08123456789"
}
```

---

### Create Paket Data

```http
POST /v1/paket-data
```

```json
{
  "name": "Paket 10GB",
  "price": 50000,
  "quota": 10,
  "active_period": 30
}
```

---

### Create Transaction

```http
POST /v1/transactions
```

```json
{
  "user_id": 1,
  "paket_data_id": 1
}
```

---

## rsitektur

Project ini menggunakan pendekatan **Clean Architecture**:

```
controllers → usecases → repositories → database
```

---


##  Author

Magdalena Pebrianty Tambunan
