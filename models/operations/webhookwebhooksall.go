// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/apideck-libraries/sdk-go/internal/utils"
	"github.com/apideck-libraries/sdk-go/models/components"
)

type WebhookWebhooksAllGlobals struct {
	// The ID of your Unify application
	AppID *string `header:"style=simple,explode=false,name=x-apideck-app-id"`
}

func (o *WebhookWebhooksAllGlobals) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

type WebhookWebhooksAllRequest struct {
	// Cursor to start from. You can find cursors for next/previous pages in the meta.cursors property of the response.
	Cursor *string `queryParam:"style=form,explode=true,name=cursor"`
	// Number of results to return. Minimum 1, Maximum 200, Default 20
	Limit *int64 `default:"20" queryParam:"style=form,explode=true,name=limit"`
}

func (w WebhookWebhooksAllRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(w, "", false)
}

func (w *WebhookWebhooksAllRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &w, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *WebhookWebhooksAllRequest) GetCursor() *string {
	if o == nil {
		return nil
	}
	return o.Cursor
}

func (o *WebhookWebhooksAllRequest) GetLimit() *int64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

type WebhookWebhooksAllResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Webhooks
	GetWebhooksResponse *components.GetWebhooksResponse
	// Unexpected error
	UnexpectedErrorResponse *components.UnexpectedErrorResponse

	Next func() (*WebhookWebhooksAllResponse, error)
}

func (o *WebhookWebhooksAllResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *WebhookWebhooksAllResponse) GetGetWebhooksResponse() *components.GetWebhooksResponse {
	if o == nil {
		return nil
	}
	return o.GetWebhooksResponse
}

func (o *WebhookWebhooksAllResponse) GetUnexpectedErrorResponse() *components.UnexpectedErrorResponse {
	if o == nil {
		return nil
	}
	return o.UnexpectedErrorResponse
}
