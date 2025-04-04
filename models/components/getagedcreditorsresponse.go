// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// GetAgedCreditorsResponse - Aged Creditors
type GetAgedCreditorsResponse struct {
	// HTTP Response Status Code
	StatusCode int64 `json:"status_code"`
	// HTTP Response Status
	Status string `json:"status"`
	// Apideck ID of service provider
	Service string `json:"service"`
	// Unified API resource name
	Resource string `json:"resource"`
	// Operation performed
	Operation string        `json:"operation"`
	Data      AgedCreditors `json:"data"`
	// Raw response from the integration when raw=true query param is provided
	Raw map[string]any `json:"_raw,omitempty"`
}

func (o *GetAgedCreditorsResponse) GetStatusCode() int64 {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetAgedCreditorsResponse) GetStatus() string {
	if o == nil {
		return ""
	}
	return o.Status
}

func (o *GetAgedCreditorsResponse) GetService() string {
	if o == nil {
		return ""
	}
	return o.Service
}

func (o *GetAgedCreditorsResponse) GetResource() string {
	if o == nil {
		return ""
	}
	return o.Resource
}

func (o *GetAgedCreditorsResponse) GetOperation() string {
	if o == nil {
		return ""
	}
	return o.Operation
}

func (o *GetAgedCreditorsResponse) GetData() AgedCreditors {
	if o == nil {
		return AgedCreditors{}
	}
	return o.Data
}

func (o *GetAgedCreditorsResponse) GetRaw() map[string]any {
	if o == nil {
		return nil
	}
	return o.Raw
}
