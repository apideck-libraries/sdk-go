// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// UpdateConsumerResponse - Consumer updated
type UpdateConsumerResponse struct {
	// HTTP Response Status Code
	StatusCode int64 `json:"status_code"`
	// HTTP Response Status
	Status string   `json:"status"`
	Data   Consumer `json:"data"`
	// Raw response from the integration when raw=true query param is provided
	Raw map[string]any `json:"_raw,omitempty"`
}

func (o *UpdateConsumerResponse) GetStatusCode() int64 {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateConsumerResponse) GetStatus() string {
	if o == nil {
		return ""
	}
	return o.Status
}

func (o *UpdateConsumerResponse) GetData() Consumer {
	if o == nil {
		return Consumer{}
	}
	return o.Data
}

func (o *UpdateConsumerResponse) GetRaw() map[string]any {
	if o == nil {
		return nil
	}
	return o.Raw
}
