// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/apideck-libraries/sdk-go/internal/utils"
	"github.com/apideck-libraries/sdk-go/models/components"
)

type AccountingBillPaymentsAddGlobals struct {
	// ID of the consumer which you want to get or push data from
	ConsumerID *string `header:"style=simple,explode=false,name=x-apideck-consumer-id"`
	// The ID of your Unify application
	AppID *string `header:"style=simple,explode=false,name=x-apideck-app-id"`
}

func (o *AccountingBillPaymentsAddGlobals) GetConsumerID() *string {
	if o == nil {
		return nil
	}
	return o.ConsumerID
}

func (o *AccountingBillPaymentsAddGlobals) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

type AccountingBillPaymentsAddRequest struct {
	// Include raw response. Mostly used for debugging purposes
	Raw *bool `default:"false" queryParam:"style=form,explode=true,name=raw"`
	// Provide the service id you want to call (e.g., pipedrive). Only needed when a consumer has activated multiple integrations for a Unified API.
	ServiceID   *string                     `header:"style=simple,explode=false,name=x-apideck-service-id"`
	BillPayment components.BillPaymentInput `request:"mediaType=application/json"`
}

func (a AccountingBillPaymentsAddRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AccountingBillPaymentsAddRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *AccountingBillPaymentsAddRequest) GetRaw() *bool {
	if o == nil {
		return nil
	}
	return o.Raw
}

func (o *AccountingBillPaymentsAddRequest) GetServiceID() *string {
	if o == nil {
		return nil
	}
	return o.ServiceID
}

func (o *AccountingBillPaymentsAddRequest) GetBillPayment() components.BillPaymentInput {
	if o == nil {
		return components.BillPaymentInput{}
	}
	return o.BillPayment
}

type AccountingBillPaymentsAddResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Bill Payment created
	CreateBillPaymentResponse *components.CreateBillPaymentResponse
	// Unexpected error
	UnexpectedErrorResponse *components.UnexpectedErrorResponse
}

func (o *AccountingBillPaymentsAddResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *AccountingBillPaymentsAddResponse) GetCreateBillPaymentResponse() *components.CreateBillPaymentResponse {
	if o == nil {
		return nil
	}
	return o.CreateBillPaymentResponse
}

func (o *AccountingBillPaymentsAddResponse) GetUnexpectedErrorResponse() *components.UnexpectedErrorResponse {
	if o == nil {
		return nil
	}
	return o.UnexpectedErrorResponse
}
