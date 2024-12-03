// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// GetEmployeesResponse - Employees
type GetEmployeesResponse struct {
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
	Data      []Employee `json:"data"`
	// Response metadata
	Meta *Meta `json:"meta,omitempty"`
	// Links to navigate to previous or next pages through the API
	Links *Links `json:"links,omitempty"`
}

func (o *GetEmployeesResponse) GetStatusCode() int64 {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetEmployeesResponse) GetStatus() string {
	if o == nil {
		return ""
	}
	return o.Status
}

func (o *GetEmployeesResponse) GetService() string {
	if o == nil {
		return ""
	}
	return o.Service
}

func (o *GetEmployeesResponse) GetResource() string {
	if o == nil {
		return ""
	}
	return o.Resource
}

func (o *GetEmployeesResponse) GetOperation() string {
	if o == nil {
		return ""
	}
	return o.Operation
}

func (o *GetEmployeesResponse) GetData() []Employee {
	if o == nil {
		return []Employee{}
	}
	return o.Data
}

func (o *GetEmployeesResponse) GetMeta() *Meta {
	if o == nil {
		return nil
	}
	return o.Meta
}

func (o *GetEmployeesResponse) GetLinks() *Links {
	if o == nil {
		return nil
	}
	return o.Links
}
