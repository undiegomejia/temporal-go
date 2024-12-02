package main

import (
	"log"
	orders "temporal-go/orders"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func main() {
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create Temporal client.", err)
	}
	defer c.Close()

	w := worker.New(c, "order", worker.Options{})

	w.RegisterWorkflow(orders.OrderProcessingWorkflow)
	w.RegisterActivity(orders.CheckInventory)
	w.RegisterActivity(orders.SendConfirmation)
	w.RegisterActivity(orders.PrepareShipping)
	w.RegisterActivity(orders.GenerateInvoice)
	w.RegisterActivity(orders.NotifyShipment)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Unable to start worker.", err)
	}
}
