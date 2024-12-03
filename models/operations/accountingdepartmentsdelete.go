// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/apideck-libraries/sdk-go/internal/utils"
	"github.com/apideck-libraries/sdk-go/models/components"
)

type AccountingDepartmentsDeleteGlobals struct {
	// ID of the consumer which you want to get or push data from
	ConsumerID *string `header:"style=simple,explode=false,name=x-apideck-consumer-id"`
	// The ID of your Unify application
	AppID *string `header:"style=simple,explode=false,name=x-apideck-app-id"`
}

func (o *AccountingDepartmentsDeleteGlobals) GetConsumerID() *string {
	if o == nil {
		return nil
	}
	return o.ConsumerID
}

func (o *AccountingDepartmentsDeleteGlobals) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

type AccountingDepartmentsDeleteRequest struct {
	// ID of the record you are acting upon.
	ID string `pathParam:"style=simple,explode=false,name=id"`
	// Provide the service id you want to call (e.g., pipedrive). Only needed when a consumer has activated multiple integrations for a Unified API.
	ServiceID *string `header:"style=simple,explode=false,name=x-apideck-service-id"`
	// Include raw response. Mostly used for debugging purposes
	Raw *bool `default:"false" queryParam:"style=form,explode=true,name=raw"`
}

func (a AccountingDepartmentsDeleteRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AccountingDepartmentsDeleteRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *AccountingDepartmentsDeleteRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *AccountingDepartmentsDeleteRequest) GetServiceID() *string {
	if o == nil {
		return nil
	}
	return o.ServiceID
}

func (o *AccountingDepartmentsDeleteRequest) GetRaw() *bool {
	if o == nil {
		return nil
	}
	return o.Raw
}

type AccountingDepartmentsDeleteResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Department deleted
	DeleteAccountingDepartmentResponse *components.DeleteAccountingDepartmentResponse
	// Unexpected error
	UnexpectedErrorResponse *components.UnexpectedErrorResponse
}

func (o *AccountingDepartmentsDeleteResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *AccountingDepartmentsDeleteResponse) GetDeleteAccountingDepartmentResponse() *components.DeleteAccountingDepartmentResponse {
	if o == nil {
		return nil
	}
	return o.DeleteAccountingDepartmentResponse
}

func (o *AccountingDepartmentsDeleteResponse) GetUnexpectedErrorResponse() *components.UnexpectedErrorResponse {
	if o == nil {
		return nil
	}
	return o.UnexpectedErrorResponse
}
