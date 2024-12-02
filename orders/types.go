package orders

type OrderDetails struct {
	OrderID         string `json:"orderId"`
	CustomerID      string `json:"customerId"`
	ProductDetails  string `json:"productDetails"`
	Quantity        int    `json:"quantity"`
	ShippingAddress string `json:"shippingAddress"`
	InvoiceId       string `json:"invoiceId"`
}
