// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/apideck-libraries/sdk-go/internal/utils"
	"github.com/apideck-libraries/sdk-go/models/components"
)

type IssueTrackingCollectionTicketCommentsUpdateGlobals struct {
	// ID of the consumer which you want to get or push data from
	ConsumerID *string `header:"style=simple,explode=false,name=x-apideck-consumer-id"`
	// The ID of your Unify application
	AppID *string `header:"style=simple,explode=false,name=x-apideck-app-id"`
}

func (o *IssueTrackingCollectionTicketCommentsUpdateGlobals) GetConsumerID() *string {
	if o == nil {
		return nil
	}
	return o.ConsumerID
}

func (o *IssueTrackingCollectionTicketCommentsUpdateGlobals) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

type IssueTrackingCollectionTicketCommentsUpdateRequest struct {
	// ID of the record you are acting upon.
	ID string `pathParam:"style=simple,explode=false,name=id"`
	// ID of the consumer which you want to get or push data from
	ConsumerID *string `header:"style=simple,explode=false,name=x-apideck-consumer-id"`
	// The ID of your Unify application
	AppID *string `header:"style=simple,explode=false,name=x-apideck-app-id"`
	// Provide the service id you want to call (e.g., pipedrive). Only needed when a consumer has activated multiple integrations for a Unified API.
	ServiceID *string `header:"style=simple,explode=false,name=x-apideck-service-id"`
	// Include raw response. Mostly used for debugging purposes
	Raw *bool `default:"false" queryParam:"style=form,explode=true,name=raw"`
	// The collection ID
	CollectionID string `pathParam:"style=simple,explode=false,name=collection_id"`
	// ID of the ticket you are acting upon.
	TicketID                string                                  `pathParam:"style=simple,explode=false,name=ticket_id"`
	CollectionTicketComment components.CollectionTicketCommentInput `request:"mediaType=application/json"`
}

func (i IssueTrackingCollectionTicketCommentsUpdateRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(i, "", false)
}

func (i *IssueTrackingCollectionTicketCommentsUpdateRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &i, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *IssueTrackingCollectionTicketCommentsUpdateRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *IssueTrackingCollectionTicketCommentsUpdateRequest) GetConsumerID() *string {
	if o == nil {
		return nil
	}
	return o.ConsumerID
}

func (o *IssueTrackingCollectionTicketCommentsUpdateRequest) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

func (o *IssueTrackingCollectionTicketCommentsUpdateRequest) GetServiceID() *string {
	if o == nil {
		return nil
	}
	return o.ServiceID
}

func (o *IssueTrackingCollectionTicketCommentsUpdateRequest) GetRaw() *bool {
	if o == nil {
		return nil
	}
	return o.Raw
}

func (o *IssueTrackingCollectionTicketCommentsUpdateRequest) GetCollectionID() string {
	if o == nil {
		return ""
	}
	return o.CollectionID
}

func (o *IssueTrackingCollectionTicketCommentsUpdateRequest) GetTicketID() string {
	if o == nil {
		return ""
	}
	return o.TicketID
}

func (o *IssueTrackingCollectionTicketCommentsUpdateRequest) GetCollectionTicketComment() components.CollectionTicketCommentInput {
	if o == nil {
		return components.CollectionTicketCommentInput{}
	}
	return o.CollectionTicketComment
}

type IssueTrackingCollectionTicketCommentsUpdateResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Update a Comment
	UpdateCommentResponse *components.UpdateCommentResponse
	// Unexpected error
	UnexpectedErrorResponse *components.UnexpectedErrorResponse
}

func (o *IssueTrackingCollectionTicketCommentsUpdateResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *IssueTrackingCollectionTicketCommentsUpdateResponse) GetUpdateCommentResponse() *components.UpdateCommentResponse {
	if o == nil {
		return nil
	}
	return o.UpdateCommentResponse
}

func (o *IssueTrackingCollectionTicketCommentsUpdateResponse) GetUnexpectedErrorResponse() *components.UnexpectedErrorResponse {
	if o == nil {
		return nil
	}
	return o.UnexpectedErrorResponse
}
