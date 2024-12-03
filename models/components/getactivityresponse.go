// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// GetActivityResponse - Activity
type GetActivityResponse struct {
	// HTTP Response Status Code
	StatusCode int64 `json:"status_code"`
	// HTTP Response Status
	Status string `json:"status"`
	// Apideck ID of service provider
	Service string `json:"service"`
	// Unified API resource name
	Resource string `json:"resource"`
	// Operation performed
	Operation string   `json:"operation"`
	Data      Activity `json:"data"`
}

func (o *GetActivityResponse) GetStatusCode() int64 {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetActivityResponse) GetStatus() string {
	if o == nil {
		return ""
	}
	return o.Status
}

func (o *GetActivityResponse) GetService() string {
	if o == nil {
		return ""
	}
	return o.Service
}

func (o *GetActivityResponse) GetResource() string {
	if o == nil {
		return ""
	}
	return o.Resource
}

func (o *GetActivityResponse) GetOperation() string {
	if o == nil {
		return ""
	}
	return o.Operation
}

func (o *GetActivityResponse) GetData() Activity {
	if o == nil {
		return Activity{}
	}
	return o.Data
}
