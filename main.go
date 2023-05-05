package main

import (
	"log"
)

func main() {
	if err := serve(); err != nil {
		log.Fatal(err)
	}

}

