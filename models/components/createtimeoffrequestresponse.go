// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// CreateTimeOffRequestResponse - TimeOffRequests
type CreateTimeOffRequestResponse struct {
	// HTTP Response Status Code
	StatusCode int64 `json:"status_code"`
	// HTTP Response Status
	Status string `json:"status"`
	// Apideck ID of service provider
	Service string `json:"service"`
	// Unified API resource name
	Resource string `json:"resource"`
	// Operation performed
	Operation string `json:"operation"`
	// A object containing a unique identifier for the resource that was created, updated, or deleted.
	Data UnifiedID `json:"data"`
}

func (o *CreateTimeOffRequestResponse) GetStatusCode() int64 {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateTimeOffRequestResponse) GetStatus() string {
	if o == nil {
		return ""
	}
	return o.Status
}

func (o *CreateTimeOffRequestResponse) GetService() string {
	if o == nil {
		return ""
	}
	return o.Service
}

func (o *CreateTimeOffRequestResponse) GetResource() string {
	if o == nil {
		return ""
	}
	return o.Resource
}

func (o *CreateTimeOffRequestResponse) GetOperation() string {
	if o == nil {
		return ""
	}
	return o.Operation
}

func (o *CreateTimeOffRequestResponse) GetData() UnifiedID {
	if o == nil {
		return UnifiedID{}
	}
	return o.Data
}
