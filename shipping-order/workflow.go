package orders

import (
	"fmt"
	"log"
	"time"

	"go.temporal.io/sdk/workflow"
)

func OrderProcessingWorkflow(ctx workflow.Context, order OrderDetails) (OrderDetails, error) {
	ao := workflow.ActivityOptions{
		StartToCloseTimeout: 3 * time.Hour,
	}

	ctx = workflow.WithActivityOptions(ctx, ao)

	// Check Inventory (Sequencial)
	var inventoryAvailable bool
	err := workflow.ExecuteActivity(ctx, CheckInventory, order).Get(ctx, &inventoryAvailable)
	if err != nil {
		return order, err
	}
	if !inventoryAvailable {
		fmt.Printf("Order %s cannot be fulfilled due to lack of inventory.\n", order.OrderID)
		return order, fmt.Errorf("order %s cannot be fulfilled due to lack of inventory", order.OrderID)
	}
	// Send Confirmation (Sequencial)
	err = workflow.ExecuteActivity(ctx, SendConfirmation, order).Get(ctx, nil)
	if err != nil {
		return order, err
	}
	// Parallel Activities: PrepareShipment and GenerateInvoice
	prepareShipmentFuture := workflow.ExecuteActivity(ctx, PrepareShipping, order)
	generateInvoiceFuture := workflow.ExecuteActivity(ctx, GenerateInvoice, order)
	// Wait prepare shipping to finish
	err = prepareShipmentFuture.Get(ctx, nil)
	if err != nil {
		return order, err
	}
	// Wait generate invoice to finish
	err = generateInvoiceFuture.Get(ctx, &order.InvoiceId)
	if err != nil {
		return order, err
	}
	// Notify Shipment (Sequential)
	err = workflow.ExecuteActivity(ctx, NotifyShipment, order).Get(ctx, nil)
	if err != nil {
		return order, err
	}
	log.Printf("Order %s proccesed succesfullt.\n", order.OrderID)
	return order, nil
}
