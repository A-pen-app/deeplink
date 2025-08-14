package deeplink

import (
	"fmt"
	"net/url"
)

func NewShareJoinLink(platform Platform, code string) (Deeplink, error) {
	if !isValidUUID(code) {
		return nil, fmt.Errorf("invalid UUID format: %s", code)
	}
	p := ShareJoinLink{
		platform:       platform,
		invitationCode: code,
	}
	return &p, nil
}

type ShareJoinLink struct {
	platform       Platform
	invitationCode string
}

func (p *ShareJoinLink) Build() (string, error) {
	config := platformConfigs[p.platform]

	deeplinkPath := fmt.Sprintf(string(LoginValue), InvitationTypeID, p.invitationCode)
	deeplinkURL := config.URLScheme + deeplinkPath
	encodedValue := url.QueryEscape(deeplinkURL)

	params := url.Values{}
	params.Add("af_xp", "custom")
	params.Add("pid", config.Name+"_dev")
	params.Add("c", string(ShareJoinCampaign))
	params.Add("deep_link_value", encodedValue)
	params.Add("af_dp", encodedValue)
	params.Add("af_force_deeplink", "true")

	return config.BaseURL + "?" + params.Encode(), nil
}
