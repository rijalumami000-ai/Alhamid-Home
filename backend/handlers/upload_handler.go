package handlers

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"alhamid-backend/config"

	"github.com/gofiber/fiber/v2"
)

type UploadHandler struct {
	Cfg *config.Config
}

func NewUploadHandler(cfg *config.Config) *UploadHandler {
	return &UploadHandler{Cfg: cfg}
}

func (h *UploadHandler) UploadFile(c *fiber.Ctx) error {
	file, err := c.FormFile("file")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "File tidak ditemukan dalam request"})
	}

	ext := strings.ToLower(filepath.Ext(file.Filename))
	allowed := map[string]bool{
		".jpg":  true,
		".jpeg": true,
		".png":  true,
		".webp": true,
		".gif":  true,
		".svg":  true,
		".pdf":  true,
	}

	if !allowed[ext] {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Format file tidak diizinkan. Gunakan JPG, PNG, WEBP, GIF, SVG, atau PDF."})
	}

	if file.Size > 10*1024*1024 { // 10MB limit
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Ukuran file melebihi batas 10MB"})
	}

	uploadDir := h.Cfg.UploadDir
	if err := os.MkdirAll(uploadDir, 0755); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Gagal menyiapkan direktori penyimpanan"})
	}

	fileName := fmt.Sprintf("%d_%s", time.Now().UnixNano(), filepath.Base(file.Filename))
	dst := filepath.Join(uploadDir, fileName)

	if err := c.SaveFile(file, dst); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Gagal menyimpan file"})
	}

	fileUrl := fmt.Sprintf("/uploads/%s", fileName)
	return c.JSON(fiber.Map{
		"message":  "Upload berhasil",
		"file_url": fileUrl,
		"filename": fileName,
	})
}

