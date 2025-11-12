package deeplink

import (
	"fmt"
	"net/url"
)

func NewResumeReadLink(platform Platform, chatID string) Deeplink {
	p := ResumeReadLink{
		platform: platform,
		chatID:   chatID,
	}
	return &p
}

type ResumeReadLink struct {
	platform Platform
	chatID   string
}

func (p *ResumeReadLink) Build() (string, error) {
	config := PlatformConfigs[p.platform]

	// 解析 base URL
	baseURL, err := url.Parse(config.BaseURL)
	if err != nil {
		return "", fmt.Errorf("invalid base URL: %w", err)
	}

	// 組合 deeplink URL
	deeplinkPath := fmt.Sprintf(string(HireChatValue), p.chatID)
	deeplinkURL := config.URLScheme + deeplinkPath

	// 設定查詢參數
	params := url.Values{}
	params.Add("af_xp", "email")
	params.Add("pid", "Email")
	params.Add("c", string(ResumeReadCampaign))
	params.Add("deep_link_value", deeplinkURL)
	params.Add("af_dp", deeplinkURL)
	params.Add("af_force_deeplink", "true")

	// 使用 url.URL 組建最終 URL
	baseURL.RawQuery = params.Encode()

	return baseURL.String(), nil
}
