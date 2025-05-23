// handlers.go
package main

import (
	"database/sql"
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/streadway/amqp"
)

var conn *sql.DB
var config_app *ConfigApp


func failOnError(err error, msg string) {
    if err != nil {
        log.Fatalf("%s: %s", msg, err)
    }
}

func SetDB(database *sql.DB) {
	conn = database
}

func SetConfig(config *ConfigApp){
    config = config_app
}

func PublishExame(exame string) {
    host_fila := config_app.URLFila

    conn_fila, err := amqp.Dial("amqp://guest:guest@"+host_fila+"/")

    failOnError(err, "Falha ao conectar ao RabbitMQ")
    defer conn_fila.Close()

    ch, err := conn_fila.Channel()
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