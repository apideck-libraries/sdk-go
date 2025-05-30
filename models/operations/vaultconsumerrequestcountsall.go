// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

type VaultConsumerRequestCountsAllGlobals struct {
	// The ID of your Unify application
	AppID *string `header:"style=simple,explode=false,name=x-apideck-app-id"`
}

func (o *VaultConsumerRequestCountsAllGlobals) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

type VaultConsumerRequestCountsAllRequest struct {
	// The ID of your Unify application
	AppID *string `header:"style=simple,explode=false,name=x-apideck-app-id"`
	// ID of the consumer to return
	ConsumerID string `pathParam:"style=simple,explode=false,name=consumer_id"`
	// Scopes results to requests that happened after datetime
	StartDatetime string `queryParam:"style=form,explode=true,name=start_datetime"`
	// Scopes results to requests that happened before datetime
	EndDatetime string `queryParam:"style=form,explode=true,name=end_datetime"`
}

func (o *VaultConsumerRequestCountsAllRequest) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

func (o *VaultConsumerRequestCountsAllRequest) GetConsumerID() string {
	if o == nil {
		return ""
	}
	return o.ConsumerID
}

func (o *VaultConsumerRequestCountsAllRequest) GetStartDatetime() string {
	if o == nil {
		return ""
	}
	return o.StartDatetime
}

func (o *VaultConsumerRequestCountsAllRequest) GetEndDatetime() string {
	if o == nil {
		return ""
	}
	return o.EndDatetime
}

type VaultConsumerRequestCountsAllResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Consumers Request Counts within Date Range
	ConsumerRequestCountsInDateRangeResponse *components.ConsumerRequestCountsInDateRangeResponse
	// Unexpected error
	UnexpectedErrorResponse *components.UnexpectedErrorResponse
}

func (o *VaultConsumerRequestCountsAllResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *VaultConsumerRequestCountsAllResponse) GetConsumerRequestCountsInDateRangeResponse() *components.ConsumerRequestCountsInDateRangeResponse {
	if o == nil {
		return nil
	}
	return o.ConsumerRequestCountsInDateRangeResponse
}

func (o *VaultConsumerRequestCountsAllResponse) GetUnexpectedErrorResponse() *components.UnexpectedErrorResponse {
	if o == nil {
		return nil
	}
	return o.UnexpectedErrorResponse
}
