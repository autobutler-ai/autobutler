package ui

import (
	"autobutler/internal/server/ui/types"
	"autobutler/internal/server/ui/views"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func SetupCalendarRoutes(router *gin.Engine) {
	uiRoute(router, "/calendar", func(c *gin.Context) {
		yearStr := c.Query("year")
		monthStr := c.Query("month")
		
		var targetTime *time.Time
		if yearStr != "" && monthStr != "" {
			year, err := strconv.Atoi(yearStr)
			if err == nil {
				month, err := strconv.Atoi(monthStr)
				if err == nil && month >= 1 && month <= 12 {
					t := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)
					targetTime = &t
				}
			}
		}
		
		if err := views.CalendarWithTime(types.NewPageState(), targetTime).Render(c.Request.Context(), c.Writer); err != nil {
			c.Status(400)
			return
		}
		c.Status(200)
	})
}
