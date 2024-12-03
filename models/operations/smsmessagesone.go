// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/apideck-libraries/sdk-go/internal/utils"
	"github.com/apideck-libraries/sdk-go/models/components"
)

type SmsMessagesOneGlobals struct {
	// ID of the consumer which you want to get or push data from
	ConsumerID *string `header:"style=simple,explode=false,name=x-apideck-consumer-id"`
	// The ID of your Unify application
	AppID *string `header:"style=simple,explode=false,name=x-apideck-app-id"`
}

func (o *SmsMessagesOneGlobals) GetConsumerID() *string {
	if o == nil {
		return nil
	}
	return o.ConsumerID
}

func (o *SmsMessagesOneGlobals) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

type SmsMessagesOneRequest struct {
	// ID of the record you are acting upon.
	ID string `pathParam:"style=simple,explode=false,name=id"`
	// Provide the service id you want to call (e.g., pipedrive). Only needed when a consumer has activated multiple integrations for a Unified API.
	ServiceID *string `header:"style=simple,explode=false,name=x-apideck-service-id"`
	// Include raw response. Mostly used for debugging purposes
	Raw *bool `default:"false" queryParam:"style=form,explode=true,name=raw"`
	// The 'fields' parameter allows API users to specify the fields they want to include in the API response. If this parameter is not present, the API will return all available fields. If this parameter is present, only the fields specified in the comma-separated string will be included in the response. Nested properties can also be requested by using a dot notation. <br /><br />Example: `fields=name,email,addresses.city`<br /><br />In the example above, the response will only include the fields "name", "email" and "addresses.city". If any other fields are available, they will be excluded.
	Fields *string `queryParam:"style=form,explode=true,name=fields"`
}

func (s SmsMessagesOneRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SmsMessagesOneRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SmsMessagesOneRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *SmsMessagesOneRequest) GetServiceID() *string {
	if o == nil {
		return nil
	}
	return o.ServiceID
}

func (o *SmsMessagesOneRequest) GetRaw() *bool {
	if o == nil {
		return nil
	}
	return o.Raw
}

func (o *SmsMessagesOneRequest) GetFields() *string {
	if o == nil {
		return nil
	}
	return o.Fields
}

type SmsMessagesOneResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Messages
	GetMessageResponse *components.GetMessageResponse
	// Unexpected error
	UnexpectedErrorResponse *components.UnexpectedErrorResponse
}

func (o *SmsMessagesOneResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *SmsMessagesOneResponse) GetGetMessageResponse() *components.GetMessageResponse {
	if o == nil {
		return nil
	}
	return o.GetMessageResponse
}

func (o *SmsMessagesOneResponse) GetUnexpectedErrorResponse() *components.UnexpectedErrorResponse {
	if o == nil {
		return nil
	}
	return o.UnexpectedErrorResponse
}
