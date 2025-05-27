package main

import (
	"fmt"
	"log"
)

func main() {

	fmt.Println("mensagem de teste")
	log.Println("log com nível info")
	SetConfig()
	db := ConectaBanco()
	SetDB(db)
	InitRoutes()
}
