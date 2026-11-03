package models

import (
	"time"
)

type User struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	Username     string    `gorm:"uniqueIndex;size:100;not null" json:"username"`
	Email        string    `gorm:"uniqueIndex;size:150;not null" json:"email"`
	PasswordHash string    `gorm:"not null" json:"-"`
	FullName     string    `gorm:"size:150" json:"full_name"`
	Role         string    `gorm:"size:50;default:'admin'" json:"role"` // admin, editor
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type Setting struct {
	Key       string    `gorm:"primaryKey;size:100" json:"key"`
	Value     string    `gorm:"type:text" json:"value"`
	Category  string    `gorm:"size:50;default:'general'" json:"category"` // general, profile, social, contact, seo
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Banner struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Title       string    `gorm:"size:255;not null" json:"title"`
	Subtitle    string    `gorm:"type:text" json:"subtitle"`
	ImageUrl    string    `gorm:"size:500;not null" json:"image_url"`
	ButtonText  string    `gorm:"size:100" json:"button_text"`
	ButtonLink  string    `gorm:"size:255" json:"button_link"`
	Badge       string    `gorm:"size:100" json:"badge"`
	Order       int       `gorm:"default:0" json:"order"`
	IsActive    bool      `gorm:"default:true" json:"is_active"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Category struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Name        string    `gorm:"size:100;not null" json:"name"`
	Slug        string    `gorm:"uniqueIndex;size:120;not null" json:"slug"`
	Type        string    `gorm:"size:50;default:'berita'" json:"type"` // berita, galeri_foto, galeri_video, prestasi
	Description string    `gorm:"type:text" json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Article struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	Title        string    `gorm:"size:255;not null" json:"title"`
	Slug         string    `gorm:"uniqueIndex;size:280;not null" json:"slug"`
	CategoryID   uint      `gorm:"not null" json:"category_id"`
	Category     Category  `gorm:"foreignKey:CategoryID" json:"category"`
	Excerpt      string    `gorm:"type:text" json:"excerpt"`
	Content      string    `gorm:"type:text;not null" json:"content"`
	ThumbnailUrl string    `gorm:"size:500" json:"thumbnail_url"`
	Author       string    `gorm:"size:100;default:'Redaksi Al-Hamid'" json:"author"`
	Status       string    `gorm:"size:50;default:'published'" json:"status"` // published, draft
	IsFeatured   bool      `gorm:"default:false" json:"is_featured"`
	ViewCount    int       `gorm:"default:0" json:"view_count"`
	PublishedAt  time.Time `json:"published_at"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type EducationUnit struct {
	ID              uint      `gorm:"primaryKey" json:"id"`
	Name            string    `gorm:"size:150;not null" json:"name"`
	Slug            string    `gorm:"uniqueIndex;size:160;not null" json:"slug"`
	Type            string    `gorm:"size:50;not null" json:"type"` // formal, non_formal
	Tagline         string    `gorm:"size:255" json:"tagline"`
	HeroImage       string    `gorm:"size:500" json:"hero_image"`
	Badge           string    `gorm:"size:100" json:"badge"` // e.g. "Terakreditasi A", "Khas Salafiyah Modern"
	ShortDesc       string    `gorm:"type:text" json:"short_desc"`
	FullDesc        string    `gorm:"type:text" json:"full_desc"`
	Curriculum      string    `gorm:"type:text" json:"curriculum"`       // JSON or markdown
	Requirements    string    `gorm:"type:text" json:"requirements"`     // Syarat pendaftaran
	FacilitiesList  string    `gorm:"type:text" json:"facilities_list"`  // Fasilitas spesifik unit
	HeadmasterName  string    `gorm:"size:150" json:"headmaster_name"`
	HeadmasterPhoto string    `gorm:"size:500" json:"headmaster_photo"`
	Order           int       `gorm:"default:0" json:"order"`
	IsActive        bool      `gorm:"default:true" json:"is_active"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type Caregiver struct {
	ID            uint      `gorm:"primaryKey" json:"id"`
	Name          string    `gorm:"size:150;not null" json:"name"`
	RoleTitle     string    `gorm:"size:150;not null" json:"role_title"` // Pengasuh Pondok, Majelis Masyayikh, dll
	PhotoUrl      string    `gorm:"size:500" json:"photo_url"`
	Bio           string    `gorm:"type:text" json:"bio"`
	SanadKeilmuan string    `gorm:"type:text" json:"sanad_keilmuan"`
	Order         int       `gorm:"default:0" json:"order"`
	IsActive      bool      `gorm:"default:true" json:"is_active"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type Facility struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Name        string    `gorm:"size:150;not null" json:"name"`
	Category    string    `gorm:"size:100" json:"category"` // Asrama, Ibadah, Akademik, Penunjang, Olahraga
	Description string    `gorm:"type:text" json:"description"`
	ImageUrl    string    `gorm:"size:500;not null" json:"image_url"`
	Order       int       `gorm:"default:0" json:"order"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Achievement struct {
	ID            uint      `gorm:"primaryKey" json:"id"`
	Title         string    `gorm:"size:255;not null" json:"title"`
	Year          int       `gorm:"not null" json:"year"`
	Level         string    `gorm:"size:100" json:"level"` // Kabupaten, Provinsi, Nasional, Internasional
	Category      string    `gorm:"size:100" json:"category"` // MQK, Tahfidz, Sains, Bahasa, Olahraga
	RecipientName string    `gorm:"size:150" json:"recipient_name"`
	Description   string    `gorm:"type:text" json:"description"`
	ImageUrl      string    `gorm:"size:500" json:"image_url"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type Testimonial struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"size:150;not null" json:"name"`
	Role      string    `gorm:"size:150;not null" json:"role"` // Alumni 2020 / Wali Santri / Akademisi
	Content   string    `gorm:"type:text;not null" json:"content"`
	AvatarUrl string    `gorm:"size:500" json:"avatar_url"`
	Rating    int       `gorm:"default:5" json:"rating"`
	Order     int       `gorm:"default:0" json:"order"`
	IsActive  bool      `gorm:"default:true" json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Alumni struct {
	ID              uint      `gorm:"primaryKey" json:"id"`
	Name            string    `gorm:"size:150;not null" json:"name"`
	GraduationYear  int       `gorm:"not null" json:"graduation_year"`
	EducationUnit   string    `gorm:"size:100" json:"education_unit"` // MA Cintamulya, Madrasah Diniyah, dll.
	Profession      string    `gorm:"size:150" json:"profession"`
	CurrentLocation string    `gorm:"size:150" json:"current_location"`
	Story           string    `gorm:"type:text" json:"story"`
	PhotoUrl        string    `gorm:"size:500" json:"photo_url"`
	SocialLink      string    `gorm:"size:255" json:"social_link"`
	Order           int       `gorm:"default:0" json:"order"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type MediaPhoto struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	Title      string    `gorm:"size:200;not null" json:"title"`
	CategoryID uint      `gorm:"not null" json:"category_id"`
	Category   Category  `gorm:"foreignKey:CategoryID" json:"category"`
	ImageUrl   string    `gorm:"size:500;not null" json:"image_url"`
	Caption    string    `gorm:"type:text" json:"caption"`
	Order      int       `gorm:"default:0" json:"order"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type MediaVideo struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Title       string    `gorm:"size:200;not null" json:"title"`
	CategoryID  uint      `gorm:"not null" json:"category_id"`
	Category    Category  `gorm:"foreignKey:CategoryID" json:"category"`
	YoutubeUrl  string    `gorm:"size:500;not null" json:"youtube_url"`
	YoutubeId   string    `gorm:"size:100;not null" json:"youtube_id"`
	Description string    `gorm:"type:text" json:"description"`
	Order       int       `gorm:"default:0" json:"order"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type ContactMessage struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"size:150;not null" json:"name"`
	Email     string    `gorm:"size:150;not null" json:"email"`
	Phone     string    `gorm:"size:50" json:"phone"`
	Subject   string    `gorm:"size:200;not null" json:"subject"`
	Message   string    `gorm:"type:text;not null" json:"message"`
	IsRead    bool      `gorm:"default:false" json:"is_read"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

