package deeplink

import (
	"fmt"
	"net/url"
)

func NewSubscriptionPauseReminderLink(platform Platform) Deeplink {
	return &SubscriptionPauseReminderLink{
		platform: platform,
	}
}

type SubscriptionPauseReminderLink struct {
	platform Platform
}

func (p *SubscriptionPauseReminderLink) Build() (string, error) {
	config := PlatformConfigs[p.platform]

	baseURL, err := url.Parse(config.BaseURL)
	if err != nil {
		return "", fmt.Errorf("invalid base URL: %w", err)
	}

	deeplinkURL := config.URLScheme + string(SubscriptionManagementValue)

	params := url.Values{}
	params.Add("af_xp", "email")
	params.Add("pid", "Email")
	params.Add("c", string(SubscriptionPauseReminderCampgn))
	params.Add("deep_link_value", deeplinkURL)
	params.Add("af_dp", deeplinkURL)
	params.Add("af_force_deeplink", "true")

	baseURL.RawQuery = params.Encode()

	return baseURL.String(), nil
}
