package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/streadway/amqp"
)

var conn *sql.DB

func SetDB(database *sql.DB) {
	conn = database
}

func PublishExame(agendamento Agendamentos) {
	config_app, errr := GetConfig()

	if errr != nil {
		fmt.Printf(" Erro ao carregar config: %v\n", errr)
	}

	host_fila := config_app.URLFila
	conn_fila, err := amqp.Dial("amqp://guest:guest@" + host_fila + ":5672/")
	//conn_fila, err := amqp.Dial("amqp://guest:guest@192.168.207.165:5672/")

	FailOnError(err, "Falha ao conectar ao RabbitMQ")

	//mapping de agendamento para o tipo de mensagem que será enviada
	var exame AgendamentoEmail
	exame.DataHora = agendamento.DataHora
	exame.EmailPaciente = agendamento.EmailPaciente
	exame.Instrucoes = agendamento.Instrucoes
	exame.NomeExame, err = GetNomeExame(agendamento.IdExame)

	FailOnError(err, "Falha ao obter nome do exame")

	exame.Paciente = agendamento.Paciente

	body, err := json.Marshal(exame)

	if err != nil {
		FailOnError(err, "Falha ao serializar exame")
	}

	FailOnError(err, "Falha ao conectar ao RabbitMQ")
	defer conn_fila.Close()

	channel, err := conn_fila.Channel()
	FailOnError(err, "Falha ao abrir canal")
	defer channel.Close()

	err = channel.Publish(
		"",                 // exchange
		"exames-pendentes", // chave de roteamento (routing key)
		false,              // mandatory
		false,              // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        []byte(body),
		})

	FailOnError(err, "Falha ao publicar mensagem")
	log.Printf("Mensagem publicada: %s", body)
}

func GetAgendamentos(c *gin.Context) {

	rows, er := conn.Query(`select 	a.nome_paciente, 
									a.email_paciente, 
									a.data, 
									a.id_exame, 
									e.tipo_exame, 
									a.instrucao, 
									a.cpf, 
									a.cartao_sus
									from Agendamentos a, Exames e
									where a.id_exame = e.id_exame 
									limit 50;`)

	if er != nil {
		log.Println("Erro ao buscar agendamentos")
	}

	defer rows.Close()

	var recuperados []RecuperaAgendamento

	for rows.Next() {

		var recuperaAgendamento RecuperaAgendamento

		err := rows.Scan(&recuperaAgendamento.NomePaciente,
			&recuperaAgendamento.EmailPaciente,
			&recuperaAgendamento.Data,
			&recuperaAgendamento.IdExame,
			&recuperaAgendamento.TipoExame,
			&recuperaAgendamento.Instrucoes,
			&recuperaAgendamento.Cpf,
			&recuperaAgendamento.CartaoSus)

		if err != nil {
			log.Print(err)
		}

		recuperados = append(recuperados, recuperaAgendamento)
	}

	c.JSON(200, gin.H{"agendamentos": recuperados})
}

func RecuperaExames(c *gin.Context) {
	rows, err := conn.Query("SELECT id_exame, tipo_exame, instrucoes FROM Exames")

	if err != nil {
		log.Println("Erro ao executar consulta:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar exames"})
		return
	}

	defer rows.Close()

	var exames []Exames

	// Itera sobre todas as linhas
	for rows.Next() {
		var exame Exames

		err := rows.Scan(&exame.IdExame, &exame.TipoExame, &exame.Instrucoes)
		if err != nil {
			log.Println("Erro ao ler linha:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao ler exames"})
			return
		}

		exames = append(exames, exame)
	}

	// Verifica se houve erro durante a iteração
	if err := rows.Err(); err != nil {
		log.Println("Erro ao iterar sobre resultados:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao processar exames"})
		return
	}

	// Retorna a lista de exames em JSON
	c.JSON(http.StatusOK, gin.H{"exames": exames})
}

func CadastraExame(c *gin.Context) {
	// Obtém os dados do corpo da requisição
	var exame Exames
	if err := c.ShouldBindJSON(&exame); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	// Insere o exame no banco de dados
	_, err := conn.Exec("INSERT INTO Exames (tipo_exame, instrucoes) VALUES (?, ?)", exame.TipoExame, exame.Instrucoes)
	if err != nil {
		log.Println("Erro ao inserir exame:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao cadastrar exame"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Exame cadastrado com sucesso"})
}

func AgendaExame(c *gin.Context) {
	var agendamento Agendamentos

	if err := c.ShouldBindJSON(&agendamento); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	// Insere o agendamento no banco de dados
	_, err := conn.Exec("INSERT INTO Agendamentos (data, id_exame, instrucao, nome_paciente, email_paciente, cpf, cartao_sus) VALUES (?, ?, ?, ?, ?, ?, ?)",
		agendamento.DataHora, agendamento.IdExame, agendamento.Instrucoes, agendamento.Paciente, agendamento.EmailPaciente, agendamento.CpfPaciente, agendamento.CartaoSus)
	if err != nil {
		log.Println("Erro ao inserir agendamento:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao cadastrar agendamento"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Agendamento cadastrado com sucesso"})

	PublishExame(agendamento)
}
