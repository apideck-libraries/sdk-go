// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// UpdateTaxRateResponse - TaxRate updated
type UpdateTaxRateResponse struct {
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
	Data      UnifiedID `json:"data"`
}

func (o *UpdateTaxRateResponse) GetStatusCode() int64 {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateTaxRateResponse) GetStatus() string {
	if o == nil {
		return ""
	}
	return o.Status
}

func (o *UpdateTaxRateResponse) GetService() string {
	if o == nil {
		return ""
	}
	return o.Service
}

func (o *UpdateTaxRateResponse) GetResource() string {
	if o == nil {
		return ""
	}
	return o.Resource
}

func (o *UpdateTaxRateResponse) GetOperation() string {
	if o == nil {
		return ""
	}
	return o.Operation
}

func (o *UpdateTaxRateResponse) GetData() UnifiedID {
	if o == nil {
		return UnifiedID{}
	}
	return o.Data
}
