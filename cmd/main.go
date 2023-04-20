package main

import (
	"fmt"
	"github.com/nurlan42/todo/cfg"
	"github.com/own/todo/internal/app"
	"log"
)

const cfgPath = "config.json"

func main() {
	c, err := cfg.NewConfig(cfgPath)
	if err != nil {
		log.Fatalf("NewConfig(): %v", err)
	}

	fmt.Printf("cfg = %#v\n", c)

	if err := app.Run(c); err != nil {
		log.Fatalf("Run(): %v", err)
	}

}
