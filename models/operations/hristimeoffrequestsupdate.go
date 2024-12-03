// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/apideck-libraries/sdk-go/internal/utils"
	"github.com/apideck-libraries/sdk-go/models/components"
)

type HrisTimeOffRequestsUpdateGlobals struct {
	// ID of the consumer which you want to get or push data from
	ConsumerID *string `header:"style=simple,explode=false,name=x-apideck-consumer-id"`
	// The ID of your Unify application
	AppID *string `header:"style=simple,explode=false,name=x-apideck-app-id"`
}

func (o *HrisTimeOffRequestsUpdateGlobals) GetConsumerID() *string {
	if o == nil {
		return nil
	}
	return o.ConsumerID
}

func (o *HrisTimeOffRequestsUpdateGlobals) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

type HrisTimeOffRequestsUpdateRequest struct {
	// ID of the record you are acting upon.
	ID string `pathParam:"style=simple,explode=false,name=id"`
	// Provide the service id you want to call (e.g., pipedrive). Only needed when a consumer has activated multiple integrations for a Unified API.
	ServiceID *string `header:"style=simple,explode=false,name=x-apideck-service-id"`
	// Include raw response. Mostly used for debugging purposes
	Raw *bool `default:"false" queryParam:"style=form,explode=true,name=raw"`
	// ID of the employee you are acting upon.
	EmployeeID     string                         `pathParam:"style=simple,explode=false,name=employee_id"`
	TimeOffRequest components.TimeOffRequestInput `request:"mediaType=application/json"`
}

func (h HrisTimeOffRequestsUpdateRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(h, "", false)
}

func (h *HrisTimeOffRequestsUpdateRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &h, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *HrisTimeOffRequestsUpdateRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *HrisTimeOffRequestsUpdateRequest) GetServiceID() *string {
	if o == nil {
		return nil
	}
	return o.ServiceID
}

func (o *HrisTimeOffRequestsUpdateRequest) GetRaw() *bool {
	if o == nil {
		return nil
	}
	return o.Raw
}

func (o *HrisTimeOffRequestsUpdateRequest) GetEmployeeID() string {
	if o == nil {
		return ""
	}
	return o.EmployeeID
}

func (o *HrisTimeOffRequestsUpdateRequest) GetTimeOffRequest() components.TimeOffRequestInput {
	if o == nil {
		return components.TimeOffRequestInput{}
	}
	return o.TimeOffRequest
}

type HrisTimeOffRequestsUpdateResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// TimeOffRequests
	UpdateTimeOffRequestResponse *components.UpdateTimeOffRequestResponse
	// Unexpected error
	UnexpectedErrorResponse *components.UnexpectedErrorResponse
}

func (o *HrisTimeOffRequestsUpdateResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *HrisTimeOffRequestsUpdateResponse) GetUpdateTimeOffRequestResponse() *components.UpdateTimeOffRequestResponse {
	if o == nil {
		return nil
	}
	return o.UpdateTimeOffRequestResponse
}

func (o *HrisTimeOffRequestsUpdateResponse) GetUnexpectedErrorResponse() *components.UnexpectedErrorResponse {
	if o == nil {
		return nil
	}
	return o.UnexpectedErrorResponse
}
