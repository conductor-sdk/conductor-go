//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
//  the License. You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
//  an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
//  specific language governing permissions and limitations under the License.

package shipment_example

type User struct {
	Name        string
	Email       string
	AddressLine string
	City        string
	ZipCode     string
	CountryCode string
	BillingType string
	BillingId   string
}

func NewUser(name string, email string, addressLine string, city string, zipCode string, countryCode string, billingType string, billingId string) *User {
	return &User{
		Name:        name,
		Email:       email,
		AddressLine: addressLine,
		City:        city,
		ZipCode:     zipCode,
		CountryCode: countryCode,
		BillingType: billingType,
		BillingId:   billingId,
	}
}
