package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/deadlift-or-die/backend/config"
	"github.com/deadlift-or-die/backend/handler"
)

func main() {
	c, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	handler.RegisterHandlers()

	log.Printf("starting server on %v:%v", c.Host, c.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%v:%v", c.Host, c.Port), nil))
}
