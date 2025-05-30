// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/apideck-libraries/sdk-go/internal/utils"
	"github.com/apideck-libraries/sdk-go/models/components"
)

type CrmPipelinesAddGlobals struct {
	// ID of the consumer which you want to get or push data from
	ConsumerID *string `header:"style=simple,explode=false,name=x-apideck-consumer-id"`
	// The ID of your Unify application
	AppID *string `header:"style=simple,explode=false,name=x-apideck-app-id"`
}

func (o *CrmPipelinesAddGlobals) GetConsumerID() *string {
	if o == nil {
		return nil
	}
	return o.ConsumerID
}

func (o *CrmPipelinesAddGlobals) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

type CrmPipelinesAddRequest struct {
	// Include raw response. Mostly used for debugging purposes
	Raw *bool `default:"false" queryParam:"style=form,explode=true,name=raw"`
	// ID of the consumer which you want to get or push data from
	ConsumerID *string `header:"style=simple,explode=false,name=x-apideck-consumer-id"`
	// The ID of your Unify application
	AppID *string `header:"style=simple,explode=false,name=x-apideck-app-id"`
	// Provide the service id you want to call (e.g., pipedrive). Only needed when a consumer has activated multiple integrations for a Unified API.
	ServiceID *string                  `header:"style=simple,explode=false,name=x-apideck-service-id"`
	Pipeline  components.PipelineInput `request:"mediaType=application/json"`
}

func (c CrmPipelinesAddRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CrmPipelinesAddRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CrmPipelinesAddRequest) GetRaw() *bool {
	if o == nil {
		return nil
	}
	return o.Raw
}

func (o *CrmPipelinesAddRequest) GetConsumerID() *string {
	if o == nil {
		return nil
	}
	return o.ConsumerID
}

func (o *CrmPipelinesAddRequest) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

func (o *CrmPipelinesAddRequest) GetServiceID() *string {
	if o == nil {
		return nil
	}
	return o.ServiceID
}

func (o *CrmPipelinesAddRequest) GetPipeline() components.PipelineInput {
	if o == nil {
		return components.PipelineInput{}
	}
	return o.Pipeline
}

type CrmPipelinesAddResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Pipeline created
	CreatePipelineResponse *components.CreatePipelineResponse
	// Unexpected error
	UnexpectedErrorResponse *components.UnexpectedErrorResponse
}

func (o *CrmPipelinesAddResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *CrmPipelinesAddResponse) GetCreatePipelineResponse() *components.CreatePipelineResponse {
	if o == nil {
		return nil
	}
	return o.CreatePipelineResponse
}

func (o *CrmPipelinesAddResponse) GetUnexpectedErrorResponse() *components.UnexpectedErrorResponse {
	if o == nil {
		return nil
	}
	return o.UnexpectedErrorResponse
}
