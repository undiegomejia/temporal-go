package orders

type OrderDetails struct {
	OrderID         string `json:"orderId"`
	CustomerID      string `json:"customerId"`
	ProductDetails  string `json:"productDetails"`
	Quantity        int    `json:"quantity"`
	ShippingAddress string `json:"shippingAddress"`
	InvoiceId       string `json:"invoiceId"`
}

// type (
// 	Workflow struct {
// 		Variables map[string]string
// 		Root      Statement
// 	}
// 	Statement struct {
// 		Activity *ActivityInvocation
// 		Sequence *Sequence
// 		Parallel *Parallel
// 	}
// 	Sequence struct {
// 		Elements []*Statement
// 	}
// 	Parallel struct {
// 		Branches []*Statement
// 	}
// 	ActivityInvocation struct {
// 		Name      string
// 		Arguments []string
// 		Result    string
// 	}
// 	executable interface {
// 		execute(ctx workflow.Context, bindings map[string]string)
// 	}
// )
