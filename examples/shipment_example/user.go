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
