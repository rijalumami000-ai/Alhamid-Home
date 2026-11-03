package handlers

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"alhamid-backend/database"
	"alhamid-backend/models"

	"github.com/gofiber/fiber/v2"
)

type AdminHandler struct{}

func NewAdminHandler() *AdminHandler {
	return &AdminHandler{}
}

func slugify(s string) string {
	re := regexp.MustCompile(`[^a-z0-9]+`)
	slug := strings.ToLower(s)
	slug = re.ReplaceAllString(slug, "-")
	slug = strings.Trim(slug, "-")
	if slug == "" {
		slug = fmt.Sprintf("item-%d", time.Now().Unix())
	}
	return slug
}

func normalizeGoogleDriveURL(raw string) string {
	raw = strings.TrimSpace(raw)
	if raw == "" {
		return ""
	}

	// 1. /file/d/FILE_ID
	reFile := regexp.MustCompile(`/file/d/([a-zA-Z0-9_-]+)`)
	if match := reFile.FindStringSubmatch(raw); len(match) > 1 {
		return fmt.Sprintf("https://lh3.googleusercontent.com/d/%s", match[1])
	}

	// 2. id=FILE_ID
	reID := regexp.MustCompile(`[?&]id=([a-zA-Z0-9_-]+)`)
	if match := reID.FindStringSubmatch(raw); len(match) > 1 {
		return fmt.Sprintf("https://lh3.googleusercontent.com/d/%s", match[1])
	}

	// 3. googleusercontent.com/d/FILE_ID
	reGuc := regexp.MustCompile(`googleusercontent\.com/d/([a-zA-Z0-9_-]+)`)
	if match := reGuc.FindStringSubmatch(raw); len(match) > 1 {
		return fmt.Sprintf("https://lh3.googleusercontent.com/d/%s", match[1])
	}

	return raw
}

// ---------------- DASHBOARD ----------------
func (h *AdminHandler) GetDashboardStats(c *fiber.Ctx) error {
	var totalArticles, totalUnits, totalPhotos, totalVideos, unreadMessages, totalCaregivers int64

	database.DB.Model(&models.Article{}).Count(&totalArticles)
	database.DB.Model(&models.EducationUnit{}).Count(&totalUnits)
	database.DB.Model(&models.MediaPhoto{}).Count(&totalPhotos)
	database.DB.Model(&models.MediaVideo{}).Count(&totalVideos)
	database.DB.Model(&models.ContactMessage{}).Where("is_read = ?", false).Count(&unreadMessages)
	database.DB.Model(&models.Caregiver{}).Count(&totalCaregivers)

	var recentArticles []models.Article
	database.DB.Preload("Category").Order("created_at DESC").Limit(5).Find(&recentArticles)

	var recentMessages []models.ContactMessage
	database.DB.Order("created_at DESC").Limit(5).Find(&recentMessages)

	return c.JSON(fiber.Map{
		"stats": fiber.Map{
			"total_articles":   totalArticles,
			"total_units":      totalUnits,
			"total_photos":     totalPhotos,
			"total_videos":     totalVideos,
			"unread_messages":  unreadMessages,
			"total_caregivers": totalCaregivers,
		},
		"recent_articles": recentArticles,
		"recent_messages": recentMessages,
	})
}

// ---------------- ARTICLES ----------------
func (h *AdminHandler) GetArticles(c *fiber.Ctx) error {
	var articles []models.Article
	database.DB.Preload("Category").Order("created_at DESC").Find(&articles)
	return c.JSON(articles)
}

func (h *AdminHandler) CreateArticle(c *fiber.Ctx) error {
	var article models.Article
	if err := c.BodyParser(&article); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Format data artikel tidak valid"})
	}

	if article.Title == "" || article.Content == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Judul dan konten artikel wajib diisi"})
	}

	if article.Slug == "" {
		article.Slug = slugify(article.Title)
	}

	// Cek keunikan slug
	var existing models.Article
	if err := database.DB.Where("slug = ?", article.Slug).First(&existing).Error; err == nil {
		article.Slug = fmt.Sprintf("%s-%d", article.Slug, time.Now().Unix())
	}

	if article.PublishedAt.IsZero() && article.Status == "published" {
		article.PublishedAt = time.Now()
	}

	article.ThumbnailUrl = normalizeGoogleDriveURL(article.ThumbnailUrl)

	if err := database.DB.Create(&article).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Gagal menyimpan artikel"})
	}

	database.DB.Preload("Category").First(&article, article.ID)
	return c.Status(fiber.StatusCreated).JSON(article)
}

func (h *AdminHandler) UpdateArticle(c *fiber.Ctx) error {
	id := c.Params("id")
	var article models.Article
	if err := database.DB.First(&article, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Artikel tidak ditemukan"})
	}

	var input models.Article
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Format data artikel tidak valid"})
	}

	article.Title = input.Title
	if input.Slug != "" {
		article.Slug = input.Slug
	}
	article.CategoryID = input.CategoryID
	article.Excerpt = input.Excerpt
	article.Content = input.Content
	article.ThumbnailUrl = normalizeGoogleDriveURL(input.ThumbnailUrl)
	article.Author = input.Author
	article.Status = input.Status
	article.IsFeatured = input.IsFeatured

	if input.Status == "published" && article.PublishedAt.IsZero() {
		article.PublishedAt = time.Now()
	}

	if err := database.DB.Save(&article).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Gagal memperbarui artikel"})
	}

	database.DB.Preload("Category").First(&article, article.ID)
	return c.JSON(article)
}

func (h *AdminHandler) DeleteArticle(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := database.DB.Delete(&models.Article{}, id).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Gagal menghapus artikel"})
	}
	return c.JSON(fiber.Map{"message": "Artikel berhasil dihapus"})
}

// ---------------- CATEGORIES ----------------
func (h *AdminHandler) GetCategories(c *fiber.Ctx) error {
	var categories []models.Category
	database.DB.Order("type ASC, name ASC").Find(&categories)
	return c.JSON(categories)
}

func (h *AdminHandler) CreateCategory(c *fiber.Ctx) error {
	var cat models.Category
	if err := c.BodyParser(&cat); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Format data kategori tidak valid"})
	}
	if cat.Slug == "" {
		cat.Slug = slugify(cat.Name)
	}
	if err := database.DB.Create(&cat).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Gagal menyimpan kategori"})
	}
	return c.Status(fiber.StatusCreated).JSON(cat)
}

func (h *AdminHandler) UpdateCategory(c *fiber.Ctx) error {
	id := c.Params("id")
	var cat models.Category
	if err := database.DB.First(&cat, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Kategori tidak ditemukan"})
	}
	if err := c.BodyParser(&cat); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Format data kategori tidak valid"})
	}
	database.DB.Save(&cat)
	return c.JSON(cat)
}

func (h *AdminHandler) DeleteCategory(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := database.DB.Delete(&models.Category{}, id).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Gagal menghapus kategori"})
	}
	return c.JSON(fiber.Map{"message": "Kategori berhasil dihapus"})
}

// ---------------- EDUCATION UNITS ----------------
func (h *AdminHandler) GetEducationUnits(c *fiber.Ctx) error {
	var units []models.EducationUnit
	database.DB.Order("\"order\" ASC").Find(&units)
	return c.JSON(units)
}

func (h *AdminHandler) CreateEducationUnit(c *fiber.Ctx) error {
	var unit models.EducationUnit
	if err := c.BodyParser(&unit); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Format data tidak valid"})
	}
	if unit.Slug == "" {
		unit.Slug = slugify(unit.Name)
	}
	unit.HeroImage = normalizeGoogleDriveURL(unit.HeroImage)
	unit.HeadmasterPhoto = normalizeGoogleDriveURL(unit.HeadmasterPhoto)
	if err := database.DB.Create(&unit).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Gagal menambah unit pendidikan"})
	}
	return c.Status(fiber.StatusCreated).JSON(unit)
}

func (h *AdminHandler) UpdateEducationUnit(c *fiber.Ctx) error {
	id := c.Params("id")
	var unit models.EducationUnit
	if err := database.DB.First(&unit, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Unit pendidikan tidak ditemukan"})
	}
	if err := c.BodyParser(&unit); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Format data tidak valid"})
	}
	unit.HeroImage = normalizeGoogleDriveURL(unit.HeroImage)
	unit.HeadmasterPhoto = normalizeGoogleDriveURL(unit.HeadmasterPhoto)
	database.DB.Save(&unit)
	return c.JSON(unit)
}

func (h *AdminHandler) DeleteEducationUnit(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := database.DB.Delete(&models.EducationUnit{}, id).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Gagal menghapus unit pendidikan"})
	}
	return c.JSON(fiber.Map{"message": "Unit pendidikan berhasil dihapus"})
}

// ---------------- CAREGIVERS (DEWAN PENGASUH) ----------------
func (h *AdminHandler) GetCaregivers(c *fiber.Ctx) error {
	var list []models.Caregiver
	database.DB.Order("\"order\" ASC").Find(&list)
	return c.JSON(list)
}

func (h *AdminHandler) CreateCaregiver(c *fiber.Ctx) error {
	var item models.Caregiver
	if err := c.BodyParser(&item); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Format data tidak valid"})
	}
	item.PhotoUrl = normalizeGoogleDriveURL(item.PhotoUrl)
	if err := database.DB.Create(&item).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Gagal menambah dewan pengasuh"})
	}
	return c.Status(fiber.StatusCreated).JSON(item)
}

func (h *AdminHandler) UpdateCaregiver(c *fiber.Ctx) error {
	id := c.Params("id")
	var item models.Caregiver
	if err := database.DB.First(&item, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Data pengasuh tidak ditemukan"})
	}
	if err := c.BodyParser(&item); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Format data tidak valid"})
	}
	item.PhotoUrl = normalizeGoogleDriveURL(item.PhotoUrl)
	database.DB.Save(&item)
	return c.JSON(item)
}

func (h *AdminHandler) DeleteCaregiver(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := database.DB.Delete(&models.Caregiver{}, id).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Gagal menghapus data pengasuh"})
	}
	return c.JSON(fiber.Map{"message": "Data pengasuh berhasil dihapus"})
}

// ---------------- FACILITIES ----------------
func (h *AdminHandler) GetFacilities(c *fiber.Ctx) error {
	var list []models.Facility
	database.DB.Order("\"order\" ASC").Find(&list)
	return c.JSON(list)
}

func (h *AdminHandler) CreateFacility(c *fiber.Ctx) error {
	var item models.Facility
	if err := c.BodyParser(&item); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Format data tidak valid"})
	}
	item.ImageUrl = normalizeGoogleDriveURL(item.ImageUrl)
	database.DB.Create(&item)
	return c.Status(fiber.StatusCreated).JSON(item)
}

func (h *AdminHandler) UpdateFacility(c *fiber.Ctx) error {
	id := c.Params("id")
	var item models.Facility
	if err := database.DB.First(&item, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Data fasilitas tidak ditemukan"})
	}
	if err := c.BodyParser(&item); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Format data tidak valid"})
	}
	item.ImageUrl = normalizeGoogleDriveURL(item.ImageUrl)
	database.DB.Save(&item)
	return c.JSON(item)
}

func (h *AdminHandler) DeleteFacility(c *fiber.Ctx) error {
	id := c.Params("id")
	database.DB.Delete(&models.Facility{}, id)
	return c.JSON(fiber.Map{"message": "Fasilitas berhasil dihapus"})
}

// ---------------- ACHIEVEMENTS ----------------
func (h *AdminHandler) GetAchievements(c *fiber.Ctx) error {
	var list []models.Achievement
	database.DB.Order("year DESC, id DESC").Find(&list)
	return c.JSON(list)
}

func (h *AdminHandler) CreateAchievement(c *fiber.Ctx) error {
	var item models.Achievement
	if err := c.BodyParser(&item); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Format data tidak valid"})
	}
	item.ImageUrl = normalizeGoogleDriveURL(item.ImageUrl)
	database.DB.Create(&item)
	return c.Status(fiber.StatusCreated).JSON(item)
}

func (h *AdminHandler) UpdateAchievement(c *fiber.Ctx) error {
	id := c.Params("id")
	var item models.Achievement
	if err := database.DB.First(&item, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Data prestasi tidak ditemukan"})
	}
	if err := c.BodyParser(&item); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Format data tidak valid"})
	}
	item.ImageUrl = normalizeGoogleDriveURL(item.ImageUrl)
	database.DB.Save(&item)
	return c.JSON(item)
}

func (h *AdminHandler) DeleteAchievement(c *fiber.Ctx) error {
	id := c.Params("id")
	database.DB.Delete(&models.Achievement{}, id)
	return c.JSON(fiber.Map{"message": "Prestasi berhasil dihapus"})
}

// ---------------- TESTIMONIALS ----------------
func (h *AdminHandler) GetTestimonials(c *fiber.Ctx) error {
	var list []models.Testimonial
	database.DB.Order("\"order\" ASC").Find(&list)
	return c.JSON(list)
}

func (h *AdminHandler) CreateTestimonial(c *fiber.Ctx) error {
	var item models.Testimonial
	if err := c.BodyParser(&item); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Format data tidak valid"})
	}
	item.AvatarUrl = normalizeGoogleDriveURL(item.AvatarUrl)
	database.DB.Create(&item)
	return c.Status(fiber.StatusCreated).JSON(item)
}

func (h *AdminHandler) UpdateTestimonial(c *fiber.Ctx) error {
	id := c.Params("id")
	var item models.Testimonial
	if err := database.DB.First(&item, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Testimoni tidak ditemukan"})
	}
	if err := c.BodyParser(&item); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Format data tidak valid"})
	}
	item.AvatarUrl = normalizeGoogleDriveURL(item.AvatarUrl)
	database.DB.Save(&item)
	return c.JSON(item)
}

func (h *AdminHandler) DeleteTestimonial(c *fiber.Ctx) error {
	id := c.Params("id")
	database.DB.Delete(&models.Testimonial{}, id)
	return c.JSON(fiber.Map{"message": "Testimoni berhasil dihapus"})
}

// ---------------- ALUMNI ----------------
func (h *AdminHandler) GetAlumni(c *fiber.Ctx) error {
	var list []models.Alumni
	database.DB.Order("graduation_year DESC, \"order\" ASC").Find(&list)
	return c.JSON(list)
}

func (h *AdminHandler) CreateAlumni(c *fiber.Ctx) error {
	var item models.Alumni
	if err := c.BodyParser(&item); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Format data tidak valid"})
	}
	item.PhotoUrl = normalizeGoogleDriveURL(item.PhotoUrl)
	database.DB.Create(&item)
	return c.Status(fiber.StatusCreated).JSON(item)
}

func (h *AdminHandler) UpdateAlumni(c *fiber.Ctx) error {
	id := c.Params("id")
	var item models.Alumni
	if err := database.DB.First(&item, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Data alumni tidak ditemukan"})
	}
	if err := c.BodyParser(&item); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Format data tidak valid"})
	}
	item.PhotoUrl = normalizeGoogleDriveURL(item.PhotoUrl)
	database.DB.Save(&item)
	return c.JSON(item)
}

func (h *AdminHandler) DeleteAlumni(c *fiber.Ctx) error {
	id := c.Params("id")
	database.DB.Delete(&models.Alumni{}, id)
	return c.JSON(fiber.Map{"message": "Data alumni berhasil dihapus"})
}

// ---------------- MEDIA PHOTOS & VIDEOS ----------------
func (h *AdminHandler) GetPhotos(c *fiber.Ctx) error {
	var list []models.MediaPhoto
	database.DB.Preload("Category").Order("id DESC").Find(&list)
	return c.JSON(list)
}

func (h *AdminHandler) CreatePhoto(c *fiber.Ctx) error {
	var item models.MediaPhoto
	if err := c.BodyParser(&item); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Format data tidak valid"})
	}
	item.ImageUrl = normalizeGoogleDriveURL(item.ImageUrl)
	database.DB.Create(&item)
	database.DB.Preload("Category").First(&item, item.ID)
	return c.Status(fiber.StatusCreated).JSON(item)
}

func (h *AdminHandler) DeletePhoto(c *fiber.Ctx) error {
	id := c.Params("id")
	database.DB.Delete(&models.MediaPhoto{}, id)
	return c.JSON(fiber.Map{"message": "Foto berhasil dihapus"})
}

func (h *AdminHandler) GetVideos(c *fiber.Ctx) error {
	var list []models.MediaVideo
	database.DB.Preload("Category").Order("id DESC").Find(&list)
	return c.JSON(list)
}

func (h *AdminHandler) CreateVideo(c *fiber.Ctx) error {
	var item models.MediaVideo
	if err := c.BodyParser(&item); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Format data tidak valid"})
	}

	// Extract YouTube ID if not given
	if item.YoutubeId == "" && item.YoutubeUrl != "" {
		re := regexp.MustCompile(`(?:v=|\/)([0-9A-Za-z_-]{11}).*`)
		matches := re.FindStringSubmatch(item.YoutubeUrl)
		if len(matches) > 1 {
			item.YoutubeId = matches[1]
		}
	}

	database.DB.Create(&item)
	database.DB.Preload("Category").First(&item, item.ID)
	return c.Status(fiber.StatusCreated).JSON(item)
}

func (h *AdminHandler) DeleteVideo(c *fiber.Ctx) error {
	id := c.Params("id")
	database.DB.Delete(&models.MediaVideo{}, id)
	return c.JSON(fiber.Map{"message": "Video berhasil dihapus"})
}

// ---------------- BANNERS ----------------
func (h *AdminHandler) GetBanners(c *fiber.Ctx) error {
	var banners []models.Banner
	database.DB.Order("\"order\" ASC").Find(&banners)
	return c.JSON(banners)
}

func (h *AdminHandler) CreateBanner(c *fiber.Ctx) error {
	var b models.Banner
	if err := c.BodyParser(&b); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Format data banner tidak valid"})
	}
	b.ImageUrl = normalizeGoogleDriveURL(b.ImageUrl)
	database.DB.Create(&b)
	return c.Status(fiber.StatusCreated).JSON(b)
}

func (h *AdminHandler) UpdateBanner(c *fiber.Ctx) error {
	id := c.Params("id")
	var b models.Banner
	if err := database.DB.First(&b, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Banner tidak ditemukan"})
	}
	if err := c.BodyParser(&b); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Format data banner tidak valid"})
	}
	b.ImageUrl = normalizeGoogleDriveURL(b.ImageUrl)
	database.DB.Save(&b)
	return c.JSON(b)
}

func (h *AdminHandler) DeleteBanner(c *fiber.Ctx) error {
	id := c.Params("id")
	database.DB.Delete(&models.Banner{}, id)
	return c.JSON(fiber.Map{"message": "Banner berhasil dihapus"})
}

// ---------------- SETTINGS ----------------
func (h *AdminHandler) GetSettings(c *fiber.Ctx) error {
	var settings []models.Setting
	database.DB.Find(&settings)
	return c.JSON(settings)
}

func (h *AdminHandler) UpdateSettings(c *fiber.Ctx) error {
	var payload map[string]string
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Format data tidak valid"})
	}

	for key, val := range payload {
		var s models.Setting
		if err := database.DB.Where("key = ?", key).First(&s).Error; err == nil {
			s.Value = val
			database.DB.Save(&s)
		} else {
			database.DB.Create(&models.Setting{
				Key:   key,
				Value: val,
			})
		}
	}

	return c.JSON(fiber.Map{"message": "Pengaturan website berhasil diperbarui"})
}

// ---------------- CONTACT MESSAGES ----------------
func (h *AdminHandler) GetMessages(c *fiber.Ctx) error {
	var list []models.ContactMessage
	database.DB.Order("created_at DESC").Find(&list)
	return c.JSON(list)
}

func (h *AdminHandler) MarkMessageRead(c *fiber.Ctx) error {
	id := c.Params("id")
	database.DB.Model(&models.ContactMessage{}).Where("id = ?", id).Update("is_read", true)
	return c.JSON(fiber.Map{"message": "Pesan ditandai telah dibaca"})
}

func (h *AdminHandler) DeleteMessage(c *fiber.Ctx) error {
	id := c.Params("id")
	database.DB.Delete(&models.ContactMessage{}, id)
	return c.JSON(fiber.Map{"message": "Pesan berhasil dihapus"})
}

