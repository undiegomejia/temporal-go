package orders

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

func CheckInventory(ctx context.Context, order OrderDetails) (bool, error) {
	resp, err := checkInventoryService("check-inventory", order.OrderID, strconv.Itoa(order.Quantity))
	return resp, err
}

func SendConfirmation(ctx context.Context, order OrderDetails) error {
	_, err := sendConfirmationService("send-confirmation", order.OrderID, order.CustomerID)
	if err != nil {
		return err
	}
	return nil
}

func PrepareShipping(ctx context.Context, order OrderDetails) error {
	err := prepareShippingService("prepare-shipping", order.OrderID, order.ShippingAddress)
	if err != nil {
		return err
	}
	return nil
}

func GenerateInvoice(ctx context.Context, order OrderDetails) (string, error) {
	invoiceId, err := generateInvoiceService("generate-invoice", order.OrderID)
	if err != nil {
		return invoiceId, err
	}
	return invoiceId, nil
}

func NotifyShipment(ctx context.Context, order OrderDetails, invoiceId string) error {
	err := notifyShipmentService("notify-shipment", order.OrderID, invoiceId)
	if err != nil {
		return err
	}
	return nil
}

func checkInventoryService(stem string, orderId string, quantity string) (bool, error) {
	base := "http://localhost:9999/" + stem + "?orderId=%s&quantity=%s"
	url := fmt.Sprintf(base, url.QueryEscape(orderId), url.QueryEscape(quantity))
	time.Sleep(10 * time.Second)
	resp, err := http.Get(url)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}
	response := string(body)
	status := resp.StatusCode
	if status >= 400 {
		return false, fmt.Errorf("HTTP ERROR %d: %s", status, response)
	}
	return true, nil
}

func prepareShippingService(stem string, orderId string, address string) error {
	base := "http://localhost:9999/" + stem + "?orderId=%s&shippingAddress=%s"
	url := fmt.Sprintf(base, url.QueryEscape(orderId), url.QueryEscape(address))
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

func sendConfirmationService(stem string, orderId string, customerId string) (bool, error) {
	base := "http://localhost:9999/" + stem + "?orderId=%s&customerId=%s"
	url := fmt.Sprintf(base, url.QueryEscape(orderId), url.QueryEscape(customerId))
	time.Sleep(10 * time.Second)
	resp, err := http.Get(url)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}
	response := string(body)
	status := resp.StatusCode
	if status >= 400 {
		return false, fmt.Errorf("HTTP ERROR %d: %s", status, response)
	}
	return true, nil
}

func generateInvoiceService(stem string, orderId string) (string, error) {
	base := "http://localhost:9999/" + stem + "?orderId=%s"
	url := fmt.Sprintf(base, url.QueryEscape(orderId))
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
