# Go Fiber

### Introduction
Backend Developer Challenge Test

Studi Kasus:
Sebuah perusahaan bergerak di bidang Outsource memiliki sebuah layanan aplikasi yang sudah berjalan, dan untuk saat ini memerlukan services tambahan terkait divisi HR untuk kebutuhan absensi semua karyawannya.

Tugas:
Buat mini Services baru menggunakan NodeJS, PHP, GO atau sejenisnya, yang terdiri dari 2 Endpoint API, diantaranya adalah:

1. Endpoint Login dan Autentikasi
2. Endpoint Clock In dan Clock Out (Endpoint harus disertakan dengan Authorisation)

*Note: Endpoint absensi perlu mencatat data IP Address beserta Latitude dan Longitude.

...

### Installation

1. Open in cmd or terminal app and navigate to this folder
2. Run following commands

```bash
go get
```

3. Import go_dev.sql on your own database. For example using phpmyadmin

---- LIST ENDPOINTS ----
1. Auth or Login: http://127.0.0.1:3000/api/auth/login
    And then on form data input
    ```bash
    username:user
    password:123456
    ```
    ```
    It will return token to put on Header Authorization on the endpoint Clock In and Clock Out
    ```
2. Clock In : http://127.0.0.1:3000/api/transaction/clock_in
3. Clock Out : http://127.0.0.1:3000/api/clock_out

And navigate to generated server link (http://127.0.0.1:3000/api/)

### Copyright

...
