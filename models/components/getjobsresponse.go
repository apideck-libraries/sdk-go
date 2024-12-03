// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// GetJobsResponse - Jobs
type GetJobsResponse struct {
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
	Data      []Job  `json:"data"`
	// Response metadata
	Meta *Meta `json:"meta,omitempty"`
	// Links to navigate to previous or next pages through the API
	Links *Links `json:"links,omitempty"`
}

func (o *GetJobsResponse) GetStatusCode() int64 {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetJobsResponse) GetStatus() string {
	if o == nil {
		return ""
	}
	return o.Status
}

func (o *GetJobsResponse) GetService() string {
	if o == nil {
		return ""
	}
	return o.Service
}

func (o *GetJobsResponse) GetResource() string {
	if o == nil {
		return ""
	}
	return o.Resource
}

func (o *GetJobsResponse) GetOperation() string {
	if o == nil {
		return ""
	}
	return o.Operation
}

func (o *GetJobsResponse) GetData() []Job {
	if o == nil {
		return []Job{}
	}
	return o.Data
}

func (o *GetJobsResponse) GetMeta() *Meta {
	if o == nil {
		return nil
	}
	return o.Meta
}

func (o *GetJobsResponse) GetLinks() *Links {
	if o == nil {
		return nil
	}
	return o.Links
}
