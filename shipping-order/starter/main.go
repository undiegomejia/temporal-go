package main

import (
	"context"
	"log"
	orders "temporal-go/shipping-order"

	"github.com/google/uuid"
	"go.temporal.io/sdk/client"
)

func main() {
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create order", err)
	}
	defer c.Close()

	options := client.StartWorkflowOptions{
		ID:        "order_" + uuid.NewString(),
		TaskQueue: "order",
	}

	we, err := c.ExecuteWorkflow(context.Background(), options, orders.OrderProcessingWorkflow, orders.OrderDetails{
		OrderID:         "78787",
		CustomerID:      "cust-89898",
		ProductDetails:  "Laptop",
		Quantity:        1,
		ShippingAddress: "123 Print St. San Jose",
	})
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
