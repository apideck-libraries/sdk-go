// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// GetAPIResponse - Apis
type GetAPIResponse struct {
	// HTTP Response Status Code
	StatusCode int64 `json:"status_code"`
	// HTTP Response Status
	Status string `json:"status"`
	Data   API    `json:"data"`
	// Response metadata
	Meta *Meta `json:"meta,omitempty"`
	// Links to navigate to previous or next pages through the API
	Links *Links `json:"links,omitempty"`
}

func (o *GetAPIResponse) GetStatusCode() int64 {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetAPIResponse) GetStatus() string {
	if o == nil {
		return ""
	}
	return o.Status
}

func (o *GetAPIResponse) GetData() API {
	if o == nil {
		return API{}
	}
	return o.Data
}

func (o *GetAPIResponse) GetMeta() *Meta {
	if o == nil {
		return nil
	}
	return o.Meta
}

func (o *GetAPIResponse) GetLinks() *Links {
	if o == nil {
		return nil
	}
	return o.Links
}
