// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// DeleteAttachmentResponse - Attachments
type DeleteAttachmentResponse struct {
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
}

func (o *DeleteAttachmentResponse) GetStatusCode() int64 {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteAttachmentResponse) GetStatus() string {
	if o == nil {
		return ""
	}
	return o.Status
}

func (o *DeleteAttachmentResponse) GetService() string {
	if o == nil {
		return ""
	}
	return o.Service
}

func (o *DeleteAttachmentResponse) GetResource() string {
	if o == nil {
		return ""
	}
	return o.Resource
}

func (o *DeleteAttachmentResponse) GetOperation() string {
	if o == nil {
		return ""
	}
	return o.Operation
}

func (o *DeleteAttachmentResponse) GetData() UnifiedID {
	if o == nil {
		return UnifiedID{}
	}
	return o.Data
}
