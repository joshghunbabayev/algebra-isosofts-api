package main

import (
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

	// r.GET("/reddli", func(c *gin.Context) {
	// 	var dropDownListItemModel tableComponentModels.DropDownListItemModel
	// 	dropDownListItemModel.DuplicateDefaults()
	// 	c.IndentedJSON(201, gin.H{})
	// })

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	r.Run(":" + port)
}
