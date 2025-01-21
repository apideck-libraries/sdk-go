// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/apideck-libraries/sdk-go/internal/utils"
	"github.com/apideck-libraries/sdk-go/models/components"
)

var FileStorageUploadSessionsFinishServerList = []string{
	"https://upload.apideck.com",
}

type FileStorageUploadSessionsFinishGlobals struct {
	// ID of the consumer which you want to get or push data from
	ConsumerID *string `header:"style=simple,explode=false,name=x-apideck-consumer-id"`
	// The ID of your Unify application
	AppID *string `header:"style=simple,explode=false,name=x-apideck-app-id"`
}

func (o *FileStorageUploadSessionsFinishGlobals) GetConsumerID() *string {
	if o == nil {
		return nil
	}
	return o.ConsumerID
}

func (o *FileStorageUploadSessionsFinishGlobals) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

type FileStorageUploadSessionsFinishRequestBody struct {
}

type FileStorageUploadSessionsFinishRequest struct {
	// ID of the record you are acting upon.
	ID string `pathParam:"style=simple,explode=false,name=id"`
	// Include raw response. Mostly used for debugging purposes
	Raw *bool `default:"false" queryParam:"style=form,explode=true,name=raw"`
	// ID of the consumer which you want to get or push data from
	ConsumerID *string `header:"style=simple,explode=false,name=x-apideck-consumer-id"`
	// The ID of your Unify application
	AppID *string `header:"style=simple,explode=false,name=x-apideck-app-id"`
	// Provide the service id you want to call (e.g., pipedrive). Only needed when a consumer has activated multiple integrations for a Unified API.
	ServiceID *string `header:"style=simple,explode=false,name=x-apideck-service-id"`
	// The RFC3230 message digest of the uploaded part. Only required for the Box connector. More information on the Box API docs [here](https://developer.box.com/reference/put-files-upload-sessions-id/#param-digest)
	Digest      *string                                     `header:"style=simple,explode=false,name=digest"`
	RequestBody *FileStorageUploadSessionsFinishRequestBody `request:"mediaType=application/json"`
}

func (f FileStorageUploadSessionsFinishRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(f, "", false)
}

func (f *FileStorageUploadSessionsFinishRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &f, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *FileStorageUploadSessionsFinishRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *FileStorageUploadSessionsFinishRequest) GetRaw() *bool {
	if o == nil {
		return nil
	}
	return o.Raw
}

func (o *FileStorageUploadSessionsFinishRequest) GetConsumerID() *string {
	if o == nil {
		return nil
	}
	return o.ConsumerID
}

func (o *FileStorageUploadSessionsFinishRequest) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

func (o *FileStorageUploadSessionsFinishRequest) GetServiceID() *string {
	if o == nil {
		return nil
	}
	return o.ServiceID
}

func (o *FileStorageUploadSessionsFinishRequest) GetDigest() *string {
	if o == nil {
		return nil
	}
	return o.Digest
}

func (o *FileStorageUploadSessionsFinishRequest) GetRequestBody() *FileStorageUploadSessionsFinishRequestBody {
	if o == nil {
		return nil
	}
	return o.RequestBody
}

type FileStorageUploadSessionsFinishResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// File
	GetFileResponse *components.GetFileResponse
	// Unexpected error
	UnexpectedErrorResponse *components.UnexpectedErrorResponse
}

func (o *FileStorageUploadSessionsFinishResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *FileStorageUploadSessionsFinishResponse) GetGetFileResponse() *components.GetFileResponse {
	if o == nil {
		return nil
	}
	return o.GetFileResponse
}

func (o *FileStorageUploadSessionsFinishResponse) GetUnexpectedErrorResponse() *components.UnexpectedErrorResponse {
	if o == nil {
		return nil
	}
	return o.UnexpectedErrorResponse
}
