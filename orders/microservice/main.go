package main

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

func checkInventoryHandler(w http.ResponseWriter, r *http.Request) {
	_, order_ok := r.URL.Query()["orderId"]
	_, quantity_ok := r.URL.Query()["quantity"]
	if order_ok && quantity_ok {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, true)
	} else {
		http.Error(w, "Missing required queries.", http.StatusBadRequest)
	}
}

func sendConfirmationHandler(w http.ResponseWriter, r *http.Request) {
	_, order_ok := r.URL.Query()["orderId"]
	_, customer_ok := r.URL.Query()["customerId"]
	if order_ok && customer_ok {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, true)
	} else {
		http.Error(w, "Missing required queries.", http.StatusBadRequest)
	}
}

func prepareShippingHandler(w http.ResponseWriter, r *http.Request) {
	_, order_ok := r.URL.Query()["orderId"]
	_, address_ok := r.URL.Query()["shippingAddress"]
	if order_ok && address_ok {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, true)
	} else {
		http.Error(w, "Missing required queries.", http.StatusBadRequest)
	}
}

func generateInvoiceHandler(w http.ResponseWriter, r *http.Request) {
	_, order_ok := r.URL.Query()["orderId"]
	if order_ok {
		invoiceId := "inv_" + uuid.NewString()
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, invoiceId)
	} else {
		http.Error(w, "Missing required queries.", http.StatusBadRequest)
	}
}

func notifyShipmentHandler(w http.ResponseWriter, r *http.Request) {
	_, order_ok := r.URL.Query()["orderId"]
	_, invoice_ok := r.URL.Query()["invoiceId"]
	if order_ok && invoice_ok {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, true)
	} else {
		http.Error(w, "Missing required queries.", http.StatusBadRequest)
	}
}

func main() {
	http.HandleFunc("/check-inventory", checkInventoryHandler)
	http.HandleFunc("/send-confirmation", sendConfirmationHandler)
	http.HandleFunc("/prepare-shipping", prepareShippingHandler)
	http.HandleFunc("/generate-invoice", generateInvoiceHandler)
	http.HandleFunc("/notify-shipment", notifyShipmentHandler)
	http.ListenAndServe(":9999", nil)
}
