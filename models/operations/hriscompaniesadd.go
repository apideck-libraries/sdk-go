// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/apideck-libraries/sdk-go/internal/utils"
	"github.com/apideck-libraries/sdk-go/models/components"
)

type HrisCompaniesAddGlobals struct {
	// ID of the consumer which you want to get or push data from
	ConsumerID *string `header:"style=simple,explode=false,name=x-apideck-consumer-id"`
	// The ID of your Unify application
	AppID *string `header:"style=simple,explode=false,name=x-apideck-app-id"`
}

func (o *HrisCompaniesAddGlobals) GetConsumerID() *string {
	if o == nil {
		return nil
	}
	return o.ConsumerID
}

func (o *HrisCompaniesAddGlobals) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

type HrisCompaniesAddRequest struct {
	// Include raw response. Mostly used for debugging purposes
	Raw *bool `default:"false" queryParam:"style=form,explode=true,name=raw"`
	// Provide the service id you want to call (e.g., pipedrive). Only needed when a consumer has activated multiple integrations for a Unified API.
	ServiceID   *string                     `header:"style=simple,explode=false,name=x-apideck-service-id"`
	HrisCompany components.HrisCompanyInput `request:"mediaType=application/json"`
}

func (h HrisCompaniesAddRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(h, "", false)
}

func (h *HrisCompaniesAddRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &h, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *HrisCompaniesAddRequest) GetRaw() *bool {
	if o == nil {
		return nil
	}
	return o.Raw
}

func (o *HrisCompaniesAddRequest) GetServiceID() *string {
	if o == nil {
		return nil
	}
	return o.ServiceID
}

func (o *HrisCompaniesAddRequest) GetHrisCompany() components.HrisCompanyInput {
	if o == nil {
		return components.HrisCompanyInput{}
	}
	return o.HrisCompany
}

type HrisCompaniesAddResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Companies
	CreateHrisCompanyResponse *components.CreateHrisCompanyResponse
	// Unexpected error
	UnexpectedErrorResponse *components.UnexpectedErrorResponse
}

func (o *HrisCompaniesAddResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *HrisCompaniesAddResponse) GetCreateHrisCompanyResponse() *components.CreateHrisCompanyResponse {
	if o == nil {
		return nil
	}
	return o.CreateHrisCompanyResponse
}

func (o *HrisCompaniesAddResponse) GetUnexpectedErrorResponse() *components.UnexpectedErrorResponse {
	if o == nil {
		return nil
	}
	return o.UnexpectedErrorResponse
}
