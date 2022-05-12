package SSOLoginRouter

import (
	"encoding/json"
)

type SSOLoginDto struct {
	Name   string          `json:"name"`
	Label  string          `json:"label"`
	Url    string          `json:"url"`
	Config json.RawMessage `json:"config"`
	Active bool            `json:"active"`
}

//todo need to re-write this after discussion as data created via this API is not delete-able

/*func TestCreateSsoLoginWithCorrectArgs(t *testing.T) {
	payload := &SSOLoginDto{"Deepak", "Deepaklabel", "UrlDeepak", json.RawMessage("ConfigDeepak"), true}
	var e1 []byte
	e1, _ = json.Marshal(payload)
	resp, err := Base.MakeApiCall("/orchestrator/sso/create", http.MethodPost, string(e1), nil, "")
	Base.HandleError(err, "createSsoLoginWithCorrectArgs")
	assert.Equal(t, 200, resp.StatusCode())
}
*/
