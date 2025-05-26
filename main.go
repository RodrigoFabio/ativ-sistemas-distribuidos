package main

import "fmt"

func main() {
	config, er := GetConfig(false)
	SetConfig(config)


	if er != nil {
		fmt.Print("ERRO")
	}

	db := ConectaBanco()
	
	SetDB(db)
	InitRoutes()
}
