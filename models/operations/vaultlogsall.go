// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/apideck-libraries/sdk-go/internal/utils"
	"github.com/apideck-libraries/sdk-go/models/components"
)

type VaultLogsAllGlobals struct {
	// The ID of your Unify application
	AppID *string `header:"style=simple,explode=false,name=x-apideck-app-id"`
	// ID of the consumer which you want to get or push data from
	ConsumerID *string `header:"style=simple,explode=false,name=x-apideck-consumer-id"`
}

func (o *VaultLogsAllGlobals) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

func (o *VaultLogsAllGlobals) GetConsumerID() *string {
	if o == nil {
		return nil
	}
	return o.ConsumerID
}

type VaultLogsAllRequest struct {
	// Filter results
	Filter *components.LogsFilter `queryParam:"style=deepObject,explode=true,name=filter"`
	// Cursor to start from. You can find cursors for next/previous pages in the meta.cursors property of the response.
	Cursor *string `queryParam:"style=form,explode=true,name=cursor"`
	// Number of results to return. Minimum 1, Maximum 200, Default 20
	Limit *int64 `default:"20" queryParam:"style=form,explode=true,name=limit"`
}

func (v VaultLogsAllRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(v, "", false)
}

func (v *VaultLogsAllRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &v, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *VaultLogsAllRequest) GetFilter() *components.LogsFilter {
	if o == nil {
		return nil
	}
	return o.Filter
}

func (o *VaultLogsAllRequest) GetCursor() *string {
	if o == nil {
		return nil
	}
	return o.Cursor
}

func (o *VaultLogsAllRequest) GetLimit() *int64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

type VaultLogsAllResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Logs
	GetLogsResponse *components.GetLogsResponse
	// Unexpected error
	UnexpectedErrorResponse *components.UnexpectedErrorResponse

	Next func() (*VaultLogsAllResponse, error)
}

func (o *VaultLogsAllResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *VaultLogsAllResponse) GetGetLogsResponse() *components.GetLogsResponse {
	if o == nil {
		return nil
	}
	return o.GetLogsResponse
}

func (o *VaultLogsAllResponse) GetUnexpectedErrorResponse() *components.UnexpectedErrorResponse {
	if o == nil {
		return nil
	}
	return o.UnexpectedErrorResponse
}
