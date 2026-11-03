package database

import (
	"log"
	"time"

	"alhamid-backend/models"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func SeedInitialData(db *gorm.DB) {
	// 1. User Admin
	var admin models.User
	if err := db.Where("username = ?", "admin").First(&admin).Error; err != nil {
		hashed, _ := bcrypt.GenerateFromPassword([]byte("alhamidku123"), bcrypt.DefaultCost)
		admin = models.User{
			Username:     "admin",
			Email:        "admin@alhamid.sch.id",
			PasswordHash: string(hashed),
			FullName:     "Administrator Al-Hamid",
			Role:         "admin",
		}
		db.Create(&admin)
		log.Println("[SEED] User admin berhasil dibuat: admin / alhamidku123")
	} else {
		hashed, _ := bcrypt.GenerateFromPassword([]byte("alhamidku123"), bcrypt.DefaultCost)
		db.Model(&admin).Update("password_hash", string(hashed))
		log.Println("[SEED] User admin password diperbarui: alhamidku123")
	}

	// 2. Settings
	seedSettings(db)

	// 3. Categories
	seedCategories(db)

	// 4. Banners
	seedBanners(db)

	// 5. Education Units
	seedEducationUnits(db)

	// 6. Caregivers
	seedCaregivers(db)

	// 7. Facilities
	seedFacilities(db)

	// 8. Achievements
	seedAchievements(db)

	// 9. Testimonials
	seedTestimonials(db)

	// 10. Alumni
	seedAlumni(db)

	// 11. Articles
	seedArticles(db)

	// 12. Media (Photos & Videos)
	seedMedia(db)
}

func seedSettings(db *gorm.DB) {
	settings := []models.Setting{
		{Key: "site_name", Value: "Pondok Pesantren Al-Hamid", Category: "general"},
		{Key: "site_font_preset", Value: "architectural", Category: "general"},
		{Key: "site_tagline", Value: "Meneguhkan Sanad Keilmuan, Mengintegrasikan Karakter, Membangun Generasi Emas", Category: "general"},
		{Key: "site_description", Value: "Lembaga pendidikan Islam terpadu yang memadukan kedalaman kajian kitab kuning (turats) dengan keunggulan akademik sains dan teknologi modern.", Category: "general"},
		{Key: "site_email", Value: "sekretariat@alhamid.sch.id", Category: "contact"},
		{Key: "site_phone", Value: "+62 812-3456-7890", Category: "contact"},
		{Key: "site_whatsapp", Value: "6281234567890", Category: "contact"},
		{Key: "site_address", Value: "Jl. Pesantren Al-Hamid No. 01, Cintamulya, Candipuro, Lampung Selatan", Category: "contact"},
		{Key: "site_maps", Value: "https://maps.google.com/?q=Pondok+Pesantren+Al-Hamid", Category: "contact"},
		{Key: "social_instagram", Value: "https://instagram.com/pesantren_alhamid", Category: "social"},
		{Key: "social_youtube", Value: "https://youtube.com/@pesantrenalhamidofficial", Category: "social"},
		{Key: "social_facebook", Value: "https://facebook.com/pesantrenalhamid", Category: "social"},
		{Key: "social_tiktok", Value: "https://tiktok.com/@pesantrenalhamid", Category: "social"},
		{Key: "about_profile", Value: `Pondok Pesantren Al-Hamid didirikan dengan komitmen kuat untuk menjadi benteng moral, intelektual, dan spiritual umat. Berakar dari tradisi luhur pesantren salafiyah yang menjaga kemurnian sanad keilmuan, Al-Hamid senantiasa dinamis merespons tantangan zaman melalui integrasi pendidikan formal modern (MTs, MA, MAI) dan pendidikan non-formal (Madrasah Diniyah dan Pendidikan Murottilil Qur'an).

Dengan atmosfer pembelajaran yang asri, disiplin tinggi, serta asuhan langsung dari para masyayikh dan asatidz berpengalaman, santri ditempa untuk mandiri, berwawasan luas, moderat (tawasuth), dan siap menjadi pelopor peradaban di berbagai kancah nasional maupun internasional.`, Category: "profile"},
		{Key: "about_vision", Value: "Terwujudnya lembaga pendidikan Islam unggul yang melahirkan generasi berakhlakul karimah, mutafaqqih fiddin, berwawasan global, dan berdaya saing tinggi.", Category: "profile"},
		{Key: "about_mission", Value: `1. Menyelenggarakan kajian keislaman berbasis kutubut turats dengan sanad keilmuan muttashil.
2. Membentuk pribadi santri yang disiplin, berintegritas, mandiri, dan berakhlak mulia.
3. Menyediakan pendidikan formal berkualitas tinggi dalam bidang sains, humaniora, dan teknologi.
4. Mengembangkan kecakapan berbahasa asing (Arab dan Inggris) serta literasi digital santri.
5. Membangun sinergi berkelanjutan dengan masyarakat, alumni, dan institusi global.`, Category: "profile"},
	}

	for _, s := range settings {
		var existing models.Setting
		if err := db.Where("key = ?", s.Key).First(&existing).Error; err != nil {
			db.Create(&s)
		}
	}
}

func seedCategories(db *gorm.DB) {
	categories := []models.Category{
		{Name: "Kajian & Kitab", Slug: "kajian-kitab", Type: "berita", Description: "Ulasan mutiara kitab kuning, tadabbur ayat, dan hadits nabawi."},
		{Name: "Kabar Pesantren", Slug: "kabar-pesantren", Type: "berita", Description: "Berita kegiatan harian, dinamika santri, dan agenda pesantren."},
		{Name: "Prestasi Santri", Slug: "prestasi-santri", Type: "berita", Description: "Catatan capaian gemilang para santri Al-Hamid di berbagai ajang."},
		{Name: "Pengumuman", Slug: "pengumuman", Type: "berita", Description: "Informasi resmi penerimaan santri baru, kalender akademik, dan perpulangan."},
		{Name: "Opini & Literasi", Slug: "opini-literasi", Type: "berita", Description: "Gagasan dan karya tulis asatidz serta santri Al-Hamid."},

		{Name: "Kegiatan Pesantren", Slug: "foto-kegiatan", Type: "galeri_foto", Description: "Dokumentasi ragam aktivitas keseharian santri."},
		{Name: "Haflah & Wisuda", Slug: "foto-haflah", Type: "galeri_foto", Description: "Momen perayaan kelulusan dan imtihan santri."},
		{Name: "Fasilitas & Kompleks", Slug: "foto-fasilitas", Type: "galeri_foto", Description: "Lansekap gedung, asrama, dan fasilitas penunjang."},

		{Name: "Kajian Masyayikh", Slug: "video-kajian", Type: "galeri_video", Description: "Video rekaman pengajian kitab kuning dan tausiyah dewan pengasuh."},
		{Name: "Profil & Dokumenter", Slug: "video-profil", Type: "galeri_video", Description: "Video dokumenter perjalanan dan profil lembaga Al-Hamid."},
		{Name: "Seni & Kreasi Santri", Slug: "video-kreasi", Type: "galeri_video", Description: "Penampilan hadroh, pidato multi bahasa, dan expo santri."},
	}

	for _, c := range categories {
		var existing models.Category
		if err := db.Where("slug = ?", c.Slug).First(&existing).Error; err != nil {
			db.Create(&c)
		}
	}
}

func seedBanners(db *gorm.DB) {
	var count int64
	db.Model(&models.Banner{}).Count(&count)
	if count > 0 {
		return
	}

	banners := []models.Banner{
		{
			Title:      "Menempa Akhlak, Mengukir Prestasi Masa Depan",
			Subtitle:   "Selamat Datang di Portal Resmi Pondok Pesantren Al-Hamid. Paduan harmoni antara kedalaman turats pesantren salaf dan keunggulan sains modern.",
			ImageUrl:   "https://images.unsplash.com/photo-1542838132-92c53300491e?auto=format&fit=crop&w=1920&q=80",
			ButtonText: "Jelajahi Pesantren",
			ButtonLink: "#tentang",
			Badge:      "Tahun Ajaran 2026/2027",
			Order:      1,
			IsActive:   true,
		},
		{
			Title:      "Penerimaan Santri Baru (PSB) Telah Dibuka",
			Subtitle:   "Bergabunglah bersama ribuan kader umat di MTs & MA Cintamulya, MAI Cintamulya, Madrasah Diniyah Al-Hamid, serta PMQ.",
			ImageUrl:   "https://images.unsplash.com/photo-1523240795612-9a054b0db644?auto=format&fit=crop&w=1920&q=80",
			ButtonText: "Daftar Sekarang",
			ButtonLink: "/kontak",
			Badge:      "Pendaftaran Online",
			Order:      2,
			IsActive:   true,
		},
		{
			Title:      "Kajian Kitab Kuning Bersanad & Mutu Akademik Prima",
			Subtitle:   "Dibimbing langsung oleh dewan masyayikh berkompeten, menghasilkan santri yang teguh berprinsip dan luwes berkiprah di masyarakat.",
			ImageUrl:   "https://images.unsplash.com/photo-1577495508048-b635879837f1?auto=format&fit=crop&w=1920&q=80",
			ButtonText: "Lihat Program Pendidikan",
			ButtonLink: "/pendidikan",
			Badge:      "Tradisi Luhur Pesantren",
			Order:      3,
			IsActive:   true,
		},
	}

	for _, b := range banners {
		db.Create(&b)
	}
}

func seedEducationUnits(db *gorm.DB) {
	var count int64
	db.Model(&models.EducationUnit{}).Count(&count)
	if count > 0 {
		return
	}

	units := []models.EducationUnit{
		// Non-Formal 1: Madrasah Diniyah Al-Hamid
		{
			Name:      "Madrasah Diniyah Al-Hamid",
			Slug:      "madrasah-diniyah-alhamid",
			Type:      "non_formal",
			Tagline:   "Pilar Utama Penguasaan Khazanah Kitab Kuning dan Sanad Keilmuan",
			HeroImage: "https://images.unsplash.com/photo-1532012164546-f432f2e3777a?auto=format&fit=crop&w=1200&q=80",
			Badge:     "Kurikulum Pesantren Salaf",
			ShortDesc: "Pendidikan keislaman berjenjang (Ula, Wustha, Ulya) yang memfokuskan santri pada penguasaan gramatika Arab (Nahwu, Shorof) serta kajian Fiqih, Ushul Fiqih, Tauhid, dan Hadits.",
			FullDesc: `Madrasah Diniyah Al-Hamid merupakan jantung spiritual dan intelektual pesantren. Menggunakan metodologi klasik yang teruji—sorogan, bandongan, dan musyawarah bahtsul masail—para santri dibimbing mengkaji teks-teks otoritatif turats Islam.

Jenjang pendidikan ditempuh selama 6 tahun yang terbagi dalam tingkat Ula (dasar kaidah bahasa & fiqih ibadah), Wustha (pendalaman tata bahasa, fiqih muamalah & tarikh), serta Ulya (kajian spesialisasi fiqih munakahat/jinayat, ushul fiqih, tafsir Jalalain, dan balaghah).`,
			Curriculum: `- **Jenjang Ula**: Jurumiyah, Al-Amtsilah At-Tashrifiyyah, Safinatun Najah, Aqidatul Awam, Al-Akhlaq lil Banin/Banat.
- **Jenjang Wustha**: Imrithi, Matan Al-Bina, Taqrib/Fathul Qorib, Taisir Musthalah Hadits, Ta'limul Muta'allim.
- **Jenjang Ulya**: Alfiyah Ibnu Malik, Fathul Mu'in, Al-Waraqat, Tafsir Jalalain, Bulughul Maram, Jawahirul Bukhari.`,
			Requirements: `- Terdaftar aktif sebagai santri mukim Pondok Pesantren Al-Hamid.
- Mengikuti tes pemetaan kemampuan membaca Al-Qur'an dan dasar kaidah bahasa.
- Memiliki komitmen tinggi mengikuti seluruh jadwal musyawarah dan tadarrus kitab.`,
			FacilitiesList: `Ruang halaqoh berkarpet nyaman, Maktabah Turats dengan ribuan koleksi manuskrip & syarah, perpustakaan audio kajian, serta dewan asatidz pembimbing 24 jam.`,
			HeadmasterName: "K.H. Syamsul Huda, S.Pd.I.",
			Order:          1,
			IsActive:       true,
		},
		// Non-Formal 2: PMQ (Pendidikan Murottilil Qur'an)
		{
			Name:      "PMQ (Pendidikan Murottilil Qur'an)",
			Slug:      "pmq-alhamid",
			Type:      "non_formal",
			Tagline:   "Mencetak Generasi Qur'ani Bertartil, Fasih, dan Berakhlak Al-Qur'an",
			HeroImage: "https://images.unsplash.com/photo-1609599006353-e629aaabfeae?auto=format&fit=crop&w=1200&q=80",
			Badge:     "Talaqqi & Nagham Bersanad",
			ShortDesc: "Program pembinaan bacaan Al-Qur'an secara intensif dengan penekanan pada makharijul huruf, sifatul huruf, hukum tajwid praktis, serta variasi lagu/nagham yang indah.",
			FullDesc: `Pendidikan Murottilil Qur'an (PMQ) adalah wadah spesialisasi bagi santri untuk memperindah dan menyempurnakan bacaan kalam Ilahi. Metode yang diterapkan mengombinasikan talaqqi musyafahah (face-to-face dengan guru bersanad) serta pembinaan tilawah berirama (Bayyati, Shoba, Hijaz, Nahawand, Rast, Sika, dan Jiharkah).

Program ini mencakup jenjang dasar (fashohah), jenjang lanjutan (nagham & tartil panggung), serta bimbingan khusus Tahfidzul Qur'an bagi santri yang berkeinginan menghafal 30 juz secara mutqin.`,
			Curriculum: `- **Tahsin & Tajwid**: Matan Al-Jazariyah, Tuhfatul Athfal, Makharij & Sifat Huruf.
- **Nagham/Irama**: 7 Maqamat Al-Qur'an standar nasional & internasional.
- **Talaqqi Sanad**: Setoran hafalan dan bacaan langsung ke ustadz pemegang sanad riwayat Imam Ashim 'an Hafsh.
- **Majelis Semaan**: Semaan rutin selapanan dan tasmi' Al-Qur'an bil ghoib.`,
			Requirements: `- Lulus tes seleksi tajwid dasar dan kelancaran membaca mushaf pojok.
- Mengikuti displin jam latihan fashohah ba'da Shubuh dan ba'da Maghrib.`,
			FacilitiesList: `Studio rekam tilawah santri, majelis talaqqi privat, mikrofon standar festival tilawah, serta mushaf tahfidz per juz.`,
			HeadmasterName: "Ust. H. Ahmad Fakhruddin, Al-Hafidz",
			Order:          2,
			IsActive:       true,
		},
		// Formal 1: MTs Cintamulya
		{
			Name:      "MTs Cintamulya",
			Slug:      "mts-cintamulya",
			Type:      "formal",
			Tagline:   "Pendidikan Menengah Pertama Berbasis Karakter dan Keunggulan Sains",
			HeroImage: "https://images.unsplash.com/photo-1580582932707-520aed937b7b?auto=format&fit=crop&w=1200&q=80",
			Badge:     "Akreditasi A Unggul",
			ShortDesc: "Madrasah Tsanawiyah formal di bawah binaan Kemenag yang mengintegrasikan kurikulum nasional dengan kurikulum kepesantrenan terpadu.",
			FullDesc: `Madrasah Tsanawiyah (MTs) Cintamulya hadir mendidik generasi usia remaja dengan pendekatan holistik. Di sini, penguasaan sains, matematika, bahasa internasional, dan teknologi dipadukan secara seimbang dengan pembiasaan adab, sholat berjamaah, dan tahfidz juz 30.

Lulusan MTs Cintamulya memiliki bekal ganda: ijazah resmi negara untuk melanjutkan ke jenjang SMA/MA manapun serta fondasi agama yang kokoh untuk kehidupan sehari-hari.`,
			Curriculum: `- **Kurikulum Merdeka Kemenag**: PAI (Aqidah Akhlak, Fikih, SKI, Qur'an Hadits), Bahasa Arab, Bahasa Inggris, Matematika, IPA, IPS, Informatika.
- **Muatan Lokal Pesantren**: Pegon, Khot Arab, Imla, Muhadhoroh 3 Bahasa, Praktik Ibadah Amaliyah.
- **Ekstrakurikuler**: Pramuka, Robotika Sederhana, Hadroh, Karate, Jurnalistik Santri, PMR.`,
			Requirements: `- Mengisi formulir pendaftaran resmi (online/offline).
- Fotokopi Ijazah/SKL SD/MI yang telah dilegalisir.
- Fotokopi Akta Kelahiran dan Kartu Keluarga (KK).
- Pas foto resmi ukuran 3x4 (4 lembar).
- Surat keterangan sehat dari dokter.`,
			FacilitiesList: `Gedung ruang kelas ber-AC dan proyektor, Laboratorium Komputer, Lab IPA terpadu, Lapangan olahraga terpadu, Perpustakaan modern.`,
			HeadmasterName: "Drs. H. M. Shodiqin, M.Pd.",
			Order:          3,
			IsActive:       true,
		},
		// Formal 2: MA Cintamulya
		{
			Name:      "MA Cintamulya",
			Slug:      "ma-cintamulya",
			Type:      "formal",
			Tagline:   "Mempersiapkan Kader Ulama Cendekiawan Menembus Kampus Idaman",
			HeroImage: "https://images.unsplash.com/photo-1523050854058-8df90110c9f1?auto=format&fit=crop&w=1200&q=80",
			Badge:     "Terakreditasi A Unggul",
			ShortDesc: "Madrasah Aliyah unggulan dengan pilihan jurusan MIPA, IPS, dan Keagamaan, didukung program intensif beasiswa PTN dan studi luar negeri.",
			FullDesc: `Madrasah Aliyah (MA) Cintamulya merupakan jenjang pendidikan menengah atas yang dirancang untuk melahirkan generasi pemikir berdaya saing global. Kurikulum diperkaya dengan program bimbingan UTBK-SNBT, persiapan kuliah ke Timur Tengah (Al-Azhar Kairo, Maroko, Tunisia), serta pengembangan riset ilmiah remaja (KIR).

Santri MA Cintamulya dilatih untuk berpikir kritis, menguasai literasi data, fasih berdebat dalam bahasa Arab dan Inggris, serta menjunjung tinggi akhlakul karimah.`,
			Curriculum: `- **Peminatan MIPA**: Fisika, Kimia, Biologi, Matematika Lanjutan, Riset Karya Tulis Ilmiah.
- **Peminatan IPS**: Ekonomi, Geografi, Sosiologi, Sejarah, Entrepreneurship.
- **Peminatan Keagamaan**: Tafsir-Ilmu Tafsir, Hadits-Ilmu Hadits, Fiqih-Ushul Fiqih, Bahasa Arab Intensif, Ilmu Kalam.
- **Program Akselerasi**: Klinik UTBK, TOEFL & TOAFL Preparation, Bimbingan Mu'adalah Kairo.`,
			Requirements: `- Lulus MTs / SMP sederajat.
- Menyerahkan rapor semester 1-5 dan surat keterangan kelakuan baik.
- Lolos seleksi tes potensi akademik dan wawancara komitmen orang tua.`,
			FacilitiesList: `Laboratorium Kimia-Fisika-Biologi, Laboratorium Multimedia & Bahasa, Smart Classroom, Asrama binaan khusus MA.`,
			HeadmasterName: "H. Abdullah Munir, S.Ag., M.Ag.",
			Order:          4,
			IsActive:       true,
		},
		// Formal 3: MAI Cintamulya
		{
			Name:      "MAI Cintamulya",
			Slug:      "mai-cintamulya",
			Type:      "formal",
			Tagline:   "Inovasi Vokasi Islam: Menyatukan Dakwah Digital, Teknologi Informasi, dan Kemandirian",
			HeroImage: "https://images.unsplash.com/photo-1531482615713-2afd69097998?auto=format&fit=crop&w=1200&q=80",
			Badge:     "Program Kejuruan & Vokasi",
			ShortDesc: "Madrasah Aliyah Kejuruan/Informatika (MAI) yang memfokuskan santri pada keahlian Rekayasa Perangkat Lunak, Multimedia Dakwah, dan Technopreneurship.",
			FullDesc: `MAI Cintamulya lahir sebagai jawaban konkret pesantren terhadap era kecerdasan buatan dan industri 4.0. Santri tidak hanya menguasai ilmu agama, tetapi juga dibekali sertifikasi kompetensi di bidang software development, desain UI/UX, produksi audio-visual, dan digital marketing.

Lulusan MAI siap berkarier langsung di industri digital, melanjutkan studi ke perguruan tinggi teknik, maupun membuka usaha kreatif secara mandiri.`,
			Curriculum: `- **Kompetensi Keahlian IT**: Web Development (HTML/CSS/JS, Go, Svelte, Database), Mobile App Basics, UI/UX Design.
- **Multimedia & Dakwah**: Video Editing, Animasi, Live Streaming Broadcast, Podcasting Islam.
- **Entrepreneurship**: Business Model Canvas, Digital Marketing, Manajemen Usaha Pesantren.`,
			Requirements: `- Lulusan SMP/MTs sederajat yang memiliki minat besar pada teknologi dan multimedia.
- Membawa laptop pribadi untuk keperluan praktikum coding & desain.`,
			FacilitiesList: `Lab Komputer High-End dengan koneksi gigabit fiber optic, Studio Podcast & Broadcasting, Ruang Co-Working Santri.`,
			HeadmasterName: "Gus H. Muhammad Zaim, Lc., M.A.",
			Order:          5,
			IsActive:       true,
		},
	}

	for _, u := range units {
		db.Create(&u)
	}
}

func seedCaregivers(db *gorm.DB) {
	var count int64
	db.Model(&models.Caregiver{}).Count(&count)
	if count > 0 {
		return
	}

	caregivers := []models.Caregiver{
		{
			Name:          "K.H. Ahmad Hamidurrohman",
			RoleTitle:     "Pengasuh Utama Pondok Pesantren",
			PhotoUrl:      "https://images.unsplash.com/photo-1507003211169-0a1dd7228f2d?auto=format&fit=crop&w=600&q=80",
			Bio:           "Ulama kharismatik yang mendedikasikan hidupnya untuk dakwah, tarbiyah, dan pemberdayaan masyarakat pesantren selama lebih dari 35 tahun.",
			SanadKeilmuan: "Muttashil hingga Syaikhona Kholil Bangkalan dan Hadratus Syaikh Hasyim Asy'ari melalui jalur Masyayikh Lirboyo dan Sarang.",
			Order:         1,
			IsActive:      true,
		},
		{
			Name:          "Nyai Hj. Siti Fatimah",
			RoleTitle:     "Ketua Yayasan & Pengasuh Keputrian",
			PhotoUrl:      "https://images.unsplash.com/photo-1544005313-94ddf0286df2?auto=format&fit=crop&w=600&q=80",
			Bio:           "Penggerak pendidikan karakter santri putri, membina ratusan santriwati dengan pendekatan keibuan yang hangat dan penuh keteladanan.",
			SanadKeilmuan: "Jalur sanad Al-Qur'an bil Ghoib dari Masyayikh Kudus dan Krapyak Yogyakarta.",
			Order:         2,
			IsActive:      true,
		},
		{
			Name:          "Gus H. Muhammad Zaim, Lc., M.A.",
			RoleTitle:     "Direktur Lembaga Pendidikan Formal & Digital",
			PhotoUrl:      "https://images.unsplash.com/photo-1500648767791-00dcc994a43e?auto=format&fit=crop&w=600&q=80",
			Bio:           "Alumnus Universitas Al-Azhar Kairo dan Magister Pendidikan Islam, memelopori integrasi teknologi modern dan digitalisasi di Al-Hamid.",
			SanadKeilmuan: "Sanad Hadits dan Ushul Fiqih bersambung ke para Masyayikh Kibar Al-Azhar Kairo, Mesir.",
			Order:         3,
			IsActive:      true,
		},
		{
			Name:          "K.H. Syamsul Huda, S.Pd.I.",
			RoleTitle:     "Katib Majelis Masyayikh & Kepala Diniyah",
			PhotoUrl:      "https://images.unsplash.com/photo-1472099645785-5658abf4ff4e?auto=format&fit=crop&w=600&q=80",
			Bio:           "Pakar gramatika Arab dan fiqih muamalah yang mengampu kajian musyawarah bahtsul masail santri tingkat tinggi.",
			SanadKeilmuan: "Sanad Alfiyah dan Fathul Mu'in melalui Masyayikh Ploso Kediri.",
			Order:         4,
			IsActive:      true,
		},
	}

	for _, c := range caregivers {
		db.Create(&c)
	}
}

func seedFacilities(db *gorm.DB) {
	var count int64
	db.Model(&models.Facility{}).Count(&count)
	if count > 0 {
		return
	}

	facilities := []models.Facility{
		{
			Name:        "Masjid Jami' Al-Hamid",
			Category:    "Ibadah",
			Description: "Pusat kegiatan spiritual santri, sholat berjamaah lima waktu, kajian kitab kuning ba'da Maghrib & Shubuh, serta halaqoh Qur'an.",
			ImageUrl:    "https://images.unsplash.com/photo-1564769625905-50e93615e769?auto=format&fit=crop&w=800&q=80",
			Order:       1,
		},
		{
			Name:        "Kompleks Asrama Putra & Putri Terpadu",
			Category:    "Asrama",
			Description: "Hunian santri yang bersih, tertata rapi, berfentilasi alami, dilengkapi lemari pribadi, serta diawasi oleh pembina asrama 24 jam.",
			ImageUrl:    "https://images.unsplash.com/photo-1555854877-bab0e564b8d5?auto=format&fit=crop&w=800&q=80",
			Order:       2,
		},
		{
			Name:        "Maktabah Turats & Perpustakaan Digital",
			Category:    "Akademik",
			Description: "Menyimpan lebih dari 5.000 judul kitab kuning, referensi sains, jurnal ilmiah, dan puluhan terminal e-library untuk riset santri.",
			ImageUrl:    "https://images.unsplash.com/photo-1521587760476-6c12a4b040da?auto=format&fit=crop&w=800&q=80",
			Order:       3,
		},
		{
			Name:        "Laboratorium Sains & Komputer MA Cintamulya",
			Category:    "Akademik",
			Description: "Dilengkapi peralatan praktikum fisika, kimia, biologi, serta puluhan unit PC high-spec untuk pembelajaran programming dan multimedia.",
			ImageUrl:    "https://images.unsplash.com/photo-1581092918056-0c4c3acd3789?auto=format&fit=crop&w=800&q=80",
			Order:       4,
		},
		{
			Name:        "Gelanggang Olahraga & Lapangan Futsal",
			Category:    "Olahraga",
			Description: "Fasilitas olahraga santri meliputi lapangan futsal rumput sintetis, lapangan voli, bulu tangkis, dan tenis meja.",
			ImageUrl:    "https://images.unsplash.com/photo-1518604666864-742395639b8d?auto=format&fit=crop&w=800&q=80",
			Order:       5,
		},
		{
			Name:        "Pos Kesehatan Pesantren (Poskestren)",
			Category:    "Penunjang",
			Description: "Layanan medis dasar santri dengan dokter dan perawat jaga, bekerja sama dengan Puskesmas dan Rumah Sakit Daerah rujukan.",
			ImageUrl:    "https://images.unsplash.com/photo-1519494026892-80bbd2d6fd0d?auto=format&fit=crop&w=800&q=80",
			Order:       6,
		},
	}

	for _, f := range facilities {
		db.Create(&f)
	}
}

func seedAchievements(db *gorm.DB) {
	var count int64
	db.Model(&models.Achievement{}).Count(&count)
	if count > 0 {
		return
	}

	achievements := []models.Achievement{
		{
			Title:         "Juara 1 Musabaqah Qira'atil Kutub (MQK) Nasional Cabang Fiqih Ulya",
			Year:          2026,
			Level:         "Nasional",
			Category:      "MQK",
			RecipientName: "Muhammad Faruq Al-Hafidz",
			Description:   "Membaca dan mempertahankan syarah Fathul Mu'in di hadapan dewan juri Kemenag RI.",
			ImageUrl:      "https://images.unsplash.com/photo-1578269174936-2709b6aeb913?auto=format&fit=crop&w=800&q=80",
		},
		{
			Title:         "Medali Emas Olimpiade Sains Madrasah (KSM) Nasional Bidang Matematika Terintegrasi",
			Year:          2025,
			Level:         "Nasional",
			Category:      "Sains",
			RecipientName: "Ananda Zahra Putri (Santri MA Cintamulya)",
			Description:   "Menyingkirkan ratusan delegasi madrasah se-Indonesia dengan skor tertinggi.",
			ImageUrl:      "https://images.unsplash.com/photo-1532094349884-543bc11b234d?auto=format&fit=crop&w=800&q=80",
		},
		{
			Title:         "Juara 1 Debat Bahasa Arab Tingkat Provinsi Lampung",
			Year:          2025,
			Level:         "Provinsi",
			Category:      "Bahasa",
			RecipientName: "Tim Debat Al-Hamid Arabic Club",
			Description:   "Memenangkan mosi isu-isu kontemporer dan maqashid syariah.",
			ImageUrl:      "https://images.unsplash.com/photo-1524178232363-1fb2b075b655?auto=format&fit=crop&w=800&q=80",
		},
		{
			Title:         "Juara Umum Festival Hadroh & Banjari se-Sumatera Bagian Selatan",
			Year:          2024,
			Level:         "Regional",
			Category:      "Seni",
			RecipientName: "Grup Hadroh Syauqul Habib Al-Hamid",
			Description:   "Meraih juara vokal terbaik dan aransemen hadroh klasik terbaik.",
			ImageUrl:      "https://images.unsplash.com/photo-1511671782779-c97d3d27a1d4?auto=format&fit=crop&w=800&q=80",
		},
	}

	for _, a := range achievements {
		db.Create(&a)
	}
}

func seedTestimonials(db *gorm.DB) {
	var count int64
	db.Model(&models.Testimonial{}).Count(&count)
	if count > 0 {
		return
	}

	testimonials := []models.Testimonial{
		{
			Name:      "Dr. H. Fathurrahman, M.A.",
			Role:      "Alumni Angkatan 2010 / Dosen UIN & Peneliti",
			Content:   "Mondok di Al-Hamid memberi saya bekal fondasi nalar kritis kitab kuning yang luar biasa berharga ketika menempuh studi doktoral di Timur Tengah. Al-Hamid mengajarkan adab sebelum ilmu.",
			AvatarUrl: "https://images.unsplash.com/photo-1534528741775-53994a69daeb?auto=format&fit=crop&w=400&q=80",
			Rating:    5,
			Order:     1,
			IsActive:  true,
		},
		{
			Name:      "Ir. Bambang Trihatmojo",
			Role:      "Wali Santri MA Cintamulya (Jakarta)",
			Content:   "Awalnya ragu menitipkan anak di pesantren, tapi melihat perkembangan putra kami yang mandiri, sholat tepat waktu, dan kini lolos seleksi beasiswa teknik, kami sangat bersyukur telah memilih Al-Hamid.",
			AvatarUrl: "https://images.unsplash.com/photo-1506794778202-cad84cf45f1d?auto=format&fit=crop&w=400&q=80",
			Rating:    5,
			Order:     2,
			IsActive:  true,
		},
		{
			Name:      "Aisyah Ramadhani, S.Kom.",
			Role:      "Alumni MAI Cintamulya / Software Engineer",
			Content:   "Jurusan IT di MAI Cintamulya benar-benar relevan dengan industri kerja! Di pesantren saya belajar ngaji sekaligus ngoding web dan mobile. Suasana belajarnya sangat kondusif.",
			AvatarUrl: "https://images.unsplash.com/photo-1517841905240-472988babdf9?auto=format&fit=crop&w=400&q=80",
			Rating:    5,
			Order:     3,
			IsActive:  true,
		},
	}

	for _, t := range testimonials {
		db.Create(&t)
	}
}

func seedAlumni(db *gorm.DB) {
	var count int64
	db.Model(&models.Alumni{}).Count(&count)
	if count > 0 {
		return
	}

	alumni := []models.Alumni{
		{
			Name:            "K.H. Ihsan Mahfudz, Lc.",
			GraduationYear:  2012,
			EducationUnit:   "Madrasah Diniyah Al-Hamid & MA Cintamulya",
			Profession:      "Pengasuh Ponpes Bina Insan & Dai Nasional",
			CurrentLocation: "Bandung, Jawa Barat",
			Story:           "Menyebarkan dakwah wasathiyah dan mendirikan cabang lembaga pendidikan Al-Qur'an.",
			PhotoUrl:        "https://images.unsplash.com/photo-1500648767791-00dcc994a43e?auto=format&fit=crop&w=500&q=80",
			SocialLink:      "https://instagram.com",
			Order:           1,
		},
		{
			Name:            "dr. Nabila Zulfa, Sp.A",
			GraduationYear:  2015,
			EducationUnit:   "MA Cintamulya (MIPA)",
			Profession:      "Dokter Spesialis Anak RS Hermina",
			CurrentLocation: "Bandar Lampung",
			Story:           "Mewujudkan cita-cita menjadi dokter yang berakhlakul karimah dan gemar mengabdi untuk santri dan dhuafa.",
			PhotoUrl:        "https://images.unsplash.com/photo-1559839734-2b71ea197ec2?auto=format&fit=crop&w=500&q=80",
			SocialLink:      "https://linkedin.com",
			Order:           2,
		},
		{
			Name:            "Rizky Fauzi, S.T., M.Sc.",
			GraduationYear:  2017,
			EducationUnit:   "MAI Cintamulya",
			Profession:      "Lead AI Engineer di Perusahaan Multinasional",
			CurrentLocation: "Singapura",
			Story:           "Berawal dari lab komputer pesantren Cintamulya hingga kini memimpin pengembangan sistem kecerdasan buatan kelas dunia.",
			PhotoUrl:        "https://images.unsplash.com/photo-1492562080023-ab3db95bfbce?auto=format&fit=crop&w=500&q=80",
			SocialLink:      "https://linkedin.com",
			Order:           3,
		},
	}

	for _, al := range alumni {
		db.Create(&al)
	}
}

func seedArticles(db *gorm.DB) {
	var count int64
	db.Model(&models.Article{}).Count(&count)
	if count > 0 {
		return
	}

	var catKajian, catKabar, catPengumuman models.Category
	db.Where("slug = ?", "kajian-kitab").First(&catKajian)
	db.Where("slug = ?", "kabar-pesantren").First(&catKabar)
	db.Where("slug = ?", "pengumuman").First(&catPengumuman)

	articles := []models.Article{
		{
			Title:        "Meneladani Akhlaqul Karimah dalam Menuntut Ilmu Turats",
			Slug:         "meneladani-akhlaqul-karimah-dalam-menuntut-ilmu",
			CategoryID:   catKajian.ID,
			Excerpt:      "Syekh Az-Zarnuji dalam Ta'limul Muta'allim menegaskan bahwa ilmu tanpa adab laksana pohon yang tak berbuah. Bagaimana santri modern menginternalisasikannya?",
			Content: `<h3>Urgensi Adab Sebelum Ilmu</h3>
<p>Dalam tradisi pesantren Ahlussunnah wal Jama'ah, adab menempati derajat teratas sebelum seseorang menyelami samudra ilmu pengetahuan. Imam Malik pernah menasihati seorang pemuda Quraisy: <em>"Pelajarilah adab sebelum engkau mempelajari suatu ilmu."</em></p>
<p>Di era modern yang serba cepat, godaan untuk bersikap instan kerap mengikis rasa tawadhu' dan takzim kepada guru. Pondok Pesantren Al-Hamid senantiasa membentengi santri dengan pembiasaan adab lahir dan batin, mulai dari cara memegang kitab, menyimak dawuh masyayikh, hingga adab bergaul antarsesama santri.</p>
<h3>Kunci Keberkahan Ilmu</h3>
<p>Keberkahan ilmu tidak diukur dari seberapa cepat seseorang menghafal atau seberapa banyak gelar akademik yang diraih, melainkan sejauh mana ilmu tersebut membawa ketenangan jiwa, ketakutan kepada Allah, serta manfaat nyata bagi kemaslahatan masyarakat.</p>`,
			ThumbnailUrl: "https://images.unsplash.com/photo-1542838132-92c53300491e?auto=format&fit=crop&w=1000&q=80",
			Author:       "K.H. Ahmad Hamidurrohman",
			Status:       "published",
			IsFeatured:   true,
			ViewCount:    342,
			PublishedAt:  time.Now().Add(-48 * time.Hour),
		},
		{
			Title:        "Haflah Akhirussanah & Wisuda Akbar Santri Al-Hamid 2026 Berlangsung Khidmat",
			Slug:         "haflah-akhirussanah-dan-wisuda-akbar-2026",
			CategoryID:   catKabar.ID,
			Excerpt:      "Ratusan santri jenjang Madrasah Diniyah, MTs MA Cintamulya, dan MAI Cintamulya berhasil menyelesaikan masa studi dengan capaian membanggakan.",
			Content: `<p>Suasana haru dan penuh syukur menyelimuti Auditorium Utama Pondok Pesantren Al-Hamid dalam perhelatan akbar Haflah Akhirussanah dan Wisuda Santri tahun 2026. Acara dihadiri oleh segenap dewan masyayikh, para asatidz, pejabat pemerintahan, serta ribuan wali santri dari berbagai penjuru tanah air.</p>
<p>Pada wisuda kali ini, tercatat puluhan santri menuntaskan hafalan 30 juz Al-Qur'an bil ghoib melalui program PMQ, serta santri kelas akhir MA Cintamulya yang telah mengantongi tiket masuk perguruan tinggi negeri favorit dan universitas di Timur Tengah.</p>`,
			ThumbnailUrl: "https://images.unsplash.com/photo-1523240795612-9a054b0db644?auto=format&fit=crop&w=1000&q=80",
			Author:       "Tim Humas Al-Hamid",
			Status:       "published",
			IsFeatured:   true,
			ViewCount:    580,
			PublishedAt:  time.Now().Add(-24 * time.Hour),
		},
		{
			Title:        "Informasi Alur Pendaftaran Santri Baru (PSB) Tahun Ajaran 2026/2027",
			Slug:         "alur-pendaftaran-santri-baru-2026-2027",
			CategoryID:   catPengumuman.ID,
			Excerpt:      "Pendaftaran Santri Baru gelombang I resmi dibuka mulai 1 Maret hingga 30 Mei 2026. Simak syarat, jadwal tes, dan tautan formulir pendaftaran.",
			Content: `<p>Panitia Penerimaan Santri Baru (PSB) Pondok Pesantren Al-Hamid mengumumkan bahwa pendaftaran untuk jenjang MTs Cintamulya, MA Cintamulya, MAI Cintamulya, serta Madrasah Diniyah dan PMQ telah dibuka.</p>
<h4>Gelombang Pendaftaran:</h4>
<ul>
<li><strong>Gelombang I</strong>: 1 Maret – 30 Mei 2026 (Free biaya formulir)</li>
<li><strong>Gelombang II</strong>: 1 Juni – 10 Juli 2026</li>
</ul>
<p>Seluruh proses verifikasi berkas dan simulasi tes dapat diakses melalui portal resmi atau langsung hadir ke sekretariat PSB di kampus Al-Hamid.</p>`,
			ThumbnailUrl: "https://images.unsplash.com/photo-1523050854058-8df90110c9f1?auto=format&fit=crop&w=1000&q=80",
			Author:       "Panitia PSB Al-Hamid",
			Status:       "published",
			IsFeatured:   false,
			ViewCount:    1205,
			PublishedAt:  time.Now(),
		},
	}

	for _, art := range articles {
		db.Create(&art)
	}
}

func seedMedia(db *gorm.DB) {
	var photoCount int64
	db.Model(&models.MediaPhoto{}).Count(&photoCount)
	if photoCount == 0 {
		var catFoto models.Category
		db.Where("slug = ?", "foto-kegiatan").First(&catFoto)

		photos := []models.MediaPhoto{
			{
				Title:      "Suasana Halaqoh Ba'da Shubuh",
				CategoryID: catFoto.ID,
				ImageUrl:   "https://images.unsplash.com/photo-1542838132-92c53300491e?auto=format&fit=crop&w=1000&q=80",
				Caption:    "Santri menyimak dengan khidmat kajian tafsir Al-Qur'an di serambi Masjid Jami'.",
			},
			{
				Title:      "Praktikum Sains di Lab MA Cintamulya",
				CategoryID: catFoto.ID,
				ImageUrl:   "https://images.unsplash.com/photo-1581092918056-0c4c3acd3789?auto=format&fit=crop&w=1000&q=80",
				Caption:    "Eksplorasi penelitian sains modern melengkapi bekal intelektual santri.",
			},
			{
				Title:      "Semarak Lomba Hadroh Antar-Asrama",
				CategoryID: catFoto.ID,
				ImageUrl:   "https://images.unsplash.com/photo-1511671782779-c97d3d27a1d4?auto=format&fit=crop&w=1000&q=80",
				Caption:    "Gita sholawat berkumandang memupuk kecintaan mendalam pada Baginda Nabi SAW.",
			},
			{
				Title:      "Latihan Debat Tiga Bahasa Santri",
				CategoryID: catFoto.ID,
				ImageUrl:   "https://images.unsplash.com/photo-1524178232363-1fb2b075b655?auto=format&fit=crop&w=1000&q=80",
				Caption:    "Melatih keberanian berbicara dan diplomasi santri berwawasan global.",
			},
		}

		for _, p := range photos {
			db.Create(&p)
		}
	}

	var videoCount int64
	db.Model(&models.MediaVideo{}).Count(&videoCount)
	if videoCount == 0 {
		var catVideo models.Category
		db.Where("slug = ?", "video-profil").First(&catVideo)

		videos := []models.MediaVideo{
			{
				Title:       "Profil Resmi Pondok Pesantren Al-Hamid",
				CategoryID:  catVideo.ID,
				YoutubeUrl:  "https://www.youtube.com/watch?v=dQw4w9WgXcQ",
				YoutubeId:   "dQw4w9WgXcQ",
				Description: "Dokumenter sinematik perjalanan dan filosofi pendidikan di Pondok Pesantren Al-Hamid.",
			},
			{
				Title:       "Kajian Akbar Kitab Al-Hikam oleh Pengasuh Utama",
				CategoryID:  catVideo.ID,
				YoutubeUrl:  "https://www.youtube.com/watch?v=kJQP7kiw5Fk",
				YoutubeId:   "kJQP7kiw5Fk",
				Description: "Kajian tasawuf mendalam yang menyentuh hati dan menuntun ketenangan batin.",
			},
			{
				Title:       "Karya Santri MAI Cintamulya: Digitalisasi Pesantren",
				CategoryID:  catVideo.ID,
				YoutubeUrl:  "https://www.youtube.com/watch?v=fJ9rUzIMcZQ",
				YoutubeId:   "fJ9rUzIMcZQ",
				Description: "Pameran hasil karya software development dan animasi santri vokasi informatika.",
			},
		}

		for _, v := range videos {
			db.Create(&v)
		}
	}
}

