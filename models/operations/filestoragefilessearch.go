// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/apideck-libraries/sdk-go/internal/utils"
	"github.com/apideck-libraries/sdk-go/models/components"
)

type FileStorageFilesSearchGlobals struct {
	// ID of the consumer which you want to get or push data from
	ConsumerID *string `header:"style=simple,explode=false,name=x-apideck-consumer-id"`
	// The ID of your Unify application
	AppID *string `header:"style=simple,explode=false,name=x-apideck-app-id"`
}

func (o *FileStorageFilesSearchGlobals) GetConsumerID() *string {
	if o == nil {
		return nil
	}
	return o.ConsumerID
}

func (o *FileStorageFilesSearchGlobals) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

type FileStorageFilesSearchRequest struct {
	// ID of the consumer which you want to get or push data from
	ConsumerID *string `header:"style=simple,explode=false,name=x-apideck-consumer-id"`
	// The ID of your Unify application
	AppID *string `header:"style=simple,explode=false,name=x-apideck-app-id"`
	// Provide the service id you want to call (e.g., pipedrive). Only needed when a consumer has activated multiple integrations for a Unified API.
	ServiceID *string `header:"style=simple,explode=false,name=x-apideck-service-id"`
	// Optional unmapped key/values that will be passed through to downstream as query parameters. Ie: ?pass_through[search]=leads becomes ?search=leads
	PassThrough map[string]any `queryParam:"style=deepObject,explode=true,name=pass_through"`
	// The 'fields' parameter allows API users to specify the fields they want to include in the API response. If this parameter is not present, the API will return all available fields. If this parameter is present, only the fields specified in the comma-separated string will be included in the response. Nested properties can also be requested by using a dot notation. <br /><br />Example: `fields=name,email,addresses.city`<br /><br />In the example above, the response will only include the fields "name", "email" and "addresses.city". If any other fields are available, they will be excluded.
	Fields *string `queryParam:"style=form,explode=true,name=fields"`
	// Cursor to start from. You can find cursors for next/previous pages in the meta.cursors property of the response.
	Cursor *string `queryParam:"style=form,explode=true,name=cursor"`
	// Number of results to return. Minimum 1, Maximum 200, Default 20
	Limit *int64 `default:"20" queryParam:"style=form,explode=true,name=limit"`
	// Apply filters
	Filter      *components.FilesFilter `queryParam:"style=deepObject,explode=true,name=filter"`
	FilesSearch components.FilesSearch  `request:"mediaType=application/json"`
}

func (f FileStorageFilesSearchRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(f, "", false)
}

func (f *FileStorageFilesSearchRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &f, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *FileStorageFilesSearchRequest) GetConsumerID() *string {
	if o == nil {
		return nil
	}
	return o.ConsumerID
}

func (o *FileStorageFilesSearchRequest) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

func (o *FileStorageFilesSearchRequest) GetServiceID() *string {
	if o == nil {
		return nil
	}
	return o.ServiceID
}

func (o *FileStorageFilesSearchRequest) GetPassThrough() map[string]any {
	if o == nil {
		return nil
	}
	return o.PassThrough
}

func (o *FileStorageFilesSearchRequest) GetFields() *string {
	if o == nil {
		return nil
	}
	return o.Fields
}

func (o *FileStorageFilesSearchRequest) GetCursor() *string {
	if o == nil {
		return nil
	}
	return o.Cursor
}

func (o *FileStorageFilesSearchRequest) GetLimit() *int64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *FileStorageFilesSearchRequest) GetFilter() *components.FilesFilter {
	if o == nil {
		return nil
	}
	return o.Filter
}

func (o *FileStorageFilesSearchRequest) GetFilesSearch() components.FilesSearch {
	if o == nil {
		return components.FilesSearch{}
	}
	return o.FilesSearch
}

type FileStorageFilesSearchResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Files
	GetFilesResponse *components.GetFilesResponse
	// Unexpected error
	UnexpectedErrorResponse *components.UnexpectedErrorResponse
}

func (o *FileStorageFilesSearchResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *FileStorageFilesSearchResponse) GetGetFilesResponse() *components.GetFilesResponse {
	if o == nil {
		return nil
	}
	return o.GetFilesResponse
}

func (o *FileStorageFilesSearchResponse) GetUnexpectedErrorResponse() *components.UnexpectedErrorResponse {
	if o == nil {
		return nil
	}
	return o.UnexpectedErrorResponse
}
