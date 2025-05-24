// database.go
package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type BancoDeDados struct {
	Host    string `json:"host"`
	Porta   int    `json:"porta"`
	Usuario string `json:"usuario"`
	Senha   string `json:"senha"`
	Banco   string `json:"banco"`
	URL     string `json:"url"`
}

// Struct para a configuração da aplicação
type ConfigApp struct {
	URLFila      string       `json:"url-fila"`
	NomeFila     string       `json:"nome-fila"`
	URLFrontend  string       `json:"url-frontend"`
	BancoDeDados BancoDeDados `json:"banco-de-dados"`
}

func GetConfig() (*ConfigApp, error) {

	var Config ConfigApp
	Config.BancoDeDados.Host = os.Getenv("DB_HOST")
	//passa eu profesor
	Config.BancoDeDados.Usuario = os.Getenv("DB_USER")
	Config.BancoDeDados.Senha = os.Getenv("DB_PASS")
	Config.BancoDeDados.Banco = os.Getenv("DB_NAME")
	Config.NomeFila = os.Getenv("NOME_FILA")
	Config.URLFila = os.Getenv("URL_FILA")
	Config.URLFrontend = os.Getenv("URL_FRONTEND")

	return &Config, nil
}

func ConectaBanco() *sql.DB {
	str_conn := GetStringConfig()
	//db, err := sql.Open("mysql", "username:password@tcp(<ip_do_banco>:3306)/banco")

	db, er := sql.Open("mysql", str_conn)

	if er != nil {
		return nil
	}
	return db
}

// func GetStringConfig() string {
// 	var Config ConfigApp
// 	Config.BancoDeDados.Host = os.Getenv("DB_HOST")

// 	Config.BancoDeDados.Usuario =  os.Getenv("DB_USER")
// 	Config.BancoDeDados.Senha = os.Getenv("DB_PASS")
// 	Config.BancoDeDados.Banco = os.Getenv("DB_NAME")
// 	Config.NomeFila = os.Getenv("NOME_FILA")
// 	Config.URLFila = os.Getenv("URL_FILA")
// 	Config.URLFrontend = os.Getenv("URL_FRONTEND")

// 	//config, err := GetConfig()
// 	//db := config.BancoDeDados
// 	//str_conn := MONTE AQUI
// 	str_conn := "root:123456@tcp(192.168.207.152:3306)/examed"
// 	// if err != nil {
// 	// 	fmt.Println("Erro ao ler o arquivo de configuração:", err)
// 	// 	return ""
// 	// }

// 	return str_conn
// }

func GetStringConfig() string {
	var Config ConfigApp

	// Config.BancoDeDados.Host = os.Getenv("DB_HOST")
	// Config.BancoDeDados.Usuario = os.Getenv("DB_USER")
	// Config.BancoDeDados.Senha = os.Getenv("DB_PASS")
	// Config.BancoDeDados.Banco = os.Getenv("DB_NAME")
	// Config.NomeFila = os.Getenv("NOME_FILA")
	// Config.URLFila = os.Getenv("URL_FILA")
	// Config.URLFrontend = os.Getenv("URL_FRONTEND")

	Config.BancoDeDados.Host = "192.168.1.31"
	Config.BancoDeDados.Usuario = "root"
	Config.BancoDeDados.Senha = "123456"
	Config.BancoDeDados.Banco = "examed"
	//Config.NomeFila = "exames-pendentes"

	// Exemplo para MySQL
	// Formato: usuario:senha@tcp(host:porta)/banco
	strConn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s",
		Config.BancoDeDados.Usuario,
		Config.BancoDeDados.Senha,
		Config.BancoDeDados.Host,
		Config.BancoDeDados.Banco,
	)

	return strConn
}
