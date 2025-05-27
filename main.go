package main

import (
	"fmt"
	"log"
)

func main() {

	fmt.Println(":::::::::::::::::::::::::::::::::::::")
	log.Println("log com nível info")

	db := ConectaBanco()
	SetDB(db)
	InitRoutes()
}
