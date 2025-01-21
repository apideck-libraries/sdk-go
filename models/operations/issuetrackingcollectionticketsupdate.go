// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/apideck-libraries/sdk-go/internal/utils"
	"github.com/apideck-libraries/sdk-go/models/components"
)

type IssueTrackingCollectionTicketsUpdateGlobals struct {
	// ID of the consumer which you want to get or push data from
	ConsumerID *string `header:"style=simple,explode=false,name=x-apideck-consumer-id"`
	// The ID of your Unify application
	AppID *string `header:"style=simple,explode=false,name=x-apideck-app-id"`
}

func (o *IssueTrackingCollectionTicketsUpdateGlobals) GetConsumerID() *string {
	if o == nil {
		return nil
	}
	return o.ConsumerID
}

func (o *IssueTrackingCollectionTicketsUpdateGlobals) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

type IssueTrackingCollectionTicketsUpdateRequest struct {
	// ID of the ticket you are acting upon.
	TicketID string `pathParam:"style=simple,explode=false,name=ticket_id"`
	// ID of the consumer which you want to get or push data from
	ConsumerID *string `header:"style=simple,explode=false,name=x-apideck-consumer-id"`
	// The ID of your Unify application
	AppID *string `header:"style=simple,explode=false,name=x-apideck-app-id"`
	// Provide the service id you want to call (e.g., pipedrive). Only needed when a consumer has activated multiple integrations for a Unified API.
	ServiceID *string `header:"style=simple,explode=false,name=x-apideck-service-id"`
	// Include raw response. Mostly used for debugging purposes
	Raw *bool `default:"false" queryParam:"style=form,explode=true,name=raw"`
	// The collection ID
	CollectionID string                 `pathParam:"style=simple,explode=false,name=collection_id"`
	Ticket       components.TicketInput `request:"mediaType=application/json"`
}

func (i IssueTrackingCollectionTicketsUpdateRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(i, "", false)
}

func (i *IssueTrackingCollectionTicketsUpdateRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &i, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *IssueTrackingCollectionTicketsUpdateRequest) GetTicketID() string {
	if o == nil {
		return ""
	}
	return o.TicketID
}

func (o *IssueTrackingCollectionTicketsUpdateRequest) GetConsumerID() *string {
	if o == nil {
		return nil
	}
	return o.ConsumerID
}

func (o *IssueTrackingCollectionTicketsUpdateRequest) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

func (o *IssueTrackingCollectionTicketsUpdateRequest) GetServiceID() *string {
	if o == nil {
		return nil
	}
	return o.ServiceID
}

func (o *IssueTrackingCollectionTicketsUpdateRequest) GetRaw() *bool {
	if o == nil {
		return nil
	}
	return o.Raw
}

func (o *IssueTrackingCollectionTicketsUpdateRequest) GetCollectionID() string {
	if o == nil {
		return ""
	}
	return o.CollectionID
}

func (o *IssueTrackingCollectionTicketsUpdateRequest) GetTicket() components.TicketInput {
	if o == nil {
		return components.TicketInput{}
	}
	return o.Ticket
}

type IssueTrackingCollectionTicketsUpdateResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Update a Ticket
	UpdateTicketResponse *components.UpdateTicketResponse
	// Unexpected error
	UnexpectedErrorResponse *components.UnexpectedErrorResponse
}

func (o *IssueTrackingCollectionTicketsUpdateResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *IssueTrackingCollectionTicketsUpdateResponse) GetUpdateTicketResponse() *components.UpdateTicketResponse {
	if o == nil {
		return nil
	}
	return o.UpdateTicketResponse
}

func (o *IssueTrackingCollectionTicketsUpdateResponse) GetUnexpectedErrorResponse() *components.UnexpectedErrorResponse {
	if o == nil {
		return nil
	}
	return o.UnexpectedErrorResponse
}
