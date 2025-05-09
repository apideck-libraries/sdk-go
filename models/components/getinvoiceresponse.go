// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// GetInvoiceResponse - Invoice
type GetInvoiceResponse struct {
	// HTTP Response Status Code
	StatusCode int64 `json:"status_code"`
	// HTTP Response Status
	Status string `json:"status"`
	// Apideck ID of service provider
	Service string `json:"service"`
	// Unified API resource name
	Resource string `json:"resource"`
	// Operation performed
	Operation string  `json:"operation"`
	Data      Invoice `json:"data"`
	// Raw response from the integration when raw=true query param is provided
	Raw map[string]any `json:"_raw,omitempty"`
}

func (o *GetInvoiceResponse) GetStatusCode() int64 {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetInvoiceResponse) GetStatus() string {
	if o == nil {
		return ""
	}
	return o.Status
}

func (o *GetInvoiceResponse) GetService() string {
	if o == nil {
		return ""
	}
	return o.Service
}

func (o *GetInvoiceResponse) GetResource() string {
	if o == nil {
		return ""
	}
	return o.Resource
}

func (o *GetInvoiceResponse) GetOperation() string {
	if o == nil {
		return ""
	}
	return o.Operation
}

func (o *GetInvoiceResponse) GetData() Invoice {
	if o == nil {
		return Invoice{}
	}
	return o.Data
}

func (o *GetInvoiceResponse) GetRaw() map[string]any {
	if o == nil {
		return nil
	}
	return o.Raw
}
