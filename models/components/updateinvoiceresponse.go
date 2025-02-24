// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// UpdateInvoiceResponse - Invoice updated
type UpdateInvoiceResponse struct {
	// HTTP Response Status Code
	StatusCode int64 `json:"status_code"`
	// HTTP Response Status
	Status string `json:"status"`
	// Apideck ID of service provider
	Service string `json:"service"`
	// Unified API resource name
	Resource string `json:"resource"`
	// Operation performed
	Operation string          `json:"operation"`
	Data      InvoiceResponse `json:"data"`
	// Raw response from the integration when raw=true query param is provided
	Raw map[string]any `json:"_raw,omitempty"`
}

func (o *UpdateInvoiceResponse) GetStatusCode() int64 {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateInvoiceResponse) GetStatus() string {
	if o == nil {
		return ""
	}
	return o.Status
}

func (o *UpdateInvoiceResponse) GetService() string {
	if o == nil {
		return ""
	}
	return o.Service
}

func (o *UpdateInvoiceResponse) GetResource() string {
	if o == nil {
		return ""
	}
	return o.Resource
}

func (o *UpdateInvoiceResponse) GetOperation() string {
	if o == nil {
		return ""
	}
	return o.Operation
}

func (o *UpdateInvoiceResponse) GetData() InvoiceResponse {
	if o == nil {
		return InvoiceResponse{}
	}
	return o.Data
}

func (o *UpdateInvoiceResponse) GetRaw() map[string]any {
	if o == nil {
		return nil
	}
	return o.Raw
}
