// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// GetProfitAndLossResponse - Profit & Loss Report
type GetProfitAndLossResponse struct {
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
	Data      ProfitAndLoss `json:"data"`
	// Raw response from the integration when raw=true query param is provided
	Raw map[string]any `json:"_raw,omitempty"`
}

func (o *GetProfitAndLossResponse) GetStatusCode() int64 {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetProfitAndLossResponse) GetStatus() string {
	if o == nil {
		return ""
	}
	return o.Status
}

func (o *GetProfitAndLossResponse) GetService() string {
	if o == nil {
		return ""
	}
	return o.Service
}

func (o *GetProfitAndLossResponse) GetResource() string {
	if o == nil {
		return ""
	}
	return o.Resource
}

func (o *GetProfitAndLossResponse) GetOperation() string {
	if o == nil {
		return ""
	}
	return o.Operation
}

func (o *GetProfitAndLossResponse) GetData() ProfitAndLoss {
	if o == nil {
		return ProfitAndLoss{}
	}
	return o.Data
}

func (o *GetProfitAndLossResponse) GetRaw() map[string]any {
	if o == nil {
		return nil
	}
	return o.Raw
}
