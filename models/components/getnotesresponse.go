// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// GetNotesResponse - Notes
type GetNotesResponse struct {
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
	Data      []Note `json:"data"`
	// Response metadata
	Meta *Meta `json:"meta,omitempty"`
	// Links to navigate to previous or next pages through the API
	Links *Links `json:"links,omitempty"`
}

func (o *GetNotesResponse) GetStatusCode() int64 {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetNotesResponse) GetStatus() string {
	if o == nil {
		return ""
	}
	return o.Status
}

func (o *GetNotesResponse) GetService() string {
	if o == nil {
		return ""
	}
	return o.Service
}

func (o *GetNotesResponse) GetResource() string {
	if o == nil {
		return ""
	}
	return o.Resource
}

func (o *GetNotesResponse) GetOperation() string {
	if o == nil {
		return ""
	}
	return o.Operation
}

func (o *GetNotesResponse) GetData() []Note {
	if o == nil {
		return []Note{}
	}
	return o.Data
}

func (o *GetNotesResponse) GetMeta() *Meta {
	if o == nil {
		return nil
	}
	return o.Meta
}

func (o *GetNotesResponse) GetLinks() *Links {
	if o == nil {
		return nil
	}
	return o.Links
}
