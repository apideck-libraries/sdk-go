// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type Security struct {
	APIKey *string `security:"scheme,type=apiKey,subtype=header,name=Authorization,env=apideck_api_key"`
}

func (o *Security) GetAPIKey() *string {
	if o == nil {
		return nil
	}
	return o.APIKey
}
