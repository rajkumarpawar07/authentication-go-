package main

import (
	// "go.mongodb.org/mongo-driver/mongo/options"
	// routes "golang-jwt-project/models"
	"golang-jwt-project/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	// "go.mongodb.org/mongo-driver/mongo"
)

func main() {
	err := godotenv.Load(".env");
	if err != nil{
		log.Fatal("Error in loading env file")
	}

	port := os.Getenv("PORT")

	if port== ""{
		port = "8000"
	}

	router := gin.New()
	router.Use(gin.Logger())

	routes.AuthRoutes(router)
	routes.UserRoutes(router)

	router.GET("/api-1", func(c *gin.Context){
		c.JSON(200, gin.H{"success":"This is API-1 Response"})
	})

	router.GET("/api-2", func(c *gin.Context) {
		c.JSON(200, gin.H{"success":"This is API-2 Response"})
	})


	router.Run(":"+port)


}


