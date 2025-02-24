// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// DeleteEmployeeResponse - Employees
type DeleteEmployeeResponse struct {
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
	// Raw response from the integration when raw=true query param is provided
	Raw map[string]any `json:"_raw,omitempty"`
}

func (o *DeleteEmployeeResponse) GetStatusCode() int64 {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteEmployeeResponse) GetStatus() string {
	if o == nil {
		return ""
	}
	return o.Status
}

func (o *DeleteEmployeeResponse) GetService() string {
	if o == nil {
		return ""
	}
	return o.Service
}

func (o *DeleteEmployeeResponse) GetResource() string {
	if o == nil {
		return ""
	}
	return o.Resource
}

func (o *DeleteEmployeeResponse) GetOperation() string {
	if o == nil {
		return ""
	}
	return o.Operation
}

func (o *DeleteEmployeeResponse) GetData() UnifiedID {
	if o == nil {
		return UnifiedID{}
	}
	return o.Data
}

func (o *DeleteEmployeeResponse) GetRaw() map[string]any {
	if o == nil {
		return nil
	}
	return o.Raw
}
