// routes.go
package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func InitRoutes() {
	router := gin.Default()

	router.Use(MiddlewareCors())
	fmt.Print("chegou aqui")
	router.GET("/api/agendamentos", GetAgendamentos)

	router.GET("/api/recupera-exames", RecuperaExames)

	router.POST("/api/cadastra-exame", CadastraExame)

	router.POST("/api/agenda-exame", AgendaExame)

	router.Run(":8080")
}

func MiddlewareCors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204) // Responde no preflight com status No Content
			return
		}

		c.Next()
	}
}
