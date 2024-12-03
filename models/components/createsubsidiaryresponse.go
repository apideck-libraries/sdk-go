// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// CreateSubsidiaryResponse - Subsidiaries
type CreateSubsidiaryResponse struct {
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

func (o *CreateSubsidiaryResponse) GetStatusCode() int64 {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateSubsidiaryResponse) GetStatus() string {
	if o == nil {
		return ""
	}
	return o.Status
}

func (o *CreateSubsidiaryResponse) GetService() string {
	if o == nil {
		return ""
	}
	return o.Service
}

func (o *CreateSubsidiaryResponse) GetResource() string {
	if o == nil {
		return ""
	}
	return o.Resource
}

func (o *CreateSubsidiaryResponse) GetOperation() string {
	if o == nil {
		return ""
	}
	return o.Operation
}

func (o *CreateSubsidiaryResponse) GetData() UnifiedID {
	if o == nil {
		return UnifiedID{}
	}
	return o.Data
}
