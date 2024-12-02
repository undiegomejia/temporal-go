package main

import (
	"context"
	"flag"
	"log"
	"os"
	dsl "temporal-go/yaml"

	"github.com/google/uuid"
	"go.temporal.io/sdk/client"
	"gopkg.in/yaml.v3"
)

func main() {
	var dslConfig string
	flag.StringVar(&dslConfig, "dslConfig", "../workflow1.yaml", "dslConfig specify the yaml file for the dsl workflow.")
	flag.Parse()

	data, err := os.ReadFile(dslConfig)
	if err != nil {
		log.Fatalln("failed to load dsl config file", err)
	}
	var dslWorkflow dsl.Workflow
	if err := yaml.Unmarshal(data, &dslWorkflow); err != nil {
		log.Fatalln("failed to unmarshal dsl config", err)
	}

	c, err := client.Dial(client.Options{
		HostPort: client.DefaultHostPort,
	})
	if err != nil {
		log.Fatalln("Unable to create order", err)
	}

	options := client.StartWorkflowOptions{
		ID:        "dsl_" + uuid.NewString(),
		TaskQueue: "dsl",
	}

	we, err := c.ExecuteWorkflow(context.Background(), options, dsl.DSLWorkflow, dslWorkflow)
	if err != nil {
		log.Fatalln("Unable to execute workflow", err)
	}
	log.Println("Started workflow", "WorkflowID", we.GetID(), "RunID", we.GetRunID())
}
