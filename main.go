package main

import (
	dashboardModels "algebra-isosofts-api/models/dashboards"
	registerComponentModels "algebra-isosofts-api/models/registers/components"
	"algebra-isosofts-api/routes"
	"fmt"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/lpernett/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	r := gin.Default()
	r.Use(cors.Default())
	r.RedirectTrailingSlash = false
	routes.APIRoutes(r.Group("/api"))

	go func() {
		var kpiModel dashboardModels.KPIModel

		ticker := time.NewTicker(1 * time.Hour)
		defer ticker.Stop()

		for range ticker.C {
			kpiModel.UpdateMonthsAll()
		}
	}()

	go func() {
		var actionModel registerComponentModels.ActionModel

		ticker := time.NewTicker(24 * time.Hour)
		defer ticker.Stop()

		for range ticker.C {
			actionModel.SendDailyNotifications()
		}
	}()

	fmt.Println(err)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	r.Run(":" + port)
}
