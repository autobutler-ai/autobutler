package v1

import (
	"autobutler/pkg/util"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/nfnt/resize"
)

const (
	thumbnailWidth  = 400
	thumbnailHeight = 400
)

func SetupThumbnailRoutes(apiV1Group *gin.RouterGroup) {
	getThumbnailRoute(apiV1Group)
}

func getThumbnailRoute(apiV1Group *gin.RouterGroup) {
	apiRoute(apiV1Group, "GET", "/thumbnails/*filePath", func(c *gin.Context) {
		filePath := c.Param("filePath")
		filesDir := util.GetFilesDir()
		fullPath := filepath.Join(filesDir, filePath)

		// Check if file exists
		if _, err := os.Stat(fullPath); os.IsNotExist(err) {
			c.Status(http.StatusNotFound)
			return
		}

		// Open the original image
		file, err := os.Open(fullPath)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return
		}
		defer file.Close()

		// Decode the image
		img, format, err := image.Decode(file)
		if err != nil {
			// If we can't decode it, just serve the original file
			c.File(fullPath)
			return
		}

		// Generate thumbnail
		thumbnail := resize.Thumbnail(thumbnailWidth, thumbnailHeight, img, resize.Lanczos3)

		// Set appropriate content type
		ext := strings.ToLower(filepath.Ext(filePath))
		switch ext {
		case ".png":
			c.Header("Content-Type", "image/png")
			if err := png.Encode(c.Writer, thumbnail); err != nil {
				c.Status(http.StatusInternalServerError)
				return
			}
		case ".jpg", ".jpeg":
			c.Header("Content-Type", "image/jpeg")
			if err := jpeg.Encode(c.Writer, thumbnail, &jpeg.Options{Quality: 85}); err != nil {
				c.Status(http.StatusInternalServerError)
				return
			}
		default:
			// For other formats, try to encode as JPEG
			c.Header("Content-Type", fmt.Sprintf("image/%s", format))
			if err := jpeg.Encode(c.Writer, thumbnail, &jpeg.Options{Quality: 85}); err != nil {
				c.Status(http.StatusInternalServerError)
				return
			}
		}

		c.Status(http.StatusOK)
	})
}
