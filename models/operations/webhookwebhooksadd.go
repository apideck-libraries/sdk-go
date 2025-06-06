// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

type WebhookWebhooksAddGlobals struct {
	// The ID of your Unify application
	AppID *string `header:"style=simple,explode=false,name=x-apideck-app-id"`
}

func (o *WebhookWebhooksAddGlobals) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

type WebhookWebhooksAddRequest struct {
	// The ID of your Unify application
	AppID                *string                         `header:"style=simple,explode=false,name=x-apideck-app-id"`
	CreateWebhookRequest components.CreateWebhookRequest `request:"mediaType=application/json"`
}

func (o *WebhookWebhooksAddRequest) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

func (o *WebhookWebhooksAddRequest) GetCreateWebhookRequest() components.CreateWebhookRequest {
	if o == nil {
		return components.CreateWebhookRequest{}
	}
	return o.CreateWebhookRequest
}

type WebhookWebhooksAddResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Webhooks
	CreateWebhookResponse *components.CreateWebhookResponse
	// Unexpected error
	UnexpectedErrorResponse *components.UnexpectedErrorResponse
}

func (o *WebhookWebhooksAddResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *WebhookWebhooksAddResponse) GetCreateWebhookResponse() *components.CreateWebhookResponse {
	if o == nil {
		return nil
	}
	return o.CreateWebhookResponse
}

func (o *WebhookWebhooksAddResponse) GetUnexpectedErrorResponse() *components.UnexpectedErrorResponse {
	if o == nil {
		return nil
	}
	return o.UnexpectedErrorResponse
}
