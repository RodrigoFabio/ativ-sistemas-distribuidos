package main

import (
    "log"
    "github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
    if err != nil {
        log.Fatalf("%s: %s", msg, err)
    }
}

func main() {
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
