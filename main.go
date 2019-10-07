package main

import (
	"fmt"
	"log"

	"mp3duration"
)

func main() {
	duration, err := mp3duration.Calculate("demo - cbr.mp3")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Duration %v\n", duration)
}
