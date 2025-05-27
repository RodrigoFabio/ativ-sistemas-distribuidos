package main

import (
	"log"
)

func main() {

	log.Println(":::::::::::::::::::::::::::::::::::::")
	log.Println(":::::::::::::INICIANDO::::::::::::::")
	log.Println(":::::::::::::::::::::::::::::::::::::")

	db := ConectaBanco()
	SetDB(db)
	InitRoutes()
}
