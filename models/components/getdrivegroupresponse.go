// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// GetDriveGroupResponse - DriveGroups
type GetDriveGroupResponse struct {
	// HTTP Response Status Code
	StatusCode int64 `json:"status_code"`
	// HTTP Response Status
	Status string `json:"status"`
	// Apideck ID of service provider
	Service string `json:"service"`
	// Unified API resource name
	Resource string `json:"resource"`
	// Operation performed
	Operation string     `json:"operation"`
	Data      DriveGroup `json:"data"`
}

func (o *GetDriveGroupResponse) GetStatusCode() int64 {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetDriveGroupResponse) GetStatus() string {
	if o == nil {
		return ""
	}
	return o.Status
}

func (o *GetDriveGroupResponse) GetService() string {
	if o == nil {
		return ""
	}
	return o.Service
}

func (o *GetDriveGroupResponse) GetResource() string {
	if o == nil {
		return ""
	}
	return o.Resource
}

func (o *GetDriveGroupResponse) GetOperation() string {
	if o == nil {
		return ""
	}
	return o.Operation
}

func (o *GetDriveGroupResponse) GetData() DriveGroup {
	if o == nil {
		return DriveGroup{}
	}
	return o.Data
}
