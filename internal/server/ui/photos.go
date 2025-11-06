package ui

import (
	"autobutler/internal/server/ui/types"
	"autobutler/internal/server/ui/views"

	"github.com/gin-gonic/gin"
)

func SetupPhotoRoutes(router *gin.Engine) {
	setupPhotoView(router)
}

func setupPhotoView(router *gin.Engine) {
	uiRoute(router, "/photos", func(c *gin.Context) {
		if err := views.Photos(types.NewPageState()).Render(c.Request.Context(), c.Writer); err != nil {
			c.Status(400)
			return
		}
		c.Status(200)
	})
	uiRoute(router, "/photos/*rootDir", func(c *gin.Context) {
		rootDir := c.Param("rootDir")
		if err := views.Photos(types.NewPageState().WithRootDir(rootDir)).Render(c.Request.Context(), c.Writer); err != nil {
			c.Status(400)
			return
		}
		c.Status(200)
	})
}
