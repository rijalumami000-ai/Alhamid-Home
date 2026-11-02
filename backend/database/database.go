package database

import (
	"fmt"
	"log"

	"alhamid-backend/config"
	"alhamid-backend/models"

	"github.com/glebarez/sqlite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitDB(cfg *config.Config) *gorm.DB {
	var err error
	var db *gorm.DB

	if cfg.DBDriver == "postgres" {
		dsn := fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=Asia/Jakarta",
			cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort, cfg.DBSSLMode,
		)
		log.Printf("[DATABASE] Connecting to PostgreSQL at %s:%s (db: %s)...", cfg.DBHost, cfg.DBPort, cfg.DBName)
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Warn),
		})

		if err != nil {
			log.Printf("[DATABASE WARNING] Gagal koneksi ke PostgreSQL: %v", err)
			log.Println("[DATABASE INFO] Beralih otomatis ke SQLite (alhamid.db) agar aplikasi tetap berjalan di lokal!")
			db, err = gorm.Open(sqlite.Open("alhamid.db"), &gorm.Config{
				Logger: logger.Default.LogMode(logger.Warn),
			})
			if err != nil {
				log.Fatalf("[DATABASE FATAL] Gagal inisialisasi database SQLite fallback: %v", err)
			}
		} else {
			log.Println("[DATABASE SUCCESS] Terhubung ke PostgreSQL!")
		}
	} else {
		log.Println("[DATABASE INFO] Menggunakan SQLite (alhamid.db)...")
		db, err = gorm.Open(sqlite.Open("alhamid.db"), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Warn),
		})
		if err != nil {
			log.Fatalf("[DATABASE FATAL] Gagal inisialisasi database SQLite: %v", err)
		}
	}

	DB = db

	// Auto Migration
	log.Println("[DATABASE] Menjalankan Auto Migration skema tabel...")
	err = db.AutoMigrate(
		&models.User{},
		&models.Setting{},
		&models.Banner{},
		&models.Category{},
		&models.Article{},
		&models.EducationUnit{},
		&models.Caregiver{},
		&models.Facility{},
		&models.Achievement{},
		&models.Testimonial{},
		&models.Alumni{},
		&models.MediaPhoto{},
		&models.MediaVideo{},
		&models.ContactMessage{},
	)
	if err != nil {
		log.Fatalf("[DATABASE FATAL] Gagal migrasi skema tabel: %v", err)
	}

	// Seed Data
	SeedInitialData(db)

	return DB
}
