// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

type VaultValidateConnectionStateGlobals struct {
	// ID of the consumer which you want to get or push data from
	ConsumerID *string `header:"style=simple,explode=false,name=x-apideck-consumer-id"`
	// The ID of your Unify application
	AppID *string `header:"style=simple,explode=false,name=x-apideck-app-id"`
}

func (o *VaultValidateConnectionStateGlobals) GetConsumerID() *string {
	if o == nil {
		return nil
	}
	return o.ConsumerID
}

func (o *VaultValidateConnectionStateGlobals) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

type VaultValidateConnectionStateRequestBody struct {
}

type VaultValidateConnectionStateRequest struct {
	// ID of the consumer which you want to get or push data from
	ConsumerID *string `header:"style=simple,explode=false,name=x-apideck-consumer-id"`
	// The ID of your Unify application
	AppID *string `header:"style=simple,explode=false,name=x-apideck-app-id"`
	// Service ID of the resource to return
	ServiceID string `pathParam:"style=simple,explode=false,name=service_id"`
	// Unified API
	UnifiedAPI  string                                   `pathParam:"style=simple,explode=false,name=unified_api"`
	RequestBody *VaultValidateConnectionStateRequestBody `request:"mediaType=application/json"`
}

func (o *VaultValidateConnectionStateRequest) GetConsumerID() *string {
	if o == nil {
		return nil
	}
	return o.ConsumerID
}

func (o *VaultValidateConnectionStateRequest) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

func (o *VaultValidateConnectionStateRequest) GetServiceID() string {
	if o == nil {
		return ""
	}
	return o.ServiceID
}

func (o *VaultValidateConnectionStateRequest) GetUnifiedAPI() string {
	if o == nil {
		return ""
	}
	return o.UnifiedAPI
}

func (o *VaultValidateConnectionStateRequest) GetRequestBody() *VaultValidateConnectionStateRequestBody {
	if o == nil {
		return nil
	}
	return o.RequestBody
}

type VaultValidateConnectionStateResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Connection access token refreshed
	ValidateConnectionStateResponse *components.ValidateConnectionStateResponse
	// Unexpected error
	UnexpectedErrorResponse *components.UnexpectedErrorResponse
}

func (o *VaultValidateConnectionStateResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *VaultValidateConnectionStateResponse) GetValidateConnectionStateResponse() *components.ValidateConnectionStateResponse {
	if o == nil {
		return nil
	}
	return o.ValidateConnectionStateResponse
}

func (o *VaultValidateConnectionStateResponse) GetUnexpectedErrorResponse() *components.UnexpectedErrorResponse {
	if o == nil {
		return nil
	}
	return o.UnexpectedErrorResponse
}
