package deeplink

import (
	"fmt"
	"net/url"
)

func NewBusinessCardReceivedLink(platform Platform, chatID string) Deeplink {
	return &BusinessCardReceivedLink{
		platform: platform,
		chatID:   chatID,
	}
}

type BusinessCardReceivedLink struct {
	platform Platform
	chatID   string
}

func (p *BusinessCardReceivedLink) Build() (string, error) {
	config := PlatformConfigs[p.platform]

	baseURL, err := url.Parse(config.BaseURL)
	if err != nil {
		return "", fmt.Errorf("invalid base URL: %w", err)
	}

	deeplinkPath := fmt.Sprintf(string(HireChatValue), p.chatID)
	deeplinkURL := config.URLScheme + deeplinkPath

	params := url.Values{}
	params.Add("af_xp", "email")
	params.Add("pid", "Email")
	params.Add("c", string(BusinessCardSentCampaign))
	params.Add("deep_link_value", deeplinkURL)
	params.Add("af_dp", deeplinkURL)
	params.Add("af_force_deeplink", "true")

	baseURL.RawQuery = params.Encode()

	return baseURL.String(), nil
}
