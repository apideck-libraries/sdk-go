// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

type VaultCustomMappingsAllGlobals struct {
	// ID of the consumer which you want to get or push data from
	ConsumerID *string `header:"style=simple,explode=false,name=x-apideck-consumer-id"`
	// The ID of your Unify application
	AppID *string `header:"style=simple,explode=false,name=x-apideck-app-id"`
}

func (o *VaultCustomMappingsAllGlobals) GetConsumerID() *string {
	if o == nil {
		return nil
	}
	return o.ConsumerID
}

func (o *VaultCustomMappingsAllGlobals) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

type VaultCustomMappingsAllRequest struct {
	// Unified API
	UnifiedAPI string `pathParam:"style=simple,explode=false,name=unified_api"`
	// Service ID of the resource to return
	ServiceID string `pathParam:"style=simple,explode=false,name=service_id"`
}

func (o *VaultCustomMappingsAllRequest) GetUnifiedAPI() string {
	if o == nil {
		return ""
	}
	return o.UnifiedAPI
}

func (o *VaultCustomMappingsAllRequest) GetServiceID() string {
	if o == nil {
		return ""
	}
	return o.ServiceID
}

type VaultCustomMappingsAllResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Custom mapping
	GetCustomMappingsResponse *components.GetCustomMappingsResponse
	// Unexpected error
	UnexpectedErrorResponse *components.UnexpectedErrorResponse
}

func (o *VaultCustomMappingsAllResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *VaultCustomMappingsAllResponse) GetGetCustomMappingsResponse() *components.GetCustomMappingsResponse {
	if o == nil {
		return nil
	}
	return o.GetCustomMappingsResponse
}

func (o *VaultCustomMappingsAllResponse) GetUnexpectedErrorResponse() *components.UnexpectedErrorResponse {
	if o == nil {
		return nil
	}
	return o.UnexpectedErrorResponse
}
