// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// GetWebhookEventLogsResponse - EventLogs
type GetWebhookEventLogsResponse struct {
	// HTTP Response Status Code
	StatusCode int64 `json:"status_code"`
	// HTTP Response Status
	Status string            `json:"status"`
	Data   []WebhookEventLog `json:"data"`
	// Response metadata
	Meta *Meta `json:"meta,omitempty"`
	// Links to navigate to previous or next pages through the API
	Links *Links `json:"links,omitempty"`
	// Raw response from the integration when raw=true query param is provided
	Raw map[string]any `json:"_raw,omitempty"`
}

func (o *GetWebhookEventLogsResponse) GetStatusCode() int64 {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetWebhookEventLogsResponse) GetStatus() string {
	if o == nil {
		return ""
	}
	return o.Status
}

func (o *GetWebhookEventLogsResponse) GetData() []WebhookEventLog {
	if o == nil {
		return []WebhookEventLog{}
	}
	return o.Data
}

func (o *GetWebhookEventLogsResponse) GetMeta() *Meta {
	if o == nil {
		return nil
	}
	return o.Meta
}

func (o *GetWebhookEventLogsResponse) GetLinks() *Links {
	if o == nil {
		return nil
	}
	return o.Links
}

func (o *GetWebhookEventLogsResponse) GetRaw() map[string]any {
	if o == nil {
		return nil
	}
	return o.Raw
}
