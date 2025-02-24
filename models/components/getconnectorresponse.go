// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// GetConnectorResponse - Connectors
type GetConnectorResponse struct {
	// HTTP Response Status Code
	StatusCode int64 `json:"status_code"`
	// HTTP Response Status
	Status string    `json:"status"`
	Data   Connector `json:"data"`
	// Raw response from the integration when raw=true query param is provided
	Raw map[string]any `json:"_raw,omitempty"`
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

func (o *GetConnectorResponse) GetRaw() map[string]any {
	if o == nil {
		return nil
	}
	return o.Raw
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
