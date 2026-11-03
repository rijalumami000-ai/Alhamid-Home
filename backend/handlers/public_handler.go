package handlers

import (
	"strconv"
	"strings"

	"alhamid-backend/database"
	"alhamid-backend/models"

	"github.com/gofiber/fiber/v2"
)

type PublicHandler struct{}

func NewPublicHandler() *PublicHandler {
	return &PublicHandler{}
}

func (h *PublicHandler) GetHome(c *fiber.Ctx) error {
	var banners []models.Banner
	database.DB.Where("is_active = ?", true).Order("\"order\" ASC, id DESC").Find(&banners)

	var units []models.EducationUnit
	database.DB.Where("is_active = ?", true).Order("\"order\" ASC").Find(&units)

	var featuredArticles []models.Article
	database.DB.Preload("Category").Where("status = ?", "published").Order("is_featured DESC, published_at DESC").Limit(6).Find(&featuredArticles)

	var caregivers []models.Caregiver
	database.DB.Where("is_active = ?", true).Order("\"order\" ASC").Limit(4).Find(&caregivers)

	var facilities []models.Facility
	database.DB.Order("\"order\" ASC").Limit(6).Find(&facilities)

	var achievements []models.Achievement
	database.DB.Order("year DESC, id DESC").Limit(4).Find(&achievements)

	var testimonials []models.Testimonial
	database.DB.Where("is_active = ?", true).Order("\"order\" ASC").Limit(6).Find(&testimonials)

	var recentPhotos []models.MediaPhoto
	database.DB.Preload("Category").Order("id DESC").Limit(8).Find(&recentPhotos)

	// Site Settings
	var settings []models.Setting
	database.DB.Find(&settings)
	settingMap := make(map[string]string)
	for _, s := range settings {
		settingMap[s.Key] = s.Value
	}

	// Quick Stats
	stats := fiber.Map{
		"santri_count":    "1.450+",
		"asatidz_count":   "85+",
		"alumni_count":    "6.200+",
		"prestasi_count":  "120+",
	}

	return c.JSON(fiber.Map{
		"banners":           banners,
		"education_units":   units,
		"featured_articles": featuredArticles,
		"caregivers":        caregivers,
		"facilities":        facilities,
		"achievements":      achievements,
		"testimonials":      testimonials,
		"recent_photos":     recentPhotos,
		"stats":             stats,
		"settings":          settingMap,
	})
}

func (h *PublicHandler) GetSettings(c *fiber.Ctx) error {
	var settings []models.Setting
	database.DB.Find(&settings)

	settingMap := make(map[string]string)
	for _, s := range settings {
		settingMap[s.Key] = s.Value
	}

	return c.JSON(settingMap)
}

func (h *PublicHandler) GetAbout(c *fiber.Ctx) error {
	var settings []models.Setting
	database.DB.Where("category IN ?", []string{"general", "profile", "contact"}).Find(&settings)

	settingMap := make(map[string]string)
	for _, s := range settings {
		settingMap[s.Key] = s.Value
	}

	var caregivers []models.Caregiver
	database.DB.Where("is_active = ?", true).Order("\"order\" ASC").Find(&caregivers)

	var facilities []models.Facility
	database.DB.Order("\"order\" ASC").Find(&facilities)

	var achievements []models.Achievement
	database.DB.Order("year DESC, id DESC").Find(&achievements)

	var testimonials []models.Testimonial
	database.DB.Where("is_active = ?", true).Order("\"order\" ASC").Find(&testimonials)

	var alumni []models.Alumni
	database.DB.Order("graduation_year DESC, \"order\" ASC").Find(&alumni)

	return c.JSON(fiber.Map{
		"profile":       settingMap["about_profile"],
		"profile_image": settingMap["about_profile_image"],
		"vision":        settingMap["about_vision"],
		"mission":       settingMap["about_mission"],
		"site_name":     settingMap["site_name"],
		"site_tagline":  settingMap["site_tagline"],
		"caregivers":   caregivers,
		"facilities":   facilities,
		"achievements": achievements,
		"testimonials": testimonials,
		"alumni":       alumni,
	})
}

func (h *PublicHandler) GetCaregivers(c *fiber.Ctx) error {
	var list []models.Caregiver
	database.DB.Where("is_active = ?", true).Order("\"order\" ASC").Find(&list)
	return c.JSON(list)
}

func (h *PublicHandler) GetFacilities(c *fiber.Ctx) error {
	var list []models.Facility
	database.DB.Order("\"order\" ASC").Find(&list)
	return c.JSON(list)
}

func (h *PublicHandler) GetAchievements(c *fiber.Ctx) error {
	var list []models.Achievement
	database.DB.Order("year DESC, id DESC").Find(&list)
	return c.JSON(list)
}

func (h *PublicHandler) GetTestimonials(c *fiber.Ctx) error {
	var list []models.Testimonial
	database.DB.Where("is_active = ?", true).Order("\"order\" ASC").Find(&list)
	return c.JSON(list)
}

func (h *PublicHandler) GetAlumni(c *fiber.Ctx) error {
	var list []models.Alumni
	database.DB.Order("graduation_year DESC, \"order\" ASC").Find(&list)
	return c.JSON(list)
}

func (h *PublicHandler) GetEducationUnits(c *fiber.Ctx) error {
	unitType := c.Query("type") // formal / non_formal
	var list []models.EducationUnit

	query := database.DB.Where("is_active = ?", true)
	if unitType != "" {
		query = query.Where("type = ?", unitType)
	}

	query.Order("\"order\" ASC").Find(&list)
	return c.JSON(list)
}

func (h *PublicHandler) GetEducationUnitBySlug(c *fiber.Ctx) error {
	slug := c.Params("slug")
	var unit models.EducationUnit

	if err := database.DB.Where("slug = ?", slug).First(&unit).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Laman pendidikan tidak ditemukan"})
	}

	return c.JSON(unit)
}

func (h *PublicHandler) GetCategories(c *fiber.Ctx) error {
	catType := c.Query("type", "berita")
	var categories []models.Category

	query := database.DB
	if catType != "all" {
		query = query.Where("type = ?", catType)
	}
	query.Order("name ASC").Find(&categories)

	return c.JSON(categories)
}

func (h *PublicHandler) GetArticles(c *fiber.Ctx) error {
	categorySlug := c.Query("category")
	search := c.Query("q")
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "9"))
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 50 {
		limit = 9
	}
	offset := (page - 1) * limit

	query := database.DB.Model(&models.Article{}).Preload("Category").Where("status = ?", "published")

	if categorySlug != "" {
		var cat models.Category
		if err := database.DB.Where("slug = ?", categorySlug).First(&cat).Error; err == nil {
			query = query.Where("category_id = ?", cat.ID)
		}
	}

	if search != "" {
		likeTerm := "%" + strings.ToLower(search) + "%"
		query = query.Where("LOWER(title) LIKE ? OR LOWER(excerpt) LIKE ? OR LOWER(content) LIKE ?", likeTerm, likeTerm, likeTerm)
	}

	var total int64
	query.Count(&total)

	var articles []models.Article
	query.Order("published_at DESC").Limit(limit).Offset(offset).Find(&articles)

	return c.JSON(fiber.Map{
		"data":  articles,
		"total": total,
		"page":  page,
		"limit": limit,
		"pages": (total + int64(limit) - 1) / int64(limit),
	})
}

func (h *PublicHandler) GetArticleBySlug(c *fiber.Ctx) error {
	slug := c.Params("slug")
	var article models.Article

	if err := database.DB.Preload("Category").Where("slug = ? AND status = ?", slug, "published").First(&article).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Artikel tidak ditemukan"})
	}

	// Increment view count
	database.DB.Model(&article).UpdateColumn("view_count", article.ViewCount+1)

	// Get related articles
	var related []models.Article
	database.DB.Preload("Category").
		Where("category_id = ? AND id != ? AND status = ?", article.CategoryID, article.ID, "published").
		Order("published_at DESC").Limit(3).Find(&related)

	return c.JSON(fiber.Map{
		"article": article,
		"related": related,
	})
}

func (h *PublicHandler) GetPhotoGallery(c *fiber.Ctx) error {
	catSlug := c.Query("category")
	var photos []models.MediaPhoto

	query := database.DB.Preload("Category")
	if catSlug != "" {
		var cat models.Category
		if err := database.DB.Where("slug = ?", catSlug).First(&cat).Error; err == nil {
			query = query.Where("category_id = ?", cat.ID)
		}
	}

	query.Order("\"order\" ASC, id DESC").Find(&photos)
	return c.JSON(photos)
}

func (h *PublicHandler) GetVideoGallery(c *fiber.Ctx) error {
	catSlug := c.Query("category")
	var videos []models.MediaVideo

	query := database.DB.Preload("Category")
	if catSlug != "" {
		var cat models.Category
		if err := database.DB.Where("slug = ?", catSlug).First(&cat).Error; err == nil {
			query = query.Where("category_id = ?", cat.ID)
		}
	}

	query.Order("\"order\" ASC, id DESC").Find(&videos)
	return c.JSON(videos)
}

func (h *PublicHandler) SubmitContact(c *fiber.Ctx) error {
	var msg models.ContactMessage
	if err := c.BodyParser(&msg); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Format formulir pesan tidak valid"})
	}

	if msg.Name == "" || msg.Email == "" || msg.Message == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Nama, email, dan isi pesan wajib diisi"})
	}

	if err := database.DB.Create(&msg).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Gagal menyimpan pesan Anda"})
	}

	return c.JSON(fiber.Map{"message": "Terima kasih! Pesan Anda telah kami terima dan akan segera ditindaklanjuti."})
}

func (h *PublicHandler) GlobalSearch(c *fiber.Ctx) error {
	q := strings.TrimSpace(c.Query("q"))
	if len(q) < 2 {
		return c.JSON(fiber.Map{
			"articles":        []models.Article{},
			"education_units": []models.EducationUnit{},
			"caregivers":      []models.Caregiver{},
			"facilities":      []models.Facility{},
		})
	}

	term := "%" + strings.ToLower(q) + "%"

	var articles []models.Article
	database.DB.Preload("Category").
		Where("status = ? AND (LOWER(title) LIKE ? OR LOWER(excerpt) LIKE ?)", "published", term, term).
		Limit(5).Find(&articles)

	var units []models.EducationUnit
	database.DB.Where("is_active = ? AND (LOWER(name) LIKE ? OR LOWER(short_desc) LIKE ?)", true, term, term).
		Limit(5).Find(&units)

	var caregivers []models.Caregiver
	database.DB.Where("is_active = ? AND (LOWER(name) LIKE ? OR LOWER(role_title) LIKE ?)", true, term, term).
		Limit(4).Find(&caregivers)

	var facilities []models.Facility
	database.DB.Where("LOWER(name) LIKE ? OR LOWER(description) LIKE ?", term, term).
		Limit(4).Find(&facilities)

	return c.JSON(fiber.Map{
		"articles":        articles,
		"education_units": units,
		"caregivers":      caregivers,
		"facilities":      facilities,
	})
}

