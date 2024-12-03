// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// GetMessagesResponse - Messages
type GetMessagesResponse struct {
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
	Data      []Message `json:"data"`
	// Response metadata
	Meta *Meta `json:"meta,omitempty"`
	// Links to navigate to previous or next pages through the API
	Links *Links `json:"links,omitempty"`
}

func (o *GetMessagesResponse) GetStatusCode() int64 {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetMessagesResponse) GetStatus() string {
	if o == nil {
		return ""
	}
	return o.Status
}

func (o *GetMessagesResponse) GetService() string {
	if o == nil {
		return ""
	}
	return o.Service
}

func (o *GetMessagesResponse) GetResource() string {
	if o == nil {
		return ""
	}
	return o.Resource
}

func (o *GetMessagesResponse) GetOperation() string {
	if o == nil {
		return ""
	}
	return o.Operation
}

func (o *GetMessagesResponse) GetData() []Message {
	if o == nil {
		return []Message{}
	}
	return o.Data
}

func (o *GetMessagesResponse) GetMeta() *Meta {
	if o == nil {
		return nil
	}
	return o.Meta
}

func (o *GetMessagesResponse) GetLinks() *Links {
	if o == nil {
		return nil
	}
	return o.Links
}
