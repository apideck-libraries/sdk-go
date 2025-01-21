// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/apideck-libraries/sdk-go/internal/utils"
	"github.com/apideck-libraries/sdk-go/models/components"
)

type AccountingDepartmentsAllGlobals struct {
	// ID of the consumer which you want to get or push data from
	ConsumerID *string `header:"style=simple,explode=false,name=x-apideck-consumer-id"`
	// The ID of your Unify application
	AppID *string `header:"style=simple,explode=false,name=x-apideck-app-id"`
}

func (o *AccountingDepartmentsAllGlobals) GetConsumerID() *string {
	if o == nil {
		return nil
	}
	return o.ConsumerID
}

func (o *AccountingDepartmentsAllGlobals) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

type AccountingDepartmentsAllRequest struct {
	// Include raw response. Mostly used for debugging purposes
	Raw *bool `default:"false" queryParam:"style=form,explode=true,name=raw"`
	// ID of the consumer which you want to get or push data from
	ConsumerID *string `header:"style=simple,explode=false,name=x-apideck-consumer-id"`
	// The ID of your Unify application
	AppID *string `header:"style=simple,explode=false,name=x-apideck-app-id"`
	// Provide the service id you want to call (e.g., pipedrive). Only needed when a consumer has activated multiple integrations for a Unified API.
	ServiceID *string `header:"style=simple,explode=false,name=x-apideck-service-id"`
	// Cursor to start from. You can find cursors for next/previous pages in the meta.cursors property of the response.
	Cursor *string `queryParam:"style=form,explode=true,name=cursor"`
	// Number of results to return. Minimum 1, Maximum 200, Default 20
	Limit *int64 `default:"20" queryParam:"style=form,explode=true,name=limit"`
	// The 'fields' parameter allows API users to specify the fields they want to include in the API response. If this parameter is not present, the API will return all available fields. If this parameter is present, only the fields specified in the comma-separated string will be included in the response. Nested properties can also be requested by using a dot notation. <br /><br />Example: `fields=name,email,addresses.city`<br /><br />In the example above, the response will only include the fields "name", "email" and "addresses.city". If any other fields are available, they will be excluded.
	Fields *string `queryParam:"style=form,explode=true,name=fields"`
	// Apply filters
	Filter *components.AccountingDepartmentsFilter `queryParam:"style=deepObject,explode=true,name=filter"`
}

func (a AccountingDepartmentsAllRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AccountingDepartmentsAllRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *AccountingDepartmentsAllRequest) GetRaw() *bool {
	if o == nil {
		return nil
	}
	return o.Raw
}

func (o *AccountingDepartmentsAllRequest) GetConsumerID() *string {
	if o == nil {
		return nil
	}
	return o.ConsumerID
}

func (o *AccountingDepartmentsAllRequest) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

func (o *AccountingDepartmentsAllRequest) GetServiceID() *string {
	if o == nil {
		return nil
	}
	return o.ServiceID
}

func (o *AccountingDepartmentsAllRequest) GetCursor() *string {
	if o == nil {
		return nil
	}
	return o.Cursor
}

func (o *AccountingDepartmentsAllRequest) GetLimit() *int64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *AccountingDepartmentsAllRequest) GetFields() *string {
	if o == nil {
		return nil
	}
	return o.Fields
}

func (o *AccountingDepartmentsAllRequest) GetFilter() *components.AccountingDepartmentsFilter {
	if o == nil {
		return nil
	}
	return o.Filter
}

type AccountingDepartmentsAllResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Departments
	GetAccountingDepartmentsResponse *components.GetAccountingDepartmentsResponse
	// Unexpected error
	UnexpectedErrorResponse *components.UnexpectedErrorResponse

	Next func() (*AccountingDepartmentsAllResponse, error)
}

func (o *AccountingDepartmentsAllResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *AccountingDepartmentsAllResponse) GetGetAccountingDepartmentsResponse() *components.GetAccountingDepartmentsResponse {
	if o == nil {
		return nil
	}
	return o.GetAccountingDepartmentsResponse
}

func (o *AccountingDepartmentsAllResponse) GetUnexpectedErrorResponse() *components.UnexpectedErrorResponse {
	if o == nil {
		return nil
	}
	return o.UnexpectedErrorResponse
}
