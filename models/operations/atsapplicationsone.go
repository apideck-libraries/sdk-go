// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/apideck-libraries/sdk-go/internal/utils"
	"github.com/apideck-libraries/sdk-go/models/components"
)

type AtsApplicationsOneGlobals struct {
	// ID of the consumer which you want to get or push data from
	ConsumerID *string `header:"style=simple,explode=false,name=x-apideck-consumer-id"`
	// The ID of your Unify application
	AppID *string `header:"style=simple,explode=false,name=x-apideck-app-id"`
}

func (o *AtsApplicationsOneGlobals) GetConsumerID() *string {
	if o == nil {
		return nil
	}
	return o.ConsumerID
}

func (o *AtsApplicationsOneGlobals) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

type AtsApplicationsOneRequest struct {
	// ID of the record you are acting upon.
	ID string `pathParam:"style=simple,explode=false,name=id"`
	// Provide the service id you want to call (e.g., pipedrive). Only needed when a consumer has activated multiple integrations for a Unified API.
	ServiceID *string `header:"style=simple,explode=false,name=x-apideck-service-id"`
	// Include raw response. Mostly used for debugging purposes
	Raw *bool `default:"false" queryParam:"style=form,explode=true,name=raw"`
}

func (a AtsApplicationsOneRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AtsApplicationsOneRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *AtsApplicationsOneRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *AtsApplicationsOneRequest) GetServiceID() *string {
	if o == nil {
		return nil
	}
	return o.ServiceID
}

func (o *AtsApplicationsOneRequest) GetRaw() *bool {
	if o == nil {
		return nil
	}
	return o.Raw
}

type AtsApplicationsOneResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Applications
	GetApplicationResponse *components.GetApplicationResponse
	// Unexpected error
	UnexpectedErrorResponse *components.UnexpectedErrorResponse
}

func (o *AtsApplicationsOneResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *AtsApplicationsOneResponse) GetGetApplicationResponse() *components.GetApplicationResponse {
	if o == nil {
		return nil
	}
	return o.GetApplicationResponse
}

func (o *AtsApplicationsOneResponse) GetUnexpectedErrorResponse() *components.UnexpectedErrorResponse {
	if o == nil {
		return nil
	}
	return o.UnexpectedErrorResponse
}
