# Portal & CMS Pondok Pesantren Al-Hamid

Sistem informasi portal resmi dan Content Management System (CMS) modern, elegan, dan berperforma tinggi untuk **Pondok Pesantren Al-Hamid**.

Arsitektur sistem dibangun dengan standar decoupled modern:
- **Backend**: Golang (Fiber v2, GORM ORM, JWT Authentication, RESTful API)
- **Frontend**: SvelteKit 5 + TypeScript + Tailwind CSS v4 + Lucide Icons
- **Database**: PostgreSQL (dilengkapi auto-fallback SQLite untuk kemudahan pengujian lokal)
- **Deployment**: Siap dideploy ke VPS menggunakan Docker & Docker Compose

---

## 🌟 Struktur Navigasi & Fitur Utama

### 1. Beranda (Home)
- **Hero Cinematic**: Slider latar belakang foto resolusi tinggi, tipografi prestisius, badge status, tombol aksi (*Call to Action*).
- **Statistik Cepat**: Counter santri aktif, dewan pengasuh/asatidz, alumni tersebar, dan prestasi.
- **Profil Singkat & Sambutan Pengasuh**: Sambutan K.H. Ahmad Hamidurrohman dengan sanad keilmuan.
- **Program Pendidikan Unggulan**: Kartu interaktif untuk jenjang formal dan non-formal.
- **Highlight Berita & Artikel**: Grid artikel kajian dan kabar pesantren terkini.
- **Prestasi & Fasilitas**: Showcase capaian kejuaraan dan sarana representatif.
- **Testimoni**: Ulasan dan cerita dari alumni serta wali santri.
- **Call to Action PSB**: Akses cepat pendaftaran santri baru via form & WhatsApp hotline.

### 2. Tentang Kami
- **Profil & Sejarah**: Riwayat berdiri, filosofi, dan Panca Jiwa Pesantren.
- **Visi & Misi**: Arah strategis dan langkah terukur lembaga.
- **Dewan Pengasuh**: Daftar masyayikh, biografi, dan silsilah sanad keilmuan muttashil.
- **Fasilitas**: Dokumentasi asrama, masjid jami', laboratorium sains, maktabah turats, dan sarana olahraga.
- **Prestasi**: Rekam jejak kejuaraan tingkat kabupaten, provinsi, hingga nasional.
- **Testimoni**: Kesan dan pesan alumni serta wali santri.
- **Ikatan Alumni (IKAH)**: Direktori kiprah lulusan di bidang akademik, birokrasi, dan industri.
- **Kontak & Sosmed**: Peta lokasi, kontak sekretariat, dan tautan media sosial.

### 3. Pendidikan (Laman Khusus per Unit)
- **Pendidikan Non-Formal**:
  1. `Madrasah Diniyah Al-Hamid`: Penguasaan gramatika bahasa Arab (Nahwu, Shorof) dan kajian kitab kuning berjenjang (Ula, Wustha, Ulya).
  2. `PMQ (Pendidikan Murottilil Qur'an)`: Pembinaan fashohah, nagham/irama 7 maqamat, serta talaqqi sanad riwayat Imam Ashim.
- **Pendidikan Formal**:
  3. `MTs Cintamulya`: Madrasah Tsanawiyah terakreditasi A unggul, integrasi kurikulum Kemenag dan kepesantrenan.
  4. `MA Cintamulya`: Madrasah Aliyah unggulan (jurusan MIPA, IPS, Keagamaan) dengan program persiapan PTN dan beasiswa Timur Tengah.
  5. `MAI Cintamulya`: Madrasah Aliyah Kejuruan Informatika dan Multimedia Dakwah Digital.

### 4. Berita & Literasi
- Filter berdasarkan kategori (Kajian Kitab, Kabar Pesantren, Prestasi, Pengumuman, Opini).
- Pencarian artikel *real-time*.
- Halaman baca artikel dengan tipografi editorial yang nyaman, meta informasi, dan tombol bagikan.

### 5. Galeri Media
- **Galeri Foto**: Pengelompokan album berdasarkan kategori dengan lightbox interaktif.
- **Galeri Video**: Integrasi resmi YouTube dengan modal player video.

### 6. Spotlight Search
- Akses cepat modal pencarian menggunakan tombol pencarian atau shortcut `Ctrl + K` / `Cmd + K`.

### 7. CMS / Admin Dashboard (Full CRUD)
Akses portal admin: `/admin/login`
- **Dashboard Ringkasan**: Statistik artikel, unit pendidikan, foto, video, pengasuh, dan pesan belum dibaca.
- **Manajemen Artikel & Kategori**: Tulis, edit, hapus, kelola slug SEO, status draft/publish, dan fitur featured.
- **Manajemen Unit Pendidikan**: Kelola kurikulum, fasilitas, syarat masuk, dan kepala unit.
- **Manajemen Dewan Pengasuh**: Kelola data masyayikh dan sanad keilmuan.
- **Manajemen Fasilitas & Prestasi**: Tambah/edit sarana dan data kejuaraan santri.
- **Manajemen Testimoni & Alumni**: Kelola direktori alumni dan testimoni.
- **Manajemen Galeri Media**: Upload foto dan sematan video YouTube.
- **Manajemen Hero Slider**: Atur urutan banner slide beranda.
- **Kotak Masuk Pesan**: Baca pesan masuk pengunjung dari form kontak, tandai sudah dibaca, atau balas via WhatsApp / Email.
- **Pengaturan Website**: Update identitas, nomor telepon, WhatsApp, email, alamat, sosial media, profil, dan visi-misi secara *real-time*.

---

## 🚀 Panduan Menjalankan di Lokal (Development)

### Prasyarat
- Go 1.22+
- Node.js 20+
- PostgreSQL (opsional untuk lokal, sistem otomatis fallback ke SQLite jika kredensial belum dikonfigurasi)

### 1. Menjalankan Backend (Go)
```bash
cd backend

# Salin file konfigurasi environment
cp .env.example .env

# Jalankan server API
go run ./cmd/server
# atau jalankan binary yang telah dicompile:
# ./server.exe
```
Backend akan berjalan di `http://localhost:8088`. Database akan otomatis ter-migrasi dan terisi data awal (*seed data*).

### 2. Menjalankan Frontend (Svelte)
```bash
cd frontend

# Install dependensi (jika belum)
npm install

# Jalankan dev server
npm run dev
```
Frontend akan berjalan di `http://localhost:5173`.

### 3. Akun Administrator Default
- **URL Login CMS**: `http://localhost:5173/admin/login`
- **Username**: `admin`
- **Kata Sandi**: `adminpassword123`

---

## 🌐 Panduan Deploy ke VPS (Production)

Proyek ini telah dilengkapi `docker-compose.yml` multi-container untuk kemudahan deploy di VPS Ubuntu/Debian:

```bash
# 1. Clone repositori ke VPS
git clone <repo-url> /var/www/alhamid
cd /var/www/alhamid

# 2. Sesuaikan konfigurasi .env (password DB, domain, dsb)
cp backend/.env.example backend/.env

# 3. Jalankan seluruh layanan dengan Docker Compose
docker compose up -d --build
```

### Konfigurasi Reverse Proxy Nginx + SSL di VPS
```nginx
server {
    server_name alhamid.sch.id www.alhamid.sch.id;

    # Frontend Svelte
    location / {
        proxy_pass http://127.0.0.1:5173;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
    }

    # Backend API Go
    location /api/ {
        proxy_pass http://127.0.0.1:8088/api/;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
    }

    # Static Uploads
    location /uploads/ {
        proxy_pass http://127.0.0.1:8088/uploads/;
    }
}
```
Aktifkan SSL gratis dengan Certbot:
```bash
certbot --nginx -d alhamid.sch.id -d www.alhamid.sch.id
```

---

## 💡 Pembahasan Strategis & Rekomendasi Pengembangan Lanjutan

1. **Sistem PPDB Terintegrasi (Penerimaan Santri Baru)**:
   - Formulir multi-step pendaftaran santri baru dengan upload berkas (ijazah, KK, akta, pas foto).
   - Verifikasi berkas oleh admin, pencetakan nomor peserta ujian seleksi secara otomatis (PDF).
2. **Kanal Donasi & Wakaf Produktif**:
   - Modul transparansi pembangunan fasilitas (misal: perluasan asrama, maktabah digital) dengan opsi pembayaran QRIS / transfer bank instan.
3. **Portal Wali Santri / SIPren (Sistem Informasi Pesantren)**:
   - Akses nilai raport madrasah diniyah & formal.
   - Catatan setoran hafalan Qur'an santri (PMQ) secara berkala.
4. **Kalender Akademik & Jadwal Pengajian Kitab Hijriyah**:
   - Widget hitung mundur waktu sholat dan kalender perayaan hari besar Islam / haul masyayikh.

