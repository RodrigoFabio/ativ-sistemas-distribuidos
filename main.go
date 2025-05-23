package main

import "fmt"

func main() {
	config, er := GetConfig()

	if er != nil {
		fmt.Print("ERRO")
	}

	db := ConectaBanco()
	SetConfig(config)
	SetDB(db)
	InitRoutes()
	PublishExame("exame")

}
