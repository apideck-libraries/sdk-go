// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package config

import (
	"context"
	"github.com/apideck-libraries/sdk-go/internal/globals"
	"github.com/apideck-libraries/sdk-go/internal/utils"
	"github.com/apideck-libraries/sdk-go/retry"
	"net/http"
	"time"
)

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type SDKConfiguration struct {
	Client      HTTPClient
	Security    func(context.Context) (interface{}, error)
	ServerURL   string
	ServerIndex int
	ServerList  []string
	UserAgent   string
	Globals     globals.Globals
	RetryConfig *retry.Config
	Timeout     *time.Duration
}

func (c *SDKConfiguration) GetServerDetails() (string, map[string]string) {
	if c.ServerURL != "" {
		return c.ServerURL, nil
	}

	return c.ServerList[c.ServerIndex], nil
}

func (c *SDKConfiguration) FillGlobalsFromEnv() {
	if c.Globals.ConsumerID == nil {
		if val := utils.ValueFromEnvVar("APIDECK_CONSUMER_ID", c.Globals.ConsumerID); val != nil {
			if typedVal, ok := val.(string); ok {
				c.Globals.ConsumerID = &typedVal
			}
		}
	}
	if c.Globals.AppID == nil {
		if val := utils.ValueFromEnvVar("APIDECK_APP_ID", c.Globals.AppID); val != nil {
			if typedVal, ok := val.(string); ok {
				c.Globals.AppID = &typedVal
			}
		}
	}
}
