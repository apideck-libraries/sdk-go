// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/apideck-libraries/sdk-go/internal/utils"
	"github.com/apideck-libraries/sdk-go/models/components"
)

type ConnectorApisAllGlobals struct {
	// The ID of your Unify application
	AppID *string `header:"style=simple,explode=false,name=x-apideck-app-id"`
}

func (o *ConnectorApisAllGlobals) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

type ConnectorApisAllRequest struct {
	// The ID of your Unify application
	AppID *string `header:"style=simple,explode=false,name=x-apideck-app-id"`
	// Cursor to start from. You can find cursors for next/previous pages in the meta.cursors property of the response.
	Cursor *string `queryParam:"style=form,explode=true,name=cursor"`
	// Number of results to return. Minimum 1, Maximum 200, Default 20
	Limit *int64 `default:"20" queryParam:"style=form,explode=true,name=limit"`
	// Apply filters
	Filter *components.ApisFilter `queryParam:"style=deepObject,explode=true,name=filter"`
}

func (c ConnectorApisAllRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *ConnectorApisAllRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ConnectorApisAllRequest) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

func (o *ConnectorApisAllRequest) GetCursor() *string {
	if o == nil {
		return nil
	}
	return o.Cursor
}

func (o *ConnectorApisAllRequest) GetLimit() *int64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *ConnectorApisAllRequest) GetFilter() *components.ApisFilter {
	if o == nil {
		return nil
	}
	return o.Filter
}

type ConnectorApisAllResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Apis
	GetApisResponse *components.GetApisResponse
	// Unexpected error
	UnexpectedErrorResponse *components.UnexpectedErrorResponse

	Next func() (*ConnectorApisAllResponse, error)
}

func (o *ConnectorApisAllResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *ConnectorApisAllResponse) GetGetApisResponse() *components.GetApisResponse {
	if o == nil {
		return nil
	}
	return o.GetApisResponse
}

func (o *ConnectorApisAllResponse) GetUnexpectedErrorResponse() *components.UnexpectedErrorResponse {
	if o == nil {
		return nil
	}
	return o.UnexpectedErrorResponse
}
