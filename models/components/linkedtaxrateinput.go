// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type LinkedTaxRateInput struct {
	// The ID of the object.
	ID *string `json:"id,omitempty"`
	// Rate of the tax rate
	Rate *float64 `json:"rate,omitempty"`
}

func (o *LinkedTaxRateInput) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *LinkedTaxRateInput) GetRate() *float64 {
	if o == nil {
		return nil
	}
	return o.Rate
}
