package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Arghya-Banerjee/urlShortener/api/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Redis Address:", os.Getenv("DB_ADDRESS"))
	fmt.Println("App Port:", os.Getenv("APP_PORT"))
	fmt.Println("Domain:", os.Getenv("DOMAIN"))

	router := gin.Default()

	setupRouters(router)

	// port := os.Getenv("APP_PORT")
	// if port == "" {
	// 	port = "8080"
	// }

	// log.Fatal(router.Run(":" + port))

	log.Fatal(router.Run(":8000"))
}

func setupRouters(router *gin.Engine) {
	router.POST("/api/v1/", routes.ShortenURL)
	router.GET("/api/v1/shortID", routes.GetByShortID)
	router.DELETE("/api/v1/:shortID", routes.DeleteURL)
	router.PUT("/api/v1/:shortID", routes.EditURL)
	router.POST("api/v1/addTag", routes.AddTag)
	router.GET("/test/", routes.TestP)
}
