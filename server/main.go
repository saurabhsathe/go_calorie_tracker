package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/saurabhsathe/go-react-calorie-tracker/routes"
)

func main() {

	port := "8000"

	router := gin.New()
	router.Use(cors.Default())
	router.Use(gin.Logger())

	router.POST("/entry/create", routes.AddEntry)
	// router.GET("/entry/showall", routes.showallEntry)

}
