package dsl

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"go.temporal.io/api/workflowservice/v1"
	"go.temporal.io/sdk/client"
)

type Activities struct{}

func (a *Activities) CheckInventory(ctx context.Context, input []string) (string, error) {
	r, err := checkInventoryService("check-inventory", input[0], input[1])
	return r, err
}

func (a *Activities) SendConfirmation(ctx context.Context, input []string) error {
	_, err := sendConfirmationService("send-confirmation", input[0], input[1])
	if err != nil {
		return err
	}
	return nil
}

func (a *Activities) PrepareShipping(ctx context.Context, input []string) error {
	err := prepareShippingService("prepare-shipping", input[0], input[1])
	if err != nil {
		return err
	}
	return nil
}

func (a *Activities) GenerateInvoice(ctx context.Context, input []string) (string, error) {
	invoiceId, err := generateInvoiceService("generate-invoice", input[0])
	if err != nil {
		return invoiceId, err
	}
	return invoiceId, nil
}

func (a *Activities) NotifyShipment(ctx context.Context, input []string) error {
	err := notifyShipmentService("notify-shipment", input[0], input[1])
	if err != nil {
		return err
	}
	return nil
}

func (a *Activities) HumanApprovement(ctx context.Context, input []string) (string, error) {
	var err error
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()
	var approved = humanAprroved("Do you want to approve this order?")
	if approved == "false" {
		fmt.Println("Order not approved :(")
		executions, err := c.ListWorkflow(context.Background(), &workflowservice.ListWorkflowExecutionsRequest{Query: ""})
		if err != nil {
			log.Fatalln("Unable to show workflow executions", err)
		}
		var workflowId string
		if executions.Size() > 0 {
			workflowId = executions.GetExecutions()[0].GetExecution().GetWorkflowId()
		}
		err = c.CancelWorkflow(context.Background(), workflowId, "")
		if err != nil {
			log.Fatalln("Unable to cancel workflow execution", err)
		}
	}
	return approved, err
}

func checkInventoryService(stem string, orderId string, quantity string) (string, error) {
	base := "http://localhost:9999/" + stem + "?orderId=%s&quantity=%s"
	url := fmt.Sprintf(base, url.QueryEscape(orderId), url.QueryEscape(quantity))
	time.Sleep(10 * time.Second)
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	response := string(body)
	status := resp.StatusCode
	if status >= 400 {
		return "", fmt.Errorf("HTTP ERROR %d: %s", status, response)
	}
	return response, nil
}

func prepareShippingService(stem string, orderId string, shippingAddress string) error {
	base := "http://localhost:9999/" + stem + "?orderId=%s&shippingAddress=%s"
	url := fmt.Sprintf(base, url.QueryEscape(orderId), url.QueryEscape(shippingAddress))
	time.Sleep(10 * time.Second)
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	response := string(body)
	status := resp.StatusCode
	if status >= 400 {
		return fmt.Errorf("HTTP ERROR %d: %s", status, response)
	}
	return nil
}

func sendConfirmationService(stem string, orderId string, customerId string) (string, error) {
	base := "http://localhost:9999/" + stem + "?orderId=%s&customerId=%s"
	url := fmt.Sprintf(base, url.QueryEscape(orderId), url.QueryEscape(customerId))
	time.Sleep(10 * time.Second)
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	response := string(body)
	status := resp.StatusCode
	if status >= 400 {
		return "", fmt.Errorf("HTTP ERROR %d: %s", status, response)
	}
	return response, nil
}

func generateInvoiceService(stem string, orderId string) (string, error) {
	base := "http://localhost:9999/" + stem + "?orderId=%s"
	url := fmt.Sprintf(base, url.QueryEscape(orderId))
	time.Sleep(40 * time.Second)
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	response := string(body)
	status := resp.StatusCode
	if status >= 400 {
		return response, fmt.Errorf("HTTP ERROR %d: %s", status, response)
	}
	return response, nil
}

func notifyShipmentService(stem string, orderId string, invoiceId string) error {
	base := "http://localhost:9999/" + stem + "?orderId=%s&invoiceId=%s"
	url := fmt.Sprintf(base, url.QueryEscape(orderId), url.QueryEscape(invoiceId))
	time.Sleep(10 * time.Second)
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	response := string(body)
	status := resp.StatusCode
	if status >= 400 {
		return fmt.Errorf("HTTP ERROR %d: %s", status, response)
	}
	return nil
}

func humanAprroved(s string) string {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("%s [y/n]: ", s)
		response, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		response = strings.ToLower(strings.TrimSpace(response))
		if response == "y" || response == "yes" {
			return "true"
		} else if response == "n" || response == "no" {
			return "false"
		}
	}
}
