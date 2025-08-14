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

	// 解析 base URL
	baseURL, err := url.Parse(config.BaseURL)
	if err != nil {
		return "", fmt.Errorf("invalid base URL: %w", err)
	}

	// 組合 deeplink URL
	deeplinkPath := fmt.Sprintf(string(LoginValue), InvitationTypeID, p.invitationCode)
	deeplinkURL := config.URLScheme + deeplinkPath
	encodedValue := url.QueryEscape(deeplinkURL)

	// 設定查詢參數
	params := url.Values{}
	params.Add("af_xp", "custom")
	params.Add("pid", config.Name+"_dev")
	params.Add("c", string(ShareJoinCampaign))
	params.Add("deep_link_value", encodedValue)
	params.Add("af_dp", encodedValue)
	params.Add("af_force_deeplink", "true")

	// 使用 url.URL 組建最終 URL
	baseURL.RawQuery = params.Encode()

	return baseURL.String(), nil
}
