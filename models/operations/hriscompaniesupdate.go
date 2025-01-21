// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/apideck-libraries/sdk-go/internal/utils"
	"github.com/apideck-libraries/sdk-go/models/components"
)

type HrisCompaniesUpdateGlobals struct {
	// ID of the consumer which you want to get or push data from
	ConsumerID *string `header:"style=simple,explode=false,name=x-apideck-consumer-id"`
	// The ID of your Unify application
	AppID *string `header:"style=simple,explode=false,name=x-apideck-app-id"`
}

func (o *HrisCompaniesUpdateGlobals) GetConsumerID() *string {
	if o == nil {
		return nil
	}
	return o.ConsumerID
}

func (o *HrisCompaniesUpdateGlobals) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

type HrisCompaniesUpdateRequest struct {
	// ID of the record you are acting upon.
	ID string `pathParam:"style=simple,explode=false,name=id"`
	// ID of the consumer which you want to get or push data from
	ConsumerID *string `header:"style=simple,explode=false,name=x-apideck-consumer-id"`
	// The ID of your Unify application
	AppID *string `header:"style=simple,explode=false,name=x-apideck-app-id"`
	// Provide the service id you want to call (e.g., pipedrive). Only needed when a consumer has activated multiple integrations for a Unified API.
	ServiceID *string `header:"style=simple,explode=false,name=x-apideck-service-id"`
	// Include raw response. Mostly used for debugging purposes
	Raw         *bool                       `default:"false" queryParam:"style=form,explode=true,name=raw"`
	HrisCompany components.HrisCompanyInput `request:"mediaType=application/json"`
}

func (h HrisCompaniesUpdateRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(h, "", false)
}

func (h *HrisCompaniesUpdateRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &h, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *HrisCompaniesUpdateRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *HrisCompaniesUpdateRequest) GetConsumerID() *string {
	if o == nil {
		return nil
	}
	return o.ConsumerID
}

func (o *HrisCompaniesUpdateRequest) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

func (o *HrisCompaniesUpdateRequest) GetServiceID() *string {
	if o == nil {
		return nil
	}
	return o.ServiceID
}

func (o *HrisCompaniesUpdateRequest) GetRaw() *bool {
	if o == nil {
		return nil
	}
	return o.Raw
}

func (o *HrisCompaniesUpdateRequest) GetHrisCompany() components.HrisCompanyInput {
	if o == nil {
		return components.HrisCompanyInput{}
	}
	return o.HrisCompany
}

type HrisCompaniesUpdateResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Companies
	UpdateHrisCompanyResponse *components.UpdateHrisCompanyResponse
	// Unexpected error
	UnexpectedErrorResponse *components.UnexpectedErrorResponse
}

func (o *HrisCompaniesUpdateResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *HrisCompaniesUpdateResponse) GetUpdateHrisCompanyResponse() *components.UpdateHrisCompanyResponse {
	if o == nil {
		return nil
	}
	return o.UpdateHrisCompanyResponse
}

func (o *HrisCompaniesUpdateResponse) GetUnexpectedErrorResponse() *components.UnexpectedErrorResponse {
	if o == nil {
		return nil
	}
	return o.UnexpectedErrorResponse
}
