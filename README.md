# 📌 OttoDigital Project

Untuk memenuhi proses recruitment perusahaan OttoDigital

## ✨ Fitur
- ✅ Fitur 1: Backend With Go
- ✅ Fitur 2: Frontend With React

## 🚀 Teknologi yang Digunakan
- **Frontend**: React.js
- **Backend**: Go
- **Database**: MySQL
- **Tools lain**: XAMPP (Atau program Apache Local lainnya)

---

## 📂 Instruksi Penggunaan

### PULL Semua file dari reporsitory ini

`gh repo clone raihanpraditya/ottodigital`

Atau

[Klik Link Ini!](https://github.com/raihanpraditya/ottodigital.git)

### BACKEND
 - Sesuaikan credential dari database anda.
`Backend/database/connections.go` Pada line 12
`dsn := "root:root@tcp(localhost:3306)/voucher_api?parseTime=true"`
root pertama merupakan username database
root kedua merupakan password database
voucher_api merupakan nama database

- Lakukan migrasi dengan mengimport file sql pada
`Backend/migrations/20241203_init_schema.up.sql`

- Run `main.go` untuk menjalankan server
dengan menggunakan `go run main.go` pada terminal anda pada `http://localhost:8080`

- Anda dapat mulai mencoba API dengan aplikasi seperti postman

- **List Endpoint**
    - `/brand` Dengan metode POST untuk melakukan input brand baru
        Atur body dengan metode RAW JSON
            `{"name": "Brand"}`
    
    - `/voucher` Dengan metode POST untuk melakukan input voucher baru
        Atur body dengan metode RAW JSON
            `{"brand_id": 1,"name": "Voucher A","cost_in_point": 5000}`
    
    - `/voucher?id=1` Dengan metode GET untuk mendapatkan single voucher tertentu

    - `/voucher/brand?id=1` Dengan metode GET untuk mendapatkan voucher dari brand tertentu

    - `/transaction/redemption` Dengan metode POST untuk membuat redemption voucher
        Atur body dengan metode RAW JSON
            `{"customer_name": "Raihan","total_points": 10000,"details": [{"voucher_id": 1,"quantity": 2}]}`
    
    - `/transaction/redemption?transactionId=1` Dengan metode GET untuk mendapatkan detail redamption yang sudah di input 


### FRONTEND

- Run `npm start` pada terminal anda

- Buka app di browser pada link `http://localhost:3000`

- Jika file tidak terdownload secara keseluruhan lakukan input pada terminal

    `n`


