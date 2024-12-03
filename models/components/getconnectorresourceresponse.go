// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// GetConnectorResourceResponse - ConnectorResources
type GetConnectorResourceResponse struct {
	// HTTP Response Status Code
	StatusCode int64 `json:"status_code"`
	// HTTP Response Status
	Status string            `json:"status"`
	Data   ConnectorResource `json:"data"`
	// Response metadata
	Meta *Meta `json:"meta,omitempty"`
	// Links to navigate to previous or next pages through the API
	Links *Links `json:"links,omitempty"`
}

func (o *GetConnectorResourceResponse) GetStatusCode() int64 {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetConnectorResourceResponse) GetStatus() string {
	if o == nil {
		return ""
	}
	return o.Status
}

func (o *GetConnectorResourceResponse) GetData() ConnectorResource {
	if o == nil {
		return ConnectorResource{}
	}
	return o.Data
}

func (o *GetConnectorResourceResponse) GetMeta() *Meta {
	if o == nil {
		return nil
	}
	return o.Meta
}

func (o *GetConnectorResourceResponse) GetLinks() *Links {
	if o == nil {
		return nil
	}
	return o.Links
}
