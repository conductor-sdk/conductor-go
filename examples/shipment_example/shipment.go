package shipment_example

type Shipment struct {
	UserId  string
	OrderNo string
}

func NewShipment(userId string, orderNo string) *Shipment {
	return &Shipment{
		UserId:  userId,
		OrderNo: orderNo,
	}
}
