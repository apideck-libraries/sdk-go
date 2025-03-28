// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// CreateWebhookResponse - Webhooks
type CreateWebhookResponse struct {
	// HTTP Response Status Code
	StatusCode int64 `json:"status_code"`
	// HTTP Response Status
	Status string  `json:"status"`
	Data   Webhook `json:"data"`
	// Raw response from the integration when raw=true query param is provided
	Raw map[string]any `json:"_raw,omitempty"`
}

func (o *CreateWebhookResponse) GetStatusCode() int64 {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateWebhookResponse) GetStatus() string {
	if o == nil {
		return ""
	}
	return o.Status
}

func (o *CreateWebhookResponse) GetData() Webhook {
	if o == nil {
		return Webhook{}
	}
	return o.Data
}

func (o *CreateWebhookResponse) GetRaw() map[string]any {
	if o == nil {
		return nil
	}
	return o.Raw
}
