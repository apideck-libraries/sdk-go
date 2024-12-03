// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// UpdateDriveResponse - Drives
type UpdateDriveResponse struct {
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

func (o *UpdateDriveResponse) GetStatusCode() int64 {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateDriveResponse) GetStatus() string {
	if o == nil {
		return ""
	}
	return o.Status
}

func (o *UpdateDriveResponse) GetService() string {
	if o == nil {
		return ""
	}
	return o.Service
}

func (o *UpdateDriveResponse) GetResource() string {
	if o == nil {
		return ""
	}
	return o.Resource
}

func (o *UpdateDriveResponse) GetOperation() string {
	if o == nil {
		return ""
	}
	return o.Operation
}

func (o *UpdateDriveResponse) GetData() UnifiedID {
	if o == nil {
		return UnifiedID{}
	}
	return o.Data
}
