package main

import (
	dashboardModels "algebra-isosofts-api/models/dashboards"
	"algebra-isosofts-api/routes"
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

	// ---- ARKA PLAN ZAMANLAYICI (GOROUTINE) BAŞLANGICI ----
	// Bu bloğu go func() içine aldığımız için r.Run()'ı bloklamaz, arka planda akar.
	go func() {
		var kpiModel dashboardModels.KPIModel

		// Test için 1 saniye bırakıyorum, 4 saat için: 4 * time.Hour yaparsın
		ticker := time.NewTicker(1 * time.Hour)
		defer ticker.Stop()

		for range ticker.C {
			kpiModel.UpdateMonthsAll()
		}
	}()
	// ---- ARKA PLAN ZAMANLAYICI BİTİŞİ ----

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	// r.Run programı burada kilitler ve istekleri dinlemeye başlar.
	// Ama yukarıdaki kod 'go' ile başladığı için o arkada dönmeye devam eder.
	r.Run(":" + port)
}
