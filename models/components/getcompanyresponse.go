// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// GetCompanyResponse - Company
type GetCompanyResponse struct {
	// HTTP Response Status Code
	StatusCode int64 `json:"status_code"`
	// HTTP Response Status
	Status string `json:"status"`
	// Apideck ID of service provider
	Service string `json:"service"`
	// Unified API resource name
	Resource string `json:"resource"`
	// Operation performed
	Operation string  `json:"operation"`
	Data      Company `json:"data"`
}

func (o *GetCompanyResponse) GetStatusCode() int64 {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetCompanyResponse) GetStatus() string {
	if o == nil {
		return ""
	}
	return o.Status
}

func (o *GetCompanyResponse) GetService() string {
	if o == nil {
		return ""
	}
	return o.Service
}

func (o *GetCompanyResponse) GetResource() string {
	if o == nil {
		return ""
	}
	return o.Resource
}

func (o *GetCompanyResponse) GetOperation() string {
	if o == nil {
		return ""
	}
	return o.Operation
}

func (o *GetCompanyResponse) GetData() Company {
	if o == nil {
		return Company{}
	}
	return o.Data
}
