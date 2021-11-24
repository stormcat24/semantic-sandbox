package main

import (
	"log"

	"github.com/stormcat24/semantic-sandbox/pkg/logic"
)

func main() {
	res := logic.Hello("tori")
	log.Println(res)
}
