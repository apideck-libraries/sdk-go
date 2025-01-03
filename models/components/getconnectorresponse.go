// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// GetConnectorResponse - Connectors
type GetConnectorResponse struct {
	// HTTP Response Status Code
	StatusCode int64 `json:"status_code"`
	// HTTP Response Status
	Status string    `json:"status"`
	Data   Connector `json:"data"`
	// Response metadata
	Meta *Meta `json:"meta,omitempty"`
	// Links to navigate to previous or next pages through the API
	Links *Links `json:"links,omitempty"`
}

func (o *GetConnectorResponse) GetStatusCode() int64 {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetConnectorResponse) GetStatus() string {
	if o == nil {
		return ""
	}
	return o.Status
}

func (o *GetConnectorResponse) GetData() Connector {
	if o == nil {
		return Connector{}
	}
	return o.Data
}

func (o *GetConnectorResponse) GetMeta() *Meta {
	if o == nil {
		return nil
	}
	return o.Meta
}

func (o *GetConnectorResponse) GetLinks() *Links {
	if o == nil {
		return nil
	}
	return o.Links
}
