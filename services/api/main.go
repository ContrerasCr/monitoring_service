package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Usuario struct {
	ID     int    `json:"id"`
	Nombre string `json:"nombre"`
	Email  string `json:"email"`
}

var usuarios = []Usuario{
	{ID: 1, Nombre: "Juan", Email: "juan@example.com"},
	{ID: 2, Nombre: "María", Email: "maria@example.com"},
}

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		// Return JSON response
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/usuarios", func(c *gin.Context) {
		c.JSON(http.StatusOK, usuarios)
	})

	r.GET("/usuarios/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(http.StatusOK, gin.H{"id": id})
	})

	if err := r.Run(); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
