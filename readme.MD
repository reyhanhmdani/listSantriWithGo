Autentikasi Pengguna:

 - Pengguna membuka aplikasi dan diminta untuk masuk atau mendaftar.
Setelah masuk, pengguna menerima token otentikasi yang menyertakan ID pengguna dan peran pengguna (admin atau user).
Pengguna Admin:

- Seorang pengguna admin dapat melakukan tindakan berikut:
Mengelola Data Santri:

- Menambahkan data santri baru (nama, jenis kelamin, angkatan, jurusan, minat, status).
Melihat daftar semua santri.
Mencari dan menampilkan daftar santri berdasarkan berbagai kriteria seperti nama, jurusan, atau minat.
Mengedit data santri (termasuk mengubah nama, jurusan, minat, dan status).
Menghapus data santri.
Mengelola Pengguna:

- Melihat daftar pengguna (admin dan pengguna).
Mengedit data pengguna (mengubah kata sandi, jika diperlukan).
Menghapus pengguna (termasuk pengguna admin lain).
Pengguna Biasa (User):

- Seorang pengguna biasa memiliki kemampuan terbatas:
Melihat daftar santri.
Mencari dan menampilkan daftar santri berdasarkan berbagai kriteria seperti nama, jurusan, atau minat.
Mengedit profil mereka sendiri (termasuk mengubah kata sandi jika diperlukan).
Menggunakan API:

- Aplikasi klien (web atau perangkat lunak lainnya) berinteraksi dengan API server melalui permintaan HTTP seperti GET, POST, PUT, dan DELETE.
Permintaan HTTP diarahkan ke endpoint yang sesuai, misalnya, untuk menambahkan santri baru, permintaan POST dikirim ke /addSantri, dan untuk mencari santri, permintaan GET dikirim ke /search.
API server memvalidasi permintaan berdasarkan token otentikasi yang disertakan.
Pesan Kesalahan:

- Server mengirim pesan kesalahan jika terjadi kesalahan dalam permintaan, seperti kesalahan otentikasi, kesalahan validasi data, atau kesalahan server.
Respon:

- Server mengirim respon yang sesuai, termasuk data yang diminta atau pesan kesalahan jika diperlukan.
Aplikasi klien menerima respons dan menampilkan hasilnya kepada pengguna.
Aksi Pengguna:

- Pengguna (admin atau user) dapat melakukan berbagai aksi seperti menambahkan data, mengedit data, mencari data, atau menghapus data berdasarkan wewenang dan fungsionalitas yang diberikan oleh peran mereka.
Pengelolaan Kesalahan:

- Aplikasi dan server mengelola berbagai jenis kesalahan yang mungkin terjadi, termasuk kesalahan validasi data, kesalahan otentikasi, dan kesalahan server.
Pengamanan:

- Keamanan diimplementasikan melalui otentikasi pengguna, manajemen izin, dan penggunaan token otentikasi.
Logging:

- Tambahan : - adanya fitur forgot password (masih pengembangan. (sudah jadi tapi masih testing))
 
- Aplikasi dan server bisa melakukan logging aktivitas seperti permintaan masuk, akses ke data, dan perubahan data untuk keperluan pemantauan dan penelusuran.
