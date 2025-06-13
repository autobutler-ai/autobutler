package api

import (
	"time"

	"github.com/exokomodo/exoflow/autobutler/backend/pkg/llm"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func StartServer() error {
	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"POST", "GET", "PUT", "OPTIONS"}
	config.AllowHeaders = []string{"*"}
	config.ExposeHeaders = []string{"Content-Length"}
	config.AllowCredentials = true
	config.MaxAge = 12 * time.Hour
	router.Use(cors.New(config))

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})
	router.GET("/chat", func(c *gin.Context) {
		prompt := c.Query("prompt")
		response, err := llm.RemoteLLMRequest(prompt)
		if err != nil {
			c.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(200, response)
	})
	if err := router.Run(":8080"); err != nil {
		return err
	}

	return nil
}
