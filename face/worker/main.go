package main

import (
	"log"
	workflow "temporal-go/face"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func main() {
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	w := worker.New(c, "workflow-face", worker.Options{})

	w.RegisterWorkflow(workflow.WorkflowOne)
	w.RegisterActivity(workflow.AddHair)
	w.RegisterActivity(workflow.AddVoice)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Unable to start worker", err)
	}
}
