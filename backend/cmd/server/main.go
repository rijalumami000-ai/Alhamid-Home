package main

import (
	"fmt"
	"log"
	"os"

	"alhamid-backend/config"
	"alhamid-backend/database"
	"alhamid-backend/handlers"
	"alhamid-backend/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	cfg := config.LoadConfig()

	// Initialize Database & Run Migrations & Seeds
	database.InitDB(cfg)

	// Ensure upload dir exists
	_ = os.MkdirAll(cfg.UploadDir, 0755)

	app := fiber.New(fiber.Config{
		AppName:      "Al-Hamid Pesantren API & CMS v1.0",
		ServerHeader: "Al-Hamid-Server",
		BodyLimit:    20 * 1024 * 1024, // 20 MB
	})

	app.Use(recover.New())
	app.Use(logger.New(logger.Config{
		Format: "[${time}] ${status} - ${method} ${path} (${latency})\n",
	}))

	// CORS Setup
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:5173, http://127.0.0.1:5173, " + cfg.ClientURL,
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowMethods:     "GET, POST, PUT, DELETE, OPTIONS",
		AllowCredentials: true,
	}))

	// Static files for uploads
	app.Static("/uploads", cfg.UploadDir)

	// Handlers
	authHandler := handlers.NewAuthHandler(cfg)
	publicHandler := handlers.NewPublicHandler()
	adminHandler := handlers.NewAdminHandler()
	uploadHandler := handlers.NewUploadHandler(cfg)

	// API Route Group
	api := app.Group("/api")

	// 1. Health Check
	api.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  "healthy",
			"service": "Pondok Pesantren Al-Hamid Backend",
			"version": "1.0.0",
		})
	})

	// 2. Public Endpoints
	public := api.Group("/public")
	public.Get("/home", publicHandler.GetHome)
	public.Get("/settings", publicHandler.GetSettings)
	public.Get("/about", publicHandler.GetAbout)
	public.Get("/caregivers", publicHandler.GetCaregivers)
	public.Get("/facilities", publicHandler.GetFacilities)
	public.Get("/achievements", publicHandler.GetAchievements)
	public.Get("/testimonials", publicHandler.GetTestimonials)
	public.Get("/alumni", publicHandler.GetAlumni)
	public.Get("/education-units", publicHandler.GetEducationUnits)
	public.Get("/education-units/:slug", publicHandler.GetEducationUnitBySlug)
	public.Get("/categories", publicHandler.GetCategories)
	public.Get("/articles", publicHandler.GetArticles)
	public.Get("/articles/:slug", publicHandler.GetArticleBySlug)
	public.Get("/gallery/photos", publicHandler.GetPhotoGallery)
	public.Get("/gallery/videos", publicHandler.GetVideoGallery)
	public.Post("/contact", publicHandler.SubmitContact)
	public.Get("/search", publicHandler.GlobalSearch)

	// 3. Auth Endpoints
	auth := api.Group("/auth")
	auth.Post("/login", authHandler.Login)
	auth.Get("/me", middleware.Protected(cfg), authHandler.Me)
	auth.Post("/logout", authHandler.Logout)

	// 4. Admin Protected Endpoints
	admin := api.Group("/admin", middleware.Protected(cfg))
	admin.Get("/dashboard", adminHandler.GetDashboardStats)

	// Articles
	admin.Get("/articles", adminHandler.GetArticles)
	admin.Post("/articles", adminHandler.CreateArticle)
	admin.Put("/articles/:id", adminHandler.UpdateArticle)
	admin.Delete("/articles/:id", adminHandler.DeleteArticle)

	// Categories
	admin.Get("/categories", adminHandler.GetCategories)
	admin.Post("/categories", adminHandler.CreateCategory)
	admin.Put("/categories/:id", adminHandler.UpdateCategory)
	admin.Delete("/categories/:id", adminHandler.DeleteCategory)

	// Education Units
	admin.Get("/education-units", adminHandler.GetEducationUnits)
	admin.Post("/education-units", adminHandler.CreateEducationUnit)
	admin.Put("/education-units/:id", adminHandler.UpdateEducationUnit)
	admin.Delete("/education-units/:id", adminHandler.DeleteEducationUnit)

	// Caregivers
	admin.Get("/caregivers", adminHandler.GetCaregivers)
	admin.Post("/caregivers", adminHandler.CreateCaregiver)
	admin.Put("/caregivers/:id", adminHandler.UpdateCaregiver)
	admin.Delete("/caregivers/:id", adminHandler.DeleteCaregiver)

	// Facilities
	admin.Get("/facilities", adminHandler.GetFacilities)
	admin.Post("/facilities", adminHandler.CreateFacility)
	admin.Put("/facilities/:id", adminHandler.UpdateFacility)
	admin.Delete("/facilities/:id", adminHandler.DeleteFacility)

	// Achievements
	admin.Get("/achievements", adminHandler.GetAchievements)
	admin.Post("/achievements", adminHandler.CreateAchievement)
	admin.Put("/achievements/:id", adminHandler.UpdateAchievement)
	admin.Delete("/achievements/:id", adminHandler.DeleteAchievement)

	// Testimonials
	admin.Get("/testimonials", adminHandler.GetTestimonials)
	admin.Post("/testimonials", adminHandler.CreateTestimonial)
	admin.Put("/testimonials/:id", adminHandler.UpdateTestimonial)
	admin.Delete("/testimonials/:id", adminHandler.DeleteTestimonial)

	// Alumni
	admin.Get("/alumni", adminHandler.GetAlumni)
	admin.Post("/alumni", adminHandler.CreateAlumni)
	admin.Put("/alumni/:id", adminHandler.UpdateAlumni)
	admin.Delete("/alumni/:id", adminHandler.DeleteAlumni)

	// Media (Photos & Videos)
	admin.Get("/photos", adminHandler.GetPhotos)
	admin.Post("/photos", adminHandler.CreatePhoto)
	admin.Delete("/photos/:id", adminHandler.DeletePhoto)

	admin.Get("/videos", adminHandler.GetVideos)
	admin.Post("/videos", adminHandler.CreateVideo)
	admin.Delete("/videos/:id", adminHandler.DeleteVideo)

	// Banners
	admin.Get("/banners", adminHandler.GetBanners)
	admin.Post("/banners", adminHandler.CreateBanner)
	admin.Put("/banners/:id", adminHandler.UpdateBanner)
	admin.Delete("/banners/:id", adminHandler.DeleteBanner)

	// Settings
	admin.Get("/settings", adminHandler.GetSettings)
	admin.Put("/settings", adminHandler.UpdateSettings)

	// Contact Messages
	admin.Get("/messages", adminHandler.GetMessages)
	admin.Put("/messages/:id/read", adminHandler.MarkMessageRead)
	admin.Delete("/messages/:id", adminHandler.DeleteMessage)

	// Upload
	admin.Post("/upload", uploadHandler.UploadFile)

	addr := fmt.Sprintf(":%s", cfg.Port)
	log.Printf("[AL-HAMID SERVER] Server berjalan di http://localhost%s\n", addr)
	if err := app.Listen(addr); err != nil {
		log.Fatalf("Server shutdown error: %v", err)
	}
}

