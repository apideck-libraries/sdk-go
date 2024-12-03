// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package sdkgo

type Ecommerce struct {
	Orders    *Orders
	Products  *Products
	Customers *ApideckCustomers
	Stores    *Stores

	sdkConfiguration sdkConfiguration
}

func newEcommerce(sdkConfig sdkConfiguration) *Ecommerce {
	return &Ecommerce{
		sdkConfiguration: sdkConfig,
		Orders:           newOrders(sdkConfig),
		Products:         newProducts(sdkConfig),
		Customers:        newApideckCustomers(sdkConfig),
		Stores:           newStores(sdkConfig),
	}
}
