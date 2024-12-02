package main

import (
	"log"
	dsl "temporal-go/yaml"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func main() {
	c, err := client.Dial(client.Options{
		HostPort: client.DefaultHostPort,
	})
	if err != nil {
		log.Fatalln("Unable to create Temporal client.", err)
	}
	defer c.Close()

	w := worker.New(c, "dsl", worker.Options{})

	w.RegisterWorkflow(dsl.DSLWorkflow)
	w.RegisterActivity(&dsl.Activities{})

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Unable to start worker.", err)
	}
}
