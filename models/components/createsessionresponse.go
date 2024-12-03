// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type CreateSessionResponseData struct {
	SessionURI   string `json:"session_uri"`
	SessionToken string `json:"session_token"`
}

func (o *CreateSessionResponseData) GetSessionURI() string {
	if o == nil {
		return ""
	}
	return o.SessionURI
}

func (o *CreateSessionResponseData) GetSessionToken() string {
	if o == nil {
		return ""
	}
	return o.SessionToken
}

// CreateSessionResponse - Session created
type CreateSessionResponse struct {
	// HTTP Response Status Code
	StatusCode int64 `json:"status_code"`
	// HTTP Response Status
	Status string                    `json:"status"`
	Data   CreateSessionResponseData `json:"data"`
}

func (o *CreateSessionResponse) GetStatusCode() int64 {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateSessionResponse) GetStatus() string {
	if o == nil {
		return ""
	}
	return o.Status
}

func (o *CreateSessionResponse) GetData() CreateSessionResponseData {
	if o == nil {
		return CreateSessionResponseData{}
	}
	return o.Data
}
