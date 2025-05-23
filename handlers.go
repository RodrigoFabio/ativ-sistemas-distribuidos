package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
    if err != nil {
        log.Fatalf("%s: %s", msg, err)
    }
}

func PublishExame(exame string) {
    conn, err := amqp.Dial("amqp://guest:guest@192.168.1.31:5672/")
    failOnError(err, "Falha ao conectar ao RabbitMQ")
    defer conn.Close()

    ch, err := conn.Channel()
    failOnError(err, "Falha ao abrir canal")
    defer ch.Close()


    body := "Olá, mundo!"
    err = ch.Publish(
        "",     // exchange
        "exames-pendentes", // chave de roteamento (routing key)
        false,  // mandatory
        false,  // immediate
        amqp.Publishing{
            ContentType: "application/json",
            Body:        []byte(body),
        })
    failOnError(err, "Falha ao publicar mensagem")
    log.Printf("Mensagem publicada: %s", body)
}

func GetAgendamentos(c *gin.Context){
     c.JSON(http.StatusOK, gin.H{
        "message": "Agendamentos",
    })
}

func CadastraExame(c *gin.Context)  {
    c.JSON(http.StatusOK, gin.H{
        "message": "Exame cadastrado com sucesso"}) 
}

func AgendaExame(c *gin.Context)  {
    c.JSON(http.StatusOK, gin.H{
        "message": "Exame agendado com sucesso"})
}