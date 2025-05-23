package main

import ("github.com/gin-gonic/gin")


func InitRoutes() {
	router := gin.Default()

	router.GET("/api/agendamentos",  GetAgendamentos)

	router.POST("/api/cadastra-exame",  CadastraExame)

	router.POST("/api/agenda-exame", AgendaExame)

	router.Run(":8080")
}