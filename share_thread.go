package deeplink

import (
	"fmt"
	"net/url"
)

func NewShareThreadLink(platform Platform, id string) Deeplink {
	return &ShareThreadLink{
		platform: platform,
		id:       id,
	}
}

type ShareThreadLink struct {
	platform Platform
	id       string
}

func (s *ShareThreadLink) Build() (string, error) {
	config := PlatformConfigs[s.platform]

	baseURL, err := url.Parse(config.BaseURL)
	if err != nil {
		return "", fmt.Errorf("invalid base URL: %w", err)
	}

	deeplinkPath := fmt.Sprintf(string(AIShareValue), s.id)
	deeplinkURL := config.URLScheme + deeplinkPath
	webDeeplinkURL := config.WebBaseURL + "/ai/gpt/share/" + s.id

	params := url.Values{}
	params.Add("c", string(ShareThreadCampaign))
	params.Add("af_xp", "custom")
	params.Add("deep_link_value", deeplinkURL)
	params.Add("af_dp", deeplinkURL)
	params.Add("af_force_deeplink", "true")
	if config.WebBaseURL != "" {
		params.Add("af_web_dp", webDeeplinkURL)
	}

	baseURL.RawQuery = params.Encode()
	return baseURL.String(), nil
}
