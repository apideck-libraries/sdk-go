// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type CreateCallbackState struct {
	// The redirect URI to be used after the connection is created.
	RedirectURI *string `json:"redirect_uri,omitempty"`
}

func (o *CreateCallbackState) GetRedirectURI() *string {
	if o == nil {
		return nil
	}
	return o.RedirectURI
}
