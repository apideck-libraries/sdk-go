// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// GetOpportunityResponse - Opportunity
type GetOpportunityResponse struct {
	// HTTP Response Status Code
	StatusCode int64 `json:"status_code"`
	// HTTP Response Status
	Status string `json:"status"`
	// Apideck ID of service provider
	Service string `json:"service"`
	// Unified API resource name
	Resource string `json:"resource"`
	// Operation performed
	Operation string      `json:"operation"`
	Data      Opportunity `json:"data"`
	// Raw response from the integration when raw=true query param is provided
	Raw map[string]any `json:"_raw,omitempty"`
}

func (o *GetOpportunityResponse) GetStatusCode() int64 {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetOpportunityResponse) GetStatus() string {
	if o == nil {
		return ""
	}
	return o.Status
}

func (o *GetOpportunityResponse) GetService() string {
	if o == nil {
		return ""
	}
	return o.Service
}

func (o *GetOpportunityResponse) GetResource() string {
	if o == nil {
		return ""
	}
	return o.Resource
}

func (o *GetOpportunityResponse) GetOperation() string {
	if o == nil {
		return ""
	}
	return o.Operation
}

func (o *GetOpportunityResponse) GetData() Opportunity {
	if o == nil {
		return Opportunity{}
	}
	return o.Data
}

func (o *GetOpportunityResponse) GetRaw() map[string]any {
	if o == nil {
		return nil
	}
	return o.Raw
}
