package main

import (
	"log"
	"os"
)

func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func GetConfig() (*ConfigApp, error) {
	var Config ConfigApp
	tipo_ambiente := os.Getenv("TIPO_AMBIENTE")

	if tipo_ambiente == "PROD" {
		Config.BancoDeDados.Host = os.Getenv("DB_HOST")
		//passa eu profesor
		Config.BancoDeDados.Usuario = os.Getenv("DB_USER")
		Config.BancoDeDados.Senha = os.Getenv("DB_PASS")
		Config.BancoDeDados.Banco = os.Getenv("DB_NAME")
		Config.NomeFila = os.Getenv("NOME_FILA")
		Config.URLFila = os.Getenv("URL_FILA")

		log.Println("DB_HOST::", Config.BancoDeDados.Host)
		log.Println("DB_USER::", Config.BancoDeDados.Usuario)
		log.Println("DB_PASS::", Config.BancoDeDados.Senha)
		log.Println("DB_NAME::", Config.BancoDeDados.Banco)
		log.Println("NOME_FILA::", Config.NomeFila)
		log.Println("URL_FILA::", Config.URLFila)

	} else {
		Config.BancoDeDados.Host = "192.168.207.163"
		Config.BancoDeDados.Usuario = "root"
		Config.BancoDeDados.Senha = "123456"
		Config.BancoDeDados.Banco = "examed"
		Config.NomeFila = "exames-pendentes"
		Config.URLFila = "192.168.207.153"
	}

	return &Config, nil
}
