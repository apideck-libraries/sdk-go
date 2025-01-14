// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package sdkgo

import (
	"context"
	"fmt"
	"github.com/apideck-libraries/sdk-go/internal/globals"
	"github.com/apideck-libraries/sdk-go/internal/hooks"
	"github.com/apideck-libraries/sdk-go/internal/utils"
	"github.com/apideck-libraries/sdk-go/models/components"
	"github.com/apideck-libraries/sdk-go/retry"
	"net/http"
	"time"
)

// ServerList contains the list of servers available to the SDK
var ServerList = []string{
	"https://unify.apideck.com",
}

// HTTPClient provides an interface for suplying the SDK with a custom HTTP client
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// String provides a helper function to return a pointer to a string
func String(s string) *string { return &s }

// Bool provides a helper function to return a pointer to a bool
func Bool(b bool) *bool { return &b }

// Int provides a helper function to return a pointer to an int
func Int(i int) *int { return &i }

// Int64 provides a helper function to return a pointer to an int64
func Int64(i int64) *int64 { return &i }

// Float32 provides a helper function to return a pointer to a float32
func Float32(f float32) *float32 { return &f }

// Float64 provides a helper function to return a pointer to a float64
func Float64(f float64) *float64 { return &f }

// Pointer provides a helper function to return a pointer to a type
func Pointer[T any](v T) *T { return &v }

type sdkConfiguration struct {
	Client            HTTPClient
	Security          func(context.Context) (interface{}, error)
	ServerURL         string
	ServerIndex       int
	Language          string
	OpenAPIDocVersion string
	SDKVersion        string
	GenVersion        string
	UserAgent         string
	Globals           globals.Globals
	RetryConfig       *retry.Config
	Hooks             *hooks.Hooks
	Timeout           *time.Duration
}

func (c *sdkConfiguration) GetServerDetails() (string, map[string]string) {
	if c.ServerURL != "" {
		return c.ServerURL, nil
	}

	return ServerList[c.ServerIndex], nil
}

// Apideck - Apideck: The Apideck OpenAPI Spec: SDK Optimized
//
// https://developers.apideck.com - Apideck Developer Docs
type Apideck struct {
	Accounting    *Accounting
	Ats           *Ats
	Crm           *Crm
	Ecommerce     *Ecommerce
	FileStorage   *FileStorage
	Hris          *Hris
	Sms           *Sms
	IssueTracking *IssueTracking
	Connector     *Connector
	Vault         *Vault
	Webhook       *Webhook

	sdkConfiguration sdkConfiguration
}

type SDKOption func(*Apideck)

// WithServerURL allows the overriding of the default server URL
func WithServerURL(serverURL string) SDKOption {
	return func(sdk *Apideck) {
		sdk.sdkConfiguration.ServerURL = serverURL
	}
}

// WithTemplatedServerURL allows the overriding of the default server URL with a templated URL populated with the provided parameters
func WithTemplatedServerURL(serverURL string, params map[string]string) SDKOption {
	return func(sdk *Apideck) {
		if params != nil {
			serverURL = utils.ReplaceParameters(serverURL, params)
		}

		sdk.sdkConfiguration.ServerURL = serverURL
	}
}

// WithServerIndex allows the overriding of the default server by index
func WithServerIndex(serverIndex int) SDKOption {
	return func(sdk *Apideck) {
		if serverIndex < 0 || serverIndex >= len(ServerList) {
			panic(fmt.Errorf("server index %d out of range", serverIndex))
		}

		sdk.sdkConfiguration.ServerIndex = serverIndex
	}
}

// WithClient allows the overriding of the default HTTP client used by the SDK
func WithClient(client HTTPClient) SDKOption {
	return func(sdk *Apideck) {
		sdk.sdkConfiguration.Client = client
	}
}

// WithSecurity configures the SDK to use the provided security details
func WithSecurity(apiKey string) SDKOption {
	return func(sdk *Apideck) {
		security := components.Security{APIKey: &apiKey}
		sdk.sdkConfiguration.Security = utils.AsSecuritySource(&security)
	}
}

// WithSecuritySource configures the SDK to invoke the Security Source function on each method call to determine authentication
func WithSecuritySource(security func(context.Context) (components.Security, error)) SDKOption {
	return func(sdk *Apideck) {
		sdk.sdkConfiguration.Security = func(ctx context.Context) (interface{}, error) {
			return security(ctx)
		}
	}
}

// WithConsumerID allows setting the ConsumerID parameter for all supported operations
func WithConsumerID(consumerID string) SDKOption {
	return func(sdk *Apideck) {
		sdk.sdkConfiguration.Globals.ConsumerID = &consumerID
	}
}

// WithAppID allows setting the AppID parameter for all supported operations
func WithAppID(appID string) SDKOption {
	return func(sdk *Apideck) {
		sdk.sdkConfiguration.Globals.AppID = &appID
	}
}

func WithRetryConfig(retryConfig retry.Config) SDKOption {
	return func(sdk *Apideck) {
		sdk.sdkConfiguration.RetryConfig = &retryConfig
	}
}

// WithTimeout Optional request timeout applied to each operation
func WithTimeout(timeout time.Duration) SDKOption {
	return func(sdk *Apideck) {
		sdk.sdkConfiguration.Timeout = &timeout
	}
}
func (sdk *Apideck) fillGlobalsFromEnv() {
	if sdk.sdkConfiguration.Globals.ConsumerID == nil {
		if val := utils.ValueFromEnvVar("APIDECK_CONSUMER_ID", sdk.sdkConfiguration.Globals.ConsumerID); val != nil {
			if typedVal, ok := val.(string); ok {
				sdk.sdkConfiguration.Globals.ConsumerID = &typedVal
			}
		}
	}

	if sdk.sdkConfiguration.Globals.AppID == nil {
		if val := utils.ValueFromEnvVar("APIDECK_APP_ID", sdk.sdkConfiguration.Globals.AppID); val != nil {
			if typedVal, ok := val.(string); ok {
				sdk.sdkConfiguration.Globals.AppID = &typedVal
			}
		}
	}

}

// New creates a new instance of the SDK with the provided options
func New(opts ...SDKOption) *Apideck {
	sdk := &Apideck{
		sdkConfiguration: sdkConfiguration{
			Language:          "go",
			OpenAPIDocVersion: "10.9.1",
			SDKVersion:        "0.3.4",
			GenVersion:        "2.493.11",
			UserAgent:         "speakeasy-sdk/go 0.3.4 2.493.11 10.9.1 github.com/apideck-libraries/sdk-go",
			Globals:           globals.Globals{},
			Hooks:             hooks.New(),
		},
	}
	for _, opt := range opts {
		opt(sdk)
	}

	sdk.fillGlobalsFromEnv()

	if sdk.sdkConfiguration.Security == nil {
		var envVarSecurity components.Security
		if utils.PopulateSecurityFromEnv(&envVarSecurity) {
			sdk.sdkConfiguration.Security = utils.AsSecuritySource(envVarSecurity)
		}
	}

	// Use WithClient to override the default client if you would like to customize the timeout
	if sdk.sdkConfiguration.Client == nil {
		sdk.sdkConfiguration.Client = &http.Client{Timeout: 60 * time.Second}
	}

	currentServerURL, _ := sdk.sdkConfiguration.GetServerDetails()
	serverURL := currentServerURL
	serverURL, sdk.sdkConfiguration.Client = sdk.sdkConfiguration.Hooks.SDKInit(currentServerURL, sdk.sdkConfiguration.Client)
	if serverURL != currentServerURL {
		sdk.sdkConfiguration.ServerURL = serverURL
	}

	sdk.Accounting = newAccounting(sdk.sdkConfiguration)

	sdk.Ats = newAts(sdk.sdkConfiguration)

	sdk.Crm = newCrm(sdk.sdkConfiguration)

	sdk.Ecommerce = newEcommerce(sdk.sdkConfiguration)

	sdk.FileStorage = newFileStorage(sdk.sdkConfiguration)

	sdk.Hris = newHris(sdk.sdkConfiguration)

	sdk.Sms = newSms(sdk.sdkConfiguration)

	sdk.IssueTracking = newIssueTracking(sdk.sdkConfiguration)

	sdk.Connector = newConnector(sdk.sdkConfiguration)

	sdk.Vault = newVault(sdk.sdkConfiguration)

	sdk.Webhook = newWebhook(sdk.sdkConfiguration)

	return sdk
}
