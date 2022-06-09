package shipment_example

type ShipmentState struct {
	PaymentCompleted bool
	EmailSent        bool
	Shipped          bool
	TrackingNumber   string
}

func NewShipmentState() *ShipmentState {
	return &ShipmentState{}
}

func (s *ShipmentState) IsPaymentCompleted() bool {
	return s.PaymentCompleted
}
