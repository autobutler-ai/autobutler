package v1

import (
	"autobutler/internal/server/ui/components/email/email_list"
	"autobutler/internal/server/ui/types"

	"github.com/gin-gonic/gin"
)

func SetupEmailRoutes(apiV1Group *gin.RouterGroup) {
	listEmails(apiV1Group)
}

func listEmails(apiV1Group *gin.RouterGroup) {
	apiRoute(apiV1Group, "GET", "/email", func(c *gin.Context) {
		// Render the email list component with mock data
		if err := email_list.Component(types.MockEmails()).Render(c.Request.Context(), c.Writer); err != nil {
			c.Status(500)
			return
		}
		c.Status(200)
	})
}
