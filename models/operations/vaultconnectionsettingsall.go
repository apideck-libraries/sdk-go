// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

type VaultConnectionSettingsAllGlobals struct {
	// ID of the consumer which you want to get or push data from
	ConsumerID *string `header:"style=simple,explode=false,name=x-apideck-consumer-id"`
	// The ID of your Unify application
	AppID *string `header:"style=simple,explode=false,name=x-apideck-app-id"`
}

func (o *VaultConnectionSettingsAllGlobals) GetConsumerID() *string {
	if o == nil {
		return nil
	}
	return o.ConsumerID
}

func (o *VaultConnectionSettingsAllGlobals) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

type VaultConnectionSettingsAllRequest struct {
	// Unified API
	UnifiedAPI string `pathParam:"style=simple,explode=false,name=unified_api"`
	// Service ID of the resource to return
	ServiceID string `pathParam:"style=simple,explode=false,name=service_id"`
	// Name of the resource (plural)
	Resource string `pathParam:"style=simple,explode=false,name=resource"`
}

func (o *VaultConnectionSettingsAllRequest) GetUnifiedAPI() string {
	if o == nil {
		return ""
	}
	return o.UnifiedAPI
}

func (o *VaultConnectionSettingsAllRequest) GetServiceID() string {
	if o == nil {
		return ""
	}
	return o.ServiceID
}

func (o *VaultConnectionSettingsAllRequest) GetResource() string {
	if o == nil {
		return ""
	}
	return o.Resource
}

type VaultConnectionSettingsAllResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Connection
	GetConnectionResponse *components.GetConnectionResponse
	// Unexpected error
	UnexpectedErrorResponse *components.UnexpectedErrorResponse
}

func (o *VaultConnectionSettingsAllResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *VaultConnectionSettingsAllResponse) GetGetConnectionResponse() *components.GetConnectionResponse {
	if o == nil {
		return nil
	}
	return o.GetConnectionResponse
}

func (o *VaultConnectionSettingsAllResponse) GetUnexpectedErrorResponse() *components.UnexpectedErrorResponse {
	if o == nil {
		return nil
	}
	return o.UnexpectedErrorResponse
}
