package main

import (
	"context"
	"log"
	"os"
	workflow "temporal-go"

	"go.temporal.io/sdk/client"
)

func main() {
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create face", err)
	}
	defer c.Close()

	options := client.StartWorkflowOptions{
		ID:        "workflow-face-id",
		TaskQueue: "workflow-face",
	}

	we, err := c.ExecuteWorkflow(context.Background(), options, workflow.WorkflowOne, os.Args[1])
	if err != nil {
		log.Fatalln("Unable to execute workflow", err)
	}
	log.Println("Started workflow", "WorkflowID", we.GetID(), "RunID", we.GetRunID())

	var result string
	err = we.Get(context.Background(), &result)
	if err != nil {
		log.Fatalln("Unable to get workflow result", err)
	}
	log.Println("Workflow result:", result)
}
