//database.go
package main

import (
	"fmt"
	"os"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type BancoDeDados struct {
	Host     string `json:"host"`
	Porta    int    `json:"porta"`
	Usuario  string `json:"usuario"`
	Senha    string `json:"senha"`
	Banco    string `json:"banco"`
	URL      string `json:"url"`
}

// Struct para a configuração da aplicação
type ConfigApp struct {
	URLFila       string      `json:"url-fila"`
	NomeFila      string      `json:"nome-fila"`
	URLFrontend   string      `json:"url-frontend"`
	BancoDeDados  BancoDeDados `json:"banco-de-dados"`
}

func GetConfig() (*ConfigApp, error) {

	var Config ConfigApp
	Config.BancoDeDados.Host = os.Getenv("DB_HOST")
	Config.BancoDeDados.Usuario =  os.Getenv("DB_USER")
	Config.BancoDeDados.Senha = os.Getenv("DB_PASS")
	Config.BancoDeDados.Banco = os.Getenv("DB_NAME")
	Config.NomeFila = os.Getenv("NOME_FILA")
	Config.URLFila = os.Getenv("URL_FILA")
	Config.URLFrontend = os.Getenv("URL_FRONTEND")
	
	return &Config, nil	
}

func ConectaBanco() *sql.DB{
	// Lê o arquivo de configuração
	
	str_conn := GetStringConfig()

	db, er := sql.Open("mysql", str_conn)

	if er != nil{
		return nil
	}

	return db
	// Conecta ao banco de dados usando as informações do arquivo de configuração

	//db, err := sql.Open("mysql", "username:password@tcp(127.0.0.1:3306)/test")
}

func GetStringConfig() string {
	config, err := GetConfig()
	db := config.BancoDeDados
	str_conn := db.Usuario + ":" + db.Senha + "@" + "tcp(" + db.Host + ":" + string(db.Porta) + ")" + "/" + db.Banco
	
	if err != nil {
		fmt.Println("Erro ao ler o arquivo de configuração:", err)
		return ""
	}

	return str_conn
}