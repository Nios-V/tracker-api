package main

import (
	"github.com/Nios-V/tracker-api/internal/database"
	"github.com/Nios-V/tracker-api/internal/repository"
	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()

	repos := repository.NewRepositories(database.DB)

	r := gin.Default()
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})
	r.Run(":8080")
}
