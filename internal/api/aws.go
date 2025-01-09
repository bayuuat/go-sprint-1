package api

import (
	"context"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/bayuuat/go-sprint-1/dto"
	"github.com/bayuuat/go-sprint-1/internal/middleware"
	"github.com/bayuuat/go-sprint-1/internal/utils"
	"github.com/gofiber/fiber/v2"
)

type awsApi struct{}

func NewAws(app *fiber.App) {

	da := awsApi{}

	user := app.Group("/v1/file")

	user.Use(middleware.JWTProtected)
	user.Post("/", da.UploadFile)
}

var wg sync.WaitGroup

func (da awsApi) UploadFile(ctx *fiber.Ctx) error {
	_, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	file, err := ctx.FormFile("file")
	if err != nil {
		return ctx.Status(400).JSON(dto.ErrorResponse{Message: err.Error()})
	}

	if file.Size > 100*1024 {
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Message: "File size exceeds 100KiB"})
	}

	// Check file type
	fileHeader := make([]byte, 512)
	openedFile, err := file.Open()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{Message: err.Error()})
	}
	defer openedFile.Close()

	_, err = openedFile.Read(fileHeader)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{Message: err.Error()})
	}

	fileType := http.DetectContentType(fileHeader)
	if fileType != "image/jpeg" && fileType != "image/jpg" && fileType != "image/png" {
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Message: "Invalid file type. Only JPEG, JPG, and PNG are allowed"})
	}

	bucketName := os.Getenv("AWS_BUCKET_NAME")

	filePath := file
	bucket := bucketName
	prefix := "images"

	res, err := utils.UploadFileToS3(filePath, bucket, prefix)
	if err != nil {
		log.Println("Error uploading file:", err)
		return ctx.Status(500).JSON(fiber.Map{"message": err.Error()})
	}

	return ctx.Status(200).JSON(fiber.Map{"uri": res})
}
