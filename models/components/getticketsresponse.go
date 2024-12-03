// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// GetTicketsResponse - List Tickets
type GetTicketsResponse struct {
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
	Data      []Ticket `json:"data"`
	// Response metadata
	Meta *Meta `json:"meta,omitempty"`
	// Links to navigate to previous or next pages through the API
	Links *Links `json:"links,omitempty"`
}

func (o *GetTicketsResponse) GetStatusCode() int64 {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetTicketsResponse) GetStatus() string {
	if o == nil {
		return ""
	}
	return o.Status
}

func (o *GetTicketsResponse) GetService() string {
	if o == nil {
		return ""
	}
	return o.Service
}

func (o *GetTicketsResponse) GetResource() string {
	if o == nil {
		return ""
	}
	return o.Resource
}

func (o *GetTicketsResponse) GetOperation() string {
	if o == nil {
		return ""
	}
	return o.Operation
}

func (o *GetTicketsResponse) GetData() []Ticket {
	if o == nil {
		return []Ticket{}
	}
	return o.Data
}

func (o *GetTicketsResponse) GetMeta() *Meta {
	if o == nil {
		return nil
	}
	return o.Meta
}

func (o *GetTicketsResponse) GetLinks() *Links {
	if o == nil {
		return nil
	}
	return o.Links
}
