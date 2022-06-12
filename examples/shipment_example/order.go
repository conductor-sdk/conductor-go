package shipment_example

import "github.com/conductor-sdk/conductor-go/examples/shipment_example/shipment_method_example"

type Order struct {
	OrderNumber    string
	Sku            string
	Quantity       int
	UnitPrice      float64
	ZipCode        string
	CountryCode    string
	ShippingMethod shipment_method_example.ShipmentMethod
}

func NewOrder(orderNumber string, sku string, quantity int, unitPrice float64) *Order {
	return &Order{
		OrderNumber: orderNumber,
		Sku:         sku,
		Quantity:    quantity,
		UnitPrice:   unitPrice,
	}
}
