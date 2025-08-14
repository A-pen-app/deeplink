package deeplink

import (
	"fmt"
	"net/url"
)

func NewReferralLink(platform Platform, code string) (Deeplink, error) {
	if !isValid6DigitCode(code) {
		return nil, fmt.Errorf("invalid 6-digit code format: %s", code)
	}
	p := ReferralLink{
		platform:       platform,
		invitationCode: code,
	}
	return &p, nil
}

type ReferralLink struct {
	platform       Platform
	invitationCode string
}

func (p *ReferralLink) Build() (string, error) {
	config := platformConfigs[p.platform]

	// 解析 base URL
	baseURL, err := url.Parse(config.BaseURL)
	if err != nil {
		return "", fmt.Errorf("invalid base URL: %w", err)
	}

	// 組合 deeplink URL
	deeplinkPath := fmt.Sprintf(string(LoginValue), InvitationTypeCode, p.invitationCode)
	deeplinkURL := config.URLScheme + deeplinkPath

	// 設定查詢參數
	params := url.Values{}
	params.Add("af_xp", "custom")
	params.Add("pid", config.Name+"_dev")
	params.Add("c", string(ReferralCampaign))
	params.Add("deep_link_value", deeplinkURL)
	params.Add("af_dp", deeplinkURL)
	params.Add("af_force_deeplink", "true")

	// 使用 url.URL 組建最終 URL
	baseURL.RawQuery = params.Encode()

	return baseURL.String(), nil
}
