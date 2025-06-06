// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package sdkgo

import (
	"github.com/apideck-libraries/sdk-go/internal/config"
	"github.com/apideck-libraries/sdk-go/internal/hooks"
)

type Webhook struct {
	Webhooks  *Webhooks
	EventLogs *EventLogs

	rootSDK          *Apideck
	sdkConfiguration config.SDKConfiguration
	hooks            *hooks.Hooks
}

func newWebhook(rootSDK *Apideck, sdkConfig config.SDKConfiguration, hooks *hooks.Hooks) *Webhook {
	return &Webhook{
		rootSDK:          rootSDK,
		sdkConfiguration: sdkConfig,
		hooks:            hooks,
		Webhooks:         newWebhooks(rootSDK, sdkConfig, hooks),
		EventLogs:        newEventLogs(rootSDK, sdkConfig, hooks),
	}
}
