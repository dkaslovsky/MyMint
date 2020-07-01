package main

import (
	"log"

	"github.com/dkaslovsky/MyMint/cmd"
)

func main() {
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
