package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)



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
	Config, er := GetConfig()
	if er != nil {
		return ""
	}
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

	//esse comando fecha a conexão após o fim da execução dessa função
	defer db.Close()

	var nome_exame string

	query := "SELECT tipo_exame FROM Exames WHERE id_exame = ?"

	stmt, err := db.Prepare(query)

	if err != nil {
		return "", err
	}

	defer stmt.Close()

	row := stmt.QueryRow(id_exame)

	err = row.Scan(&nome_exame)

	if err != nil {
		return "", err
	}

	// query := "SELECT tipo_exame FROM Exames WHERE id_exame = ?"
	// err := db.QueryRow(query, id_exame).Scan(&nome_exame)
	// if err != nil {
	// 	return "", err
	// }

	return nome_exame, nil
}
