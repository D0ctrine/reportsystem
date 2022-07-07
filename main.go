package main

import (
	"log"

	"report/rest"
)

func main() {
	log.Println("Main log....")
	log.Fatal(rest.RunAPI(":8000"))
}
