// handlers.go
package main

import (
	"database/sql"
	"log"
   // "fmt"
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

// func GetAgendamentos(c *gin.Context){
//     conn.Query
//      c.JSON(http.StatusOK, gin.H{
//         "message": "Agendamentos",
//     })
// }

// func GetAgendamentos(c *gin.Context) {
//     // Executa a consulta
//     rows, err := conn.Query("SELECT * FROM Exames")
//     if err != nil {
//         log.Println("Erro ao executar consulta:", err)
//         c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar exames"})
//         return
//     }
//     defer rows.Close()

//     // Lê a primeira linha
//     if rows.Next() {
//         var id int
//         var exame string
//         var instrucoes string
//         err = rows.Scan(&id, &exame, &instrucoes)
//         if err != nil {
//             log.Println("Erro ao ler linha:", err)
//             c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao ler exame"})
//             return
//         }

//         // Retorna os dados como string
//         resultado := fmt.Sprintf("ID: %d, exame: %s, isntrucoes: %s", id, exame, instrucoes)
//         c.JSON(http.StatusOK, gin.H{"exame": resultado})
//         return
//     }

//     // Se não houver dados
//     c.JSON(http.StatusOK, gin.H{"message": "Nenhum exame encontrado"})
// }


func GetAgendamentos(c *gin.Context) {
    // Executa a consulta
    rows, err := conn.Query("SELECT * FROM Exames")
    if err != nil {
        log.Println("Erro ao executar consulta:", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar exames"})
        return
    }
    defer rows.Close()

    // Slice para armazenar os exames
    var exames []map[string]interface{}

    // Itera sobre todas as linhas
    for rows.Next() {
        var id int
        var exame string
        var instrucoes string

        err := rows.Scan(&id, &exame, &instrucoes)
        if err != nil {
            log.Println("Erro ao ler linha:", err)
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao ler exames"})
            return
        }

        exameData := map[string]interface{}{
            "id":         id,
            "exame":      exame,
            "instrucoes": instrucoes,
        }

        exames = append(exames, exameData)
    }

    // Verifica se houve algum erro na iteração
    if err := rows.Err(); err != nil {
        log.Println("Erro ao iterar sobre resultados:", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao processar exames"})
        return
    }

    // Retorna a lista de exames em JSON
    c.JSON(http.StatusOK, gin.H{"exames": exames})
}



func CadastraExame(c *gin.Context)  {
    c.JSON(http.StatusOK, gin.H{
        "message": "Exame cadastrado com sucesso"}) 
}

func AgendaExame(c *gin.Context)  {
    c.JSON(http.StatusOK, gin.H{
        "message": "Exame agendado com sucesso"})
}