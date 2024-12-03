// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/apideck-libraries/sdk-go/internal/utils"
	"github.com/apideck-libraries/sdk-go/models/components"
)

type AccountingLedgerAccountsAllGlobals struct {
	// ID of the consumer which you want to get or push data from
	ConsumerID *string `header:"style=simple,explode=false,name=x-apideck-consumer-id"`
	// The ID of your Unify application
	AppID *string `header:"style=simple,explode=false,name=x-apideck-app-id"`
}

func (o *AccountingLedgerAccountsAllGlobals) GetConsumerID() *string {
	if o == nil {
		return nil
	}
	return o.ConsumerID
}

func (o *AccountingLedgerAccountsAllGlobals) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

type AccountingLedgerAccountsAllRequest struct {
	// Include raw response. Mostly used for debugging purposes
	Raw *bool `default:"false" queryParam:"style=form,explode=true,name=raw"`
	// Provide the service id you want to call (e.g., pipedrive). Only needed when a consumer has activated multiple integrations for a Unified API.
	ServiceID *string `header:"style=simple,explode=false,name=x-apideck-service-id"`
	// Cursor to start from. You can find cursors for next/previous pages in the meta.cursors property of the response.
	Cursor *string `queryParam:"style=form,explode=true,name=cursor"`
	// Number of results to return. Minimum 1, Maximum 200, Default 20
	Limit *int64 `default:"20" queryParam:"style=form,explode=true,name=limit"`
	// Apply filters
	Filter *components.LedgerAccountsFilter `queryParam:"style=deepObject,explode=true,name=filter"`
	// Apply sorting
	Sort *components.LedgerAccountsSort `queryParam:"style=deepObject,explode=true,name=sort"`
	// Optional unmapped key/values that will be passed through to downstream as query parameters. Ie: ?pass_through[search]=leads becomes ?search=leads
	PassThrough map[string]any `queryParam:"style=deepObject,explode=true,name=pass_through"`
	// The 'fields' parameter allows API users to specify the fields they want to include in the API response. If this parameter is not present, the API will return all available fields. If this parameter is present, only the fields specified in the comma-separated string will be included in the response. Nested properties can also be requested by using a dot notation. <br /><br />Example: `fields=name,email,addresses.city`<br /><br />In the example above, the response will only include the fields "name", "email" and "addresses.city". If any other fields are available, they will be excluded.
	Fields *string `queryParam:"style=form,explode=true,name=fields"`
}

func (a AccountingLedgerAccountsAllRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AccountingLedgerAccountsAllRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *AccountingLedgerAccountsAllRequest) GetRaw() *bool {
	if o == nil {
		return nil
	}
	return o.Raw
}

func (o *AccountingLedgerAccountsAllRequest) GetServiceID() *string {
	if o == nil {
		return nil
	}
	return o.ServiceID
}

func (o *AccountingLedgerAccountsAllRequest) GetCursor() *string {
	if o == nil {
		return nil
	}
	return o.Cursor
}

func (o *AccountingLedgerAccountsAllRequest) GetLimit() *int64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *AccountingLedgerAccountsAllRequest) GetFilter() *components.LedgerAccountsFilter {
	if o == nil {
		return nil
	}
	return o.Filter
}

func (o *AccountingLedgerAccountsAllRequest) GetSort() *components.LedgerAccountsSort {
	if o == nil {
		return nil
	}
	return o.Sort
}

func (o *AccountingLedgerAccountsAllRequest) GetPassThrough() map[string]any {
	if o == nil {
		return nil
	}
	return o.PassThrough
}

func (o *AccountingLedgerAccountsAllRequest) GetFields() *string {
	if o == nil {
		return nil
	}
	return o.Fields
}

type AccountingLedgerAccountsAllResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// LedgerAccounts
	GetLedgerAccountsResponse *components.GetLedgerAccountsResponse
	// Unexpected error
	UnexpectedErrorResponse *components.UnexpectedErrorResponse
}

func (o *AccountingLedgerAccountsAllResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *AccountingLedgerAccountsAllResponse) GetGetLedgerAccountsResponse() *components.GetLedgerAccountsResponse {
	if o == nil {
		return nil
	}
	return o.GetLedgerAccountsResponse
}

func (o *AccountingLedgerAccountsAllResponse) GetUnexpectedErrorResponse() *components.UnexpectedErrorResponse {
	if o == nil {
		return nil
	}
	return o.UnexpectedErrorResponse
}
