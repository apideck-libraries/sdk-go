// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/apideck-libraries/sdk-go/internal/utils"
)

// ConnectionStatus - Status of the connection.
type ConnectionStatus string

const (
	ConnectionStatusLive      ConnectionStatus = "live"
	ConnectionStatusUpcoming  ConnectionStatus = "upcoming"
	ConnectionStatusRequested ConnectionStatus = "requested"
)

func (e ConnectionStatus) ToPointer() *ConnectionStatus {
	return &e
}
func (e *ConnectionStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "live":
		fallthrough
	case "upcoming":
		fallthrough
	case "requested":
		*e = ConnectionStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ConnectionStatus: %v", v)
	}
}

type Target string

const (
	TargetCustomFields Target = "custom_fields"
	TargetResource     Target = "resource"
)

func (e Target) ToPointer() *Target {
	return &e
}
func (e *Target) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "custom_fields":
		fallthrough
	case "resource":
		*e = Target(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Target: %v", v)
	}
}

type ConnectionValue5Type string

const (
	ConnectionValue5TypeStr     ConnectionValue5Type = "str"
	ConnectionValue5TypeInteger ConnectionValue5Type = "integer"
	ConnectionValue5TypeNumber  ConnectionValue5Type = "number"
)

type ConnectionValue5 struct {
	Str     *string  `queryParam:"inline"`
	Integer *int64   `queryParam:"inline"`
	Number  *float64 `queryParam:"inline"`

	Type ConnectionValue5Type
}

func CreateConnectionValue5Str(str string) ConnectionValue5 {
	typ := ConnectionValue5TypeStr

	return ConnectionValue5{
		Str:  &str,
		Type: typ,
	}
}

func CreateConnectionValue5Integer(integer int64) ConnectionValue5 {
	typ := ConnectionValue5TypeInteger

	return ConnectionValue5{
		Integer: &integer,
		Type:    typ,
	}
}

func CreateConnectionValue5Number(number float64) ConnectionValue5 {
	typ := ConnectionValue5TypeNumber

	return ConnectionValue5{
		Number: &number,
		Type:   typ,
	}
}

func (u *ConnectionValue5) UnmarshalJSON(data []byte) error {

	var str string = ""
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = &str
		u.Type = ConnectionValue5TypeStr
		return nil
	}

	var integer int64 = int64(0)
	if err := utils.UnmarshalJSON(data, &integer, "", true, true); err == nil {
		u.Integer = &integer
		u.Type = ConnectionValue5TypeInteger
		return nil
	}

	var number float64 = float64(0)
	if err := utils.UnmarshalJSON(data, &number, "", true, true); err == nil {
		u.Number = &number
		u.Type = ConnectionValue5TypeNumber
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for ConnectionValue5", string(data))
}

func (u ConnectionValue5) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.Integer != nil {
		return utils.MarshalJSON(u.Integer, "", true)
	}

	if u.Number != nil {
		return utils.MarshalJSON(u.Number, "", true)
	}

	return nil, errors.New("could not marshal union type ConnectionValue5: all fields are null")
}

type ConnectionValueType string

const (
	ConnectionValueTypeStr                     ConnectionValueType = "str"
	ConnectionValueTypeInteger                 ConnectionValueType = "integer"
	ConnectionValueTypeNumber                  ConnectionValueType = "number"
	ConnectionValueTypeBoolean                 ConnectionValueType = "boolean"
	ConnectionValueTypeArrayOfConnectionValue5 ConnectionValueType = "arrayOfConnectionValue5"
)

type ConnectionValue struct {
	Str                     *string            `queryParam:"inline"`
	Integer                 *int64             `queryParam:"inline"`
	Number                  *float64           `queryParam:"inline"`
	Boolean                 *bool              `queryParam:"inline"`
	ArrayOfConnectionValue5 []ConnectionValue5 `queryParam:"inline"`

	Type ConnectionValueType
}

func CreateConnectionValueStr(str string) ConnectionValue {
	typ := ConnectionValueTypeStr

	return ConnectionValue{
		Str:  &str,
		Type: typ,
	}
}

func CreateConnectionValueInteger(integer int64) ConnectionValue {
	typ := ConnectionValueTypeInteger

	return ConnectionValue{
		Integer: &integer,
		Type:    typ,
	}
}

func CreateConnectionValueNumber(number float64) ConnectionValue {
	typ := ConnectionValueTypeNumber

	return ConnectionValue{
		Number: &number,
		Type:   typ,
	}
}

func CreateConnectionValueBoolean(boolean bool) ConnectionValue {
	typ := ConnectionValueTypeBoolean

	return ConnectionValue{
		Boolean: &boolean,
		Type:    typ,
	}
}

func CreateConnectionValueArrayOfConnectionValue5(arrayOfConnectionValue5 []ConnectionValue5) ConnectionValue {
	typ := ConnectionValueTypeArrayOfConnectionValue5

	return ConnectionValue{
		ArrayOfConnectionValue5: arrayOfConnectionValue5,
		Type:                    typ,
	}
}

func (u *ConnectionValue) UnmarshalJSON(data []byte) error {

	var str string = ""
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = &str
		u.Type = ConnectionValueTypeStr
		return nil
	}

	var integer int64 = int64(0)
	if err := utils.UnmarshalJSON(data, &integer, "", true, true); err == nil {
		u.Integer = &integer
		u.Type = ConnectionValueTypeInteger
		return nil
	}

	var number float64 = float64(0)
	if err := utils.UnmarshalJSON(data, &number, "", true, true); err == nil {
		u.Number = &number
		u.Type = ConnectionValueTypeNumber
		return nil
	}

	var boolean bool = false
	if err := utils.UnmarshalJSON(data, &boolean, "", true, true); err == nil {
		u.Boolean = &boolean
		u.Type = ConnectionValueTypeBoolean
		return nil
	}

	var arrayOfConnectionValue5 []ConnectionValue5 = []ConnectionValue5{}
	if err := utils.UnmarshalJSON(data, &arrayOfConnectionValue5, "", true, true); err == nil {
		u.ArrayOfConnectionValue5 = arrayOfConnectionValue5
		u.Type = ConnectionValueTypeArrayOfConnectionValue5
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for ConnectionValue", string(data))
}

func (u ConnectionValue) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.Integer != nil {
		return utils.MarshalJSON(u.Integer, "", true)
	}

	if u.Number != nil {
		return utils.MarshalJSON(u.Number, "", true)
	}

	if u.Boolean != nil {
		return utils.MarshalJSON(u.Boolean, "", true)
	}

	if u.ArrayOfConnectionValue5 != nil {
		return utils.MarshalJSON(u.ArrayOfConnectionValue5, "", true)
	}

	return nil, errors.New("could not marshal union type ConnectionValue: all fields are null")
}

type Defaults struct {
	Target  *Target           `json:"target,omitempty"`
	ID      *string           `json:"id,omitempty"`
	Options []FormFieldOption `json:"options,omitempty"`
	Value   *ConnectionValue  `json:"value,omitempty"`
}

func (o *Defaults) GetTarget() *Target {
	if o == nil {
		return nil
	}
	return o.Target
}

func (o *Defaults) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *Defaults) GetOptions() []FormFieldOption {
	if o == nil {
		return nil
	}
	return o.Options
}

func (o *Defaults) GetValue() *ConnectionValue {
	if o == nil {
		return nil
	}
	return o.Value
}

type Configuration struct {
	Resource *string    `json:"resource,omitempty"`
	Defaults []Defaults `json:"defaults,omitempty"`
}

func (o *Configuration) GetResource() *string {
	if o == nil {
		return nil
	}
	return o.Resource
}

func (o *Configuration) GetDefaults() []Defaults {
	if o == nil {
		return nil
	}
	return o.Defaults
}

type Connection struct {
	// The unique identifier of the connection.
	ID *string `json:"id,omitempty"`
	// The ID of the service this connection belongs to.
	ServiceID *string `json:"service_id,omitempty"`
	// The name of the connection
	Name    *string `json:"name,omitempty"`
	TagLine *string `json:"tag_line,omitempty"`
	// The unified API category where the connection belongs to.
	UnifiedAPI *string `json:"unified_api,omitempty"`
	// [Connection state flow](#section/Connection-state)
	State *ConnectionState `json:"state,omitempty"`
	// The current state of the Integration.
	IntegrationState *IntegrationState `json:"integration_state,omitempty"`
	// Type of authorization used by the connector
	AuthType *AuthType `json:"auth_type,omitempty"`
	// OAuth grant type used by the connector. More info: https://oauth.net/2/grant-types
	OauthGrantType *OAuthGrantType `json:"oauth_grant_type,omitempty"`
	// Status of the connection.
	Status *ConnectionStatus `json:"status,omitempty"`
	// Whether the connection is enabled or not. You can enable or disable a connection using the Update Connection API.
	Enabled *bool `json:"enabled,omitempty"`
	// The website URL of the connection
	Website *string `json:"website,omitempty"`
	// A visual icon of the connection, that will be shown in the Vault
	Icon *string `json:"icon,omitempty"`
	// The logo of the connection, that will be shown in the Vault
	Logo *string `json:"logo,omitempty"`
	// The OAuth redirect URI. Redirect your users to this URI to let them authorize your app in the connector's UI. Before you can use this URI, you must add `redirect_uri` as a query parameter to the `authorize_url`. Be sure to URL encode the `redirect_uri` part. Your users will be redirected to this `redirect_uri` after they granted access to your app in the connector's UI.
	AuthorizeURL *string `json:"authorize_url,omitempty"`
	// The OAuth revoke URI. Redirect your users to this URI to revoke this connection. Before you can use this URI, you must add `redirect_uri` as a query parameter. Your users will be redirected to this `redirect_uri` after they granted access to your app in the connector's UI.
	RevokeURL *string `json:"revoke_url,omitempty"`
	// Connection settings. Values will persist to `form_fields` with corresponding id
	Settings map[string]any `json:"settings,omitempty"`
	// Attach your own consumer specific metadata
	Metadata map[string]any `json:"metadata,omitempty"`
	// The settings that are wanted to create a connection.
	FormFields              []FormField     `json:"form_fields,omitempty"`
	Configuration           []Configuration `json:"configuration,omitempty"`
	ConfigurableResources   []string        `json:"configurable_resources,omitempty"`
	ResourceSchemaSupport   []string        `json:"resource_schema_support,omitempty"`
	ResourceSettingsSupport []string        `json:"resource_settings_support,omitempty"`
	ValidationSupport       *bool           `json:"validation_support,omitempty"`
	SchemaSupport           *bool           `json:"schema_support,omitempty"`
	// List of settings that are required to be configured on integration before authorization can occur
	SettingsRequiredForAuthorization []string              `json:"settings_required_for_authorization,omitempty"`
	Subscriptions                    []WebhookSubscription `json:"subscriptions,omitempty"`
	// Whether the connector has a guide available in the developer docs or not (https://docs.apideck.com/connectors/{service_id}/docs/consumer+connection).
	HasGuide  *bool    `json:"has_guide,omitempty"`
	CreatedAt *float64 `json:"created_at,omitempty"`
	// List of custom mappings configured for this connection
	CustomMappings []CustomMapping `json:"custom_mappings,omitempty"`
	UpdatedAt      *float64        `json:"updated_at,omitempty"`
}

func (o *Connection) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *Connection) GetServiceID() *string {
	if o == nil {
		return nil
	}
	return o.ServiceID
}

func (o *Connection) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *Connection) GetTagLine() *string {
	if o == nil {
		return nil
	}
	return o.TagLine
}

func (o *Connection) GetUnifiedAPI() *string {
	if o == nil {
		return nil
	}
	return o.UnifiedAPI
}

func (o *Connection) GetState() *ConnectionState {
	if o == nil {
		return nil
	}
	return o.State
}

func (o *Connection) GetIntegrationState() *IntegrationState {
	if o == nil {
		return nil
	}
	return o.IntegrationState
}

func (o *Connection) GetAuthType() *AuthType {
	if o == nil {
		return nil
	}
	return o.AuthType
}

func (o *Connection) GetOauthGrantType() *OAuthGrantType {
	if o == nil {
		return nil
	}
	return o.OauthGrantType
}

func (o *Connection) GetStatus() *ConnectionStatus {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *Connection) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *Connection) GetWebsite() *string {
	if o == nil {
		return nil
	}
	return o.Website
}

func (o *Connection) GetIcon() *string {
	if o == nil {
		return nil
	}
	return o.Icon
}

func (o *Connection) GetLogo() *string {
	if o == nil {
		return nil
	}
	return o.Logo
}

func (o *Connection) GetAuthorizeURL() *string {
	if o == nil {
		return nil
	}
	return o.AuthorizeURL
}

func (o *Connection) GetRevokeURL() *string {
	if o == nil {
		return nil
	}
	return o.RevokeURL
}

func (o *Connection) GetSettings() map[string]any {
	if o == nil {
		return nil
	}
	return o.Settings
}

func (o *Connection) GetMetadata() map[string]any {
	if o == nil {
		return nil
	}
	return o.Metadata
}

func (o *Connection) GetFormFields() []FormField {
	if o == nil {
		return nil
	}
	return o.FormFields
}

func (o *Connection) GetConfiguration() []Configuration {
	if o == nil {
		return nil
	}
	return o.Configuration
}

func (o *Connection) GetConfigurableResources() []string {
	if o == nil {
		return nil
	}
	return o.ConfigurableResources
}

func (o *Connection) GetResourceSchemaSupport() []string {
	if o == nil {
		return nil
	}
	return o.ResourceSchemaSupport
}

func (o *Connection) GetResourceSettingsSupport() []string {
	if o == nil {
		return nil
	}
	return o.ResourceSettingsSupport
}

func (o *Connection) GetValidationSupport() *bool {
	if o == nil {
		return nil
	}
	return o.ValidationSupport
}

func (o *Connection) GetSchemaSupport() *bool {
	if o == nil {
		return nil
	}
	return o.SchemaSupport
}

func (o *Connection) GetSettingsRequiredForAuthorization() []string {
	if o == nil {
		return nil
	}
	return o.SettingsRequiredForAuthorization
}

func (o *Connection) GetSubscriptions() []WebhookSubscription {
	if o == nil {
		return nil
	}
	return o.Subscriptions
}

func (o *Connection) GetHasGuide() *bool {
	if o == nil {
		return nil
	}
	return o.HasGuide
}

func (o *Connection) GetCreatedAt() *float64 {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *Connection) GetCustomMappings() []CustomMapping {
	if o == nil {
		return nil
	}
	return o.CustomMappings
}

func (o *Connection) GetUpdatedAt() *float64 {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

type ConnectionDefaults struct {
	ID      *string           `json:"id,omitempty"`
	Options []FormFieldOption `json:"options,omitempty"`
	Value   *ConnectionValue  `json:"value,omitempty"`
}

func (o *ConnectionDefaults) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *ConnectionDefaults) GetOptions() []FormFieldOption {
	if o == nil {
		return nil
	}
	return o.Options
}

func (o *ConnectionDefaults) GetValue() *ConnectionValue {
	if o == nil {
		return nil
	}
	return o.Value
}

type ConnectionConfiguration struct {
	Resource *string              `json:"resource,omitempty"`
	Defaults []ConnectionDefaults `json:"defaults,omitempty"`
}

func (o *ConnectionConfiguration) GetResource() *string {
	if o == nil {
		return nil
	}
	return o.Resource
}

func (o *ConnectionConfiguration) GetDefaults() []ConnectionDefaults {
	if o == nil {
		return nil
	}
	return o.Defaults
}

type ConnectionInput struct {
	// Whether the connection is enabled or not. You can enable or disable a connection using the Update Connection API.
	Enabled *bool `json:"enabled,omitempty"`
	// Connection settings. Values will persist to `form_fields` with corresponding id
	Settings map[string]any `json:"settings,omitempty"`
	// Attach your own consumer specific metadata
	Metadata      map[string]any            `json:"metadata,omitempty"`
	Configuration []ConnectionConfiguration `json:"configuration,omitempty"`
	// List of custom mappings configured for this connection
	CustomMappings []CustomMappingInput `json:"custom_mappings,omitempty"`
}

func (o *ConnectionInput) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *ConnectionInput) GetSettings() map[string]any {
	if o == nil {
		return nil
	}
	return o.Settings
}

func (o *ConnectionInput) GetMetadata() map[string]any {
	if o == nil {
		return nil
	}
	return o.Metadata
}

func (o *ConnectionInput) GetConfiguration() []ConnectionConfiguration {
	if o == nil {
		return nil
	}
	return o.Configuration
}

func (o *ConnectionInput) GetCustomMappings() []CustomMappingInput {
	if o == nil {
		return nil
	}
	return o.CustomMappings
}
