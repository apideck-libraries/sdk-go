// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/apideck-libraries/sdk-go/internal/utils"
	"github.com/apideck-libraries/sdk-go/models/components"
)

type AtsApplicantsUpdateGlobals struct {
	// ID of the consumer which you want to get or push data from
	ConsumerID *string `header:"style=simple,explode=false,name=x-apideck-consumer-id"`
	// The ID of your Unify application
	AppID *string `header:"style=simple,explode=false,name=x-apideck-app-id"`
}

func (o *AtsApplicantsUpdateGlobals) GetConsumerID() *string {
	if o == nil {
		return nil
	}
	return o.ConsumerID
}

func (o *AtsApplicantsUpdateGlobals) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

type AtsApplicantsUpdateRequest struct {
	// ID of the record you are acting upon.
	ID string `pathParam:"style=simple,explode=false,name=id"`
	// Provide the service id you want to call (e.g., pipedrive). Only needed when a consumer has activated multiple integrations for a Unified API.
	ServiceID *string `header:"style=simple,explode=false,name=x-apideck-service-id"`
	// Include raw response. Mostly used for debugging purposes
	Raw       *bool                     `default:"false" queryParam:"style=form,explode=true,name=raw"`
	Applicant components.ApplicantInput `request:"mediaType=application/json"`
}

func (a AtsApplicantsUpdateRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AtsApplicantsUpdateRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *AtsApplicantsUpdateRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *AtsApplicantsUpdateRequest) GetServiceID() *string {
	if o == nil {
		return nil
	}
	return o.ServiceID
}

func (o *AtsApplicantsUpdateRequest) GetRaw() *bool {
	if o == nil {
		return nil
	}
	return o.Raw
}

func (o *AtsApplicantsUpdateRequest) GetApplicant() components.ApplicantInput {
	if o == nil {
		return components.ApplicantInput{}
	}
	return o.Applicant
}

type AtsApplicantsUpdateResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Applicants
	UpdateApplicantResponse *components.UpdateApplicantResponse
	// Unexpected error
	UnexpectedErrorResponse *components.UnexpectedErrorResponse
}

func (o *AtsApplicantsUpdateResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *AtsApplicantsUpdateResponse) GetUpdateApplicantResponse() *components.UpdateApplicantResponse {
	if o == nil {
		return nil
	}
	return o.UpdateApplicantResponse
}

func (o *AtsApplicantsUpdateResponse) GetUnexpectedErrorResponse() *components.UnexpectedErrorResponse {
	if o == nil {
		return nil
	}
	return o.UnexpectedErrorResponse
}
