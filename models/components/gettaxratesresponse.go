// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// GetTaxRatesResponse - TaxRates
type GetTaxRatesResponse struct {
	// HTTP Response Status Code
	StatusCode int64 `json:"status_code"`
	// HTTP Response Status
	Status string `json:"status"`
	// Apideck ID of service provider
	Service string `json:"service"`
	// Unified API resource name
	Resource string `json:"resource"`
	// Operation performed
	Operation string    `json:"operation"`
	Data      []TaxRate `json:"data"`
	// Response metadata
	Meta *Meta `json:"meta,omitempty"`
	// Links to navigate to previous or next pages through the API
	Links *Links `json:"links,omitempty"`
	// Raw response from the integration when raw=true query param is provided
	Raw map[string]any `json:"_raw,omitempty"`
}

func (o *GetTaxRatesResponse) GetStatusCode() int64 {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetTaxRatesResponse) GetStatus() string {
	if o == nil {
		return ""
	}
	return o.Status
}

func (o *GetTaxRatesResponse) GetService() string {
	if o == nil {
		return ""
	}
	return o.Service
}

func (o *GetTaxRatesResponse) GetResource() string {
	if o == nil {
		return ""
	}
	return o.Resource
}

func (o *GetTaxRatesResponse) GetOperation() string {
	if o == nil {
		return ""
	}
	return o.Operation
}

func (o *GetTaxRatesResponse) GetData() []TaxRate {
	if o == nil {
		return []TaxRate{}
	}
	return o.Data
}

func (o *GetTaxRatesResponse) GetMeta() *Meta {
	if o == nil {
		return nil
	}
	return o.Meta
}

func (o *GetTaxRatesResponse) GetLinks() *Links {
	if o == nil {
		return nil
	}
	return o.Links
}

func (o *GetTaxRatesResponse) GetRaw() map[string]any {
	if o == nil {
		return nil
	}
	return o.Raw
}
