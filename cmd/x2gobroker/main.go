package main

import (
	"fmt"
	"github.com/basterrus/x2gobroker/internal/config"
	"log"
)

func main() {
	cfg, err := config.InitConfig()
	if err != nil {
		log.Fatalf("error load config file: %s", err)
	}

	fmt.Print(cfg)

}
