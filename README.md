RESTful API dengan Golang

Deskripsi Proyek

Proyek ini adalah implementasi RESTful API menggunakan bahasa pemrograman Golang. API ini menyediakan endpoint "users" untuk mengelola data pengguna. Selain itu, API ini juga memiliki fitur untuk memeriksa API key yang digunakan dalam setiap permintaan, serta menggunakan logger untuk mencatat kegiatan dan informasi penting.

Fitur API
1. Endpoint "users", Endpoint ini digunakan untuk mengelola data pengguna. Berikut adalah daftar metode HTTP yang didukung oleh endpoint "users": GET /users: Mengembalikan daftar pengguna yang ada dalam database. Setiap permintaan ke endpoint "users" akan melewati pemeriksaan API key untuk otorisasi.

2. Pemeriksaan API Key, Setiap permintaan ke API ini harus menyertakan API key yang valid untuk otorisasi. API key ini akan diverifikasi sebelum permintaan diproses. Jika API key tidak valid, permintaan akan ditolak dengan kode status 401 Unauthorized.

3. Logger, API ini menggunakan logger untuk mencatat kegiatan dan informasi penting. Setiap permintaan dan tanggapan yang diterima oleh API akan dicatat bersama dengan informasi terkait seperti waktu, metode HTTP, endpoint, kode status, dll. Log ini dapat digunakan untuk melacak kegiatan sistem, pemecahan masalah, atau analisis kinerja.
