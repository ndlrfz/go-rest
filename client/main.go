package main

import (
	"client/client"
	"fmt"
	"log"
)

func main() {
	c := client.NewClient("http://localhost:8080")
	if err := c.Login("user", "password"); err != nil {
		log.Fatal(err)
	}

	value, err := c.Random()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Random value:", value)

	if err = c.SetSeed(20); err != nil {
		log.Fatal(err)
	}

	value, err = c.Random()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Random value with seed 20:", value)
}
