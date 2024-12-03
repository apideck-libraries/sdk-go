package hooks

import (
	"net/http"
)

type AuthHook struct{}

var _ beforeRequestHook = (*AuthHook)(nil)

func (h *AuthHook) BeforeRequest(_ BeforeRequestContext, req *http.Request) (*http.Request, error) {
	if apiKey := req.Header.Get("Authorization"); apiKey != "" {
		req.Header.Set("Authorization", "Bearer "+apiKey)
	}
	return req, nil
}