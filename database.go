package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func GetConfig(tipo_ambiente bool) (*ConfigApp, error) {
	var Config ConfigApp

	if tipo_ambiente {
		Config.BancoDeDados.Host = os.Getenv("DB_HOST")
		//passa eu profesor
		Config.BancoDeDados.Usuario = os.Getenv("DB_USER")
		Config.BancoDeDados.Senha = os.Getenv("DB_PASS")
		Config.BancoDeDados.Banco = os.Getenv("DB_NAME")
		Config.NomeFila = os.Getenv("NOME_FILA")
		Config.URLFila = os.Getenv("URL_FILA")

	} else {
		Config.BancoDeDados.Host = "192.168.207.163"
		Config.BancoDeDados.Usuario = "root"
		Config.BancoDeDados.Senha = "123456"
		Config.BancoDeDados.Banco = "examed"
		Config.NomeFila = "exames-pendentes"
		Config.URLFila = "192.168.207.153:5672"
	}

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

func GetStringConfig() string {
	Config, er := GetConfig(true)
	//var Config ConfigApp

	if er != nil {
		return ""
	}

	// CASO NÃO ESTEJA USANDO O DOCKER, DESCOMENTE A LINHA ABAIXO
	// Config.BancoDeDados.Host = "192.168.1.31"
	// Config.BancoDeDados.Usuario = "root"
	// Config.BancoDeDados.Senha = "123456"
	// Config.BancoDeDados.Banco = "examed"
	// Config.NomeFila = "exames-pendentes"

	// Exemplo para MySQL
	// Formato: usuario:senha@tcp(host:porta)/banco
	strConn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s",
		Config.BancoDeDados.Usuario,
		Config.BancoDeDados.Senha,
		Config.BancoDeDados.Host,
		Config.BancoDeDados.Banco,
	)
	fmt.Print("::::::::", strConn, ":::::::::")
	return strConn
}

func GetNomeExame(id_exame int) (string, error) {
	db := ConectaBanco()
	defer db.Close()

	var nome_exame string
	query := "SELECT tipo_exame FROM Exames WHERE id_exame = ?"
	err := db.QueryRow(query, id_exame).Scan(&nome_exame)
	if err != nil {
		return "", err
	}

	return nome_exame, nil
}
