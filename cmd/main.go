package main

import (
	"fmt"
	"log"

	"github.com/faiyaz032/chaos-rabbit/internal/config"
)

func main() {
	cfg, err := config.Load("http.yaml")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(cfg.HTTP.Listen)
	fmt.Println(cfg.HTTP.Target)
}
