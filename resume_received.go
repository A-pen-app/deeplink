package deeplink

import (
	"fmt"
	"net/url"
)

func NewResumeReceivedLink(platform Platform, chatID string) Deeplink {
	p := ResumeReceivedLink{
		platform: platform,
		chatID:   chatID,
	}
	return &p
}

type ResumeReceivedLink struct {
	platform Platform
	chatID   string
}

func (p *ResumeReceivedLink) Build() (string, error) {
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
	params.Add("c", string(ResumeReceiveCampaign))
	params.Add("deep_link_value", deeplinkURL)
	params.Add("af_dp", deeplinkURL)
	params.Add("af_force_deeplink", "true")

	// 使用 url.URL 組建最終 URL
	baseURL.RawQuery = params.Encode()

	return baseURL.String(), nil
}
