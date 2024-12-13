// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/apideck-libraries/sdk-go/internal/utils"
	"github.com/apideck-libraries/sdk-go/models/components"
)

type IssueTrackingCollectionTicketCommentsAllGlobals struct {
	// ID of the consumer which you want to get or push data from
	ConsumerID *string `header:"style=simple,explode=false,name=x-apideck-consumer-id"`
	// The ID of your Unify application
	AppID *string `header:"style=simple,explode=false,name=x-apideck-app-id"`
}

func (o *IssueTrackingCollectionTicketCommentsAllGlobals) GetConsumerID() *string {
	if o == nil {
		return nil
	}
	return o.ConsumerID
}

func (o *IssueTrackingCollectionTicketCommentsAllGlobals) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

type IssueTrackingCollectionTicketCommentsAllRequest struct {
	// The collection ID
	CollectionID string `pathParam:"style=simple,explode=false,name=collection_id"`
	// ID of the ticket you are acting upon.
	TicketID string `pathParam:"style=simple,explode=false,name=ticket_id"`
	// Include raw response. Mostly used for debugging purposes
	Raw *bool `default:"false" queryParam:"style=form,explode=true,name=raw"`
	// Provide the service id you want to call (e.g., pipedrive). Only needed when a consumer has activated multiple integrations for a Unified API.
	ServiceID *string `header:"style=simple,explode=false,name=x-apideck-service-id"`
	// Cursor to start from. You can find cursors for next/previous pages in the meta.cursors property of the response.
	Cursor *string `queryParam:"style=form,explode=true,name=cursor"`
	// Number of results to return. Minimum 1, Maximum 200, Default 20
	Limit *int64 `default:"20" queryParam:"style=form,explode=true,name=limit"`
	// Apply sorting
	Sort *components.CommentsSort `queryParam:"style=deepObject,explode=true,name=sort"`
	// Optional unmapped key/values that will be passed through to downstream as query parameters. Ie: ?pass_through[search]=leads becomes ?search=leads
	PassThrough map[string]any `queryParam:"style=deepObject,explode=true,name=pass_through"`
	// The 'fields' parameter allows API users to specify the fields they want to include in the API response. If this parameter is not present, the API will return all available fields. If this parameter is present, only the fields specified in the comma-separated string will be included in the response. Nested properties can also be requested by using a dot notation. <br /><br />Example: `fields=name,email,addresses.city`<br /><br />In the example above, the response will only include the fields "name", "email" and "addresses.city". If any other fields are available, they will be excluded.
	Fields *string `queryParam:"style=form,explode=true,name=fields"`
}

func (i IssueTrackingCollectionTicketCommentsAllRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(i, "", false)
}

func (i *IssueTrackingCollectionTicketCommentsAllRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &i, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *IssueTrackingCollectionTicketCommentsAllRequest) GetCollectionID() string {
	if o == nil {
		return ""
	}
	return o.CollectionID
}

func (o *IssueTrackingCollectionTicketCommentsAllRequest) GetTicketID() string {
	if o == nil {
		return ""
	}
	return o.TicketID
}

func (o *IssueTrackingCollectionTicketCommentsAllRequest) GetRaw() *bool {
	if o == nil {
		return nil
	}
	return o.Raw
}

func (o *IssueTrackingCollectionTicketCommentsAllRequest) GetServiceID() *string {
	if o == nil {
		return nil
	}
	return o.ServiceID
}

func (o *IssueTrackingCollectionTicketCommentsAllRequest) GetCursor() *string {
	if o == nil {
		return nil
	}
	return o.Cursor
}

func (o *IssueTrackingCollectionTicketCommentsAllRequest) GetLimit() *int64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *IssueTrackingCollectionTicketCommentsAllRequest) GetSort() *components.CommentsSort {
	if o == nil {
		return nil
	}
	return o.Sort
}

func (o *IssueTrackingCollectionTicketCommentsAllRequest) GetPassThrough() map[string]any {
	if o == nil {
		return nil
	}
	return o.PassThrough
}

func (o *IssueTrackingCollectionTicketCommentsAllRequest) GetFields() *string {
	if o == nil {
		return nil
	}
	return o.Fields
}

type IssueTrackingCollectionTicketCommentsAllResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// List Comments
	GetCommentsResponse *components.GetCommentsResponse
	// Unexpected error
	UnexpectedErrorResponse *components.UnexpectedErrorResponse

	Next func() (*IssueTrackingCollectionTicketCommentsAllResponse, error)
}

func (o *IssueTrackingCollectionTicketCommentsAllResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *IssueTrackingCollectionTicketCommentsAllResponse) GetGetCommentsResponse() *components.GetCommentsResponse {
	if o == nil {
		return nil
	}
	return o.GetCommentsResponse
}

func (o *IssueTrackingCollectionTicketCommentsAllResponse) GetUnexpectedErrorResponse() *components.UnexpectedErrorResponse {
	if o == nil {
		return nil
	}
	return o.UnexpectedErrorResponse
}
