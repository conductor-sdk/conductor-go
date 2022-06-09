package shipment_method_example

type ShipmentMethod string

const (
	Ground     ShipmentMethod = "GROUND"
	NextDayAir ShipmentMethod = "NEXT_DAY_AIR"
	SameDay    ShipmentMethod = "SAME_DAY"
)
