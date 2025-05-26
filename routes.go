// routes.go
package main

import (
	"github.com/gin-gonic/gin"
)

func InitRoutes() {
	router := gin.Default()

	router.Use(MiddlewareCors())

	router.GET("/api/agendamentos", GetAgendamentos)

	router.GET("/api/recupera-exames", RecuperaExames)

	router.POST("/api/cadastra-exame", CadastraExame)

	router.POST("/api/agenda-exame", AgendaExame)

	router.Run(":8080")
}

func MiddlewareCors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Origin", "*")
		c.Header("Access-Control-Methods", "*")
		c.Header("Access-Control-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Credentials", "true")

		c.Next()
	}
}
