// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type UpdateWebhookRequest struct {
	// A description of the object.
	Description *string `json:"description,omitempty"`
	// The status of the webhook.
	Status *Status `json:"status,omitempty"`
	// The delivery url of the webhook endpoint.
	DeliveryURL *string `json:"delivery_url,omitempty"`
	// The list of subscribed events for this webhook. [`*`] indicates that all events are enabled.
	Events []WebhookEventType `json:"events,omitempty"`
}

func (o *UpdateWebhookRequest) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *UpdateWebhookRequest) GetStatus() *Status {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *UpdateWebhookRequest) GetDeliveryURL() *string {
	if o == nil {
		return nil
	}
	return o.DeliveryURL
}

func (o *UpdateWebhookRequest) GetEvents() []WebhookEventType {
	if o == nil {
		return nil
	}
	return o.Events
}
