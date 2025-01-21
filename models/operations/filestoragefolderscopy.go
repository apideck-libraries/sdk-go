// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/apideck-libraries/sdk-go/internal/utils"
	"github.com/apideck-libraries/sdk-go/models/components"
)

type FileStorageFoldersCopyGlobals struct {
	// ID of the consumer which you want to get or push data from
	ConsumerID *string `header:"style=simple,explode=false,name=x-apideck-consumer-id"`
	// The ID of your Unify application
	AppID *string `header:"style=simple,explode=false,name=x-apideck-app-id"`
}

func (o *FileStorageFoldersCopyGlobals) GetConsumerID() *string {
	if o == nil {
		return nil
	}
	return o.ConsumerID
}

func (o *FileStorageFoldersCopyGlobals) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

type FileStorageFoldersCopyRequest struct {
	// ID of the record you are acting upon.
	ID string `pathParam:"style=simple,explode=false,name=id"`
	// ID of the consumer which you want to get or push data from
	ConsumerID *string `header:"style=simple,explode=false,name=x-apideck-consumer-id"`
	// The ID of your Unify application
	AppID *string `header:"style=simple,explode=false,name=x-apideck-app-id"`
	// Provide the service id you want to call (e.g., pipedrive). Only needed when a consumer has activated multiple integrations for a Unified API.
	ServiceID *string `header:"style=simple,explode=false,name=x-apideck-service-id"`
	// Include raw response. Mostly used for debugging purposes
	Raw *bool `default:"false" queryParam:"style=form,explode=true,name=raw"`
	// The 'fields' parameter allows API users to specify the fields they want to include in the API response. If this parameter is not present, the API will return all available fields. If this parameter is present, only the fields specified in the comma-separated string will be included in the response. Nested properties can also be requested by using a dot notation. <br /><br />Example: `fields=name,email,addresses.city`<br /><br />In the example above, the response will only include the fields "name", "email" and "addresses.city". If any other fields are available, they will be excluded.
	Fields            *string                      `queryParam:"style=form,explode=true,name=fields"`
	CopyFolderRequest components.CopyFolderRequest `request:"mediaType=application/json"`
}

func (f FileStorageFoldersCopyRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(f, "", false)
}

func (f *FileStorageFoldersCopyRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &f, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *FileStorageFoldersCopyRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *FileStorageFoldersCopyRequest) GetConsumerID() *string {
	if o == nil {
		return nil
	}
	return o.ConsumerID
}

func (o *FileStorageFoldersCopyRequest) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

func (o *FileStorageFoldersCopyRequest) GetServiceID() *string {
	if o == nil {
		return nil
	}
	return o.ServiceID
}

func (o *FileStorageFoldersCopyRequest) GetRaw() *bool {
	if o == nil {
		return nil
	}
	return o.Raw
}

func (o *FileStorageFoldersCopyRequest) GetFields() *string {
	if o == nil {
		return nil
	}
	return o.Fields
}

func (o *FileStorageFoldersCopyRequest) GetCopyFolderRequest() components.CopyFolderRequest {
	if o == nil {
		return components.CopyFolderRequest{}
	}
	return o.CopyFolderRequest
}

type FileStorageFoldersCopyResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Folders
	UpdateFolderResponse *components.UpdateFolderResponse
	// Unexpected error
	UnexpectedErrorResponse *components.UnexpectedErrorResponse
}

func (o *FileStorageFoldersCopyResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *FileStorageFoldersCopyResponse) GetUpdateFolderResponse() *components.UpdateFolderResponse {
	if o == nil {
		return nil
	}
	return o.UpdateFolderResponse
}

func (o *FileStorageFoldersCopyResponse) GetUnexpectedErrorResponse() *components.UnexpectedErrorResponse {
	if o == nil {
		return nil
	}
	return o.UnexpectedErrorResponse
}
