package api

import "github.com/gin-gonic/gin"

func StartServer() error {
	router := gin.Default()
	// Define your routes here
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})
	// Start the server on port 8080
	if err := router.Run(":8080"); err != nil {
		return err
	}

	return nil
}
