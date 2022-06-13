//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
//  the License. You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
//  an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
//  specific language governing permissions and limitations under the License.

package shipment

type Order struct {
	OrderNumber    string
	Sku            string
	Quantity       int
	UnitPrice      float64
	ZipCode        string
	CountryCode    string
	ShippingMethod ShipmentMethod
}

func NewOrder(orderNumber string, sku string, quantity int, unitPrice float64) *Order {
	return &Order{
		OrderNumber: orderNumber,
		Sku:         sku,
		Quantity:    quantity,
		UnitPrice:   unitPrice,
	}
}
