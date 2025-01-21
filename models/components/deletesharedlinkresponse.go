// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// DeleteSharedLinkResponse - Shared Links
type DeleteSharedLinkResponse struct {
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

func (o *DeleteSharedLinkResponse) GetStatusCode() int64 {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteSharedLinkResponse) GetStatus() string {
	if o == nil {
		return ""
	}
	return o.Status
}

func (o *DeleteSharedLinkResponse) GetService() string {
	if o == nil {
		return ""
	}
	return o.Service
}

func (o *DeleteSharedLinkResponse) GetResource() string {
	if o == nil {
		return ""
	}
	return o.Resource
}

func (o *DeleteSharedLinkResponse) GetOperation() string {
	if o == nil {
		return ""
	}
	return o.Operation
}

func (o *DeleteSharedLinkResponse) GetData() UnifiedID {
	if o == nil {
		return UnifiedID{}
	}
	return o.Data
}
