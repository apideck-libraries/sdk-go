// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type DeleteConsumerResponseData struct {
	// Unique consumer identifier. You can freely choose a consumer ID yourself. Most of the time, this is an ID of your internal data model that represents a user or account in your system (for example account:12345). If the consumer doesn't exist yet, Vault will upsert a consumer based on your ID.
	ConsumerID *string `json:"consumer_id,omitempty"`
}

func (o *DeleteConsumerResponseData) GetConsumerID() *string {
	if o == nil {
		return nil
	}
	return o.ConsumerID
}

// DeleteConsumerResponse - Consumer deleted
type DeleteConsumerResponse struct {
	// HTTP Response Status Code
	StatusCode int64 `json:"status_code"`
	// HTTP Response Status
	Status string                     `json:"status"`
	Data   DeleteConsumerResponseData `json:"data"`
	// Raw response from the integration when raw=true query param is provided
	Raw map[string]any `json:"_raw,omitempty"`
}

func (o *DeleteConsumerResponse) GetStatusCode() int64 {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteConsumerResponse) GetStatus() string {
	if o == nil {
		return ""
	}
	return o.Status
}

func (o *DeleteConsumerResponse) GetData() DeleteConsumerResponseData {
	if o == nil {
		return DeleteConsumerResponseData{}
	}
	return o.Data
}

func (o *DeleteConsumerResponse) GetRaw() map[string]any {
	if o == nil {
		return nil
	}
	return o.Raw
}
