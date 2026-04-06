package main

import (
	dashboardModels "algebra-isosofts-api/models/dashboards"
	registerModels "algebra-isosofts-api/models/registers"
	tableComponentModels "algebra-isosofts-api/models/tableComponents"
	"algebra-isosofts-api/routes"
	"os"

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

	r.GET("/reddli", func(c *gin.Context) {
		var dropDownListItemModel tableComponentModels.DropDownListItemModel
		dropDownListItemModel.DuplicateDefaults()
		c.IndentedJSON(201, gin.H{})
	})

	r.GET("/kpid", func(c *gin.Context) {
		var kpiModel dashboardModels.KPIModel
		kpiModel.DuplicateDefaults("qwqwqwqwqwqw")
		c.IndentedJSON(201, gin.H{})
	})

	r.GET("/rn", func(c *gin.Context) {
		var commonModel registerModels.CommonModel
		a, _ := commonModel.GetRegNo("32cP24T2HXM62zx5Zu2D4Jd10173QS", "cus")
		c.IndentedJSON(201, a)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	r.Run(":" + port)
}
