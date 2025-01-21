// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

type ConnectorConnectorDocsOneGlobals struct {
	// The ID of your Unify application
	AppID *string `header:"style=simple,explode=false,name=x-apideck-app-id"`
}

func (o *ConnectorConnectorDocsOneGlobals) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

type ConnectorConnectorDocsOneRequest struct {
	// The ID of your Unify application
	AppID *string `header:"style=simple,explode=false,name=x-apideck-app-id"`
	// ID of the record you are acting upon.
	ID string `pathParam:"style=simple,explode=false,name=id"`
	// ID of the Doc
	DocID string `pathParam:"style=simple,explode=false,name=doc_id"`
}

func (o *ConnectorConnectorDocsOneRequest) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

func (o *ConnectorConnectorDocsOneRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *ConnectorConnectorDocsOneRequest) GetDocID() string {
	if o == nil {
		return ""
	}
	return o.DocID
}

type ConnectorConnectorDocsOneResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Connectors
	GetConnectorDocResponse *string
	// Unexpected error
	UnexpectedErrorResponse *components.UnexpectedErrorResponse
}

func (o *ConnectorConnectorDocsOneResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *ConnectorConnectorDocsOneResponse) GetGetConnectorDocResponse() *string {
	if o == nil {
		return nil
	}
	return o.GetConnectorDocResponse
}

func (o *ConnectorConnectorDocsOneResponse) GetUnexpectedErrorResponse() *components.UnexpectedErrorResponse {
	if o == nil {
		return nil
	}
	return o.UnexpectedErrorResponse
}
