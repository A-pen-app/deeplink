package deeplink

import (
	"fmt"
	"net/url"
)

func NewShareShortVideoLink(platform Platform, shortVideoID string) Deeplink {
	return &ShareShortVideoLink{
		platform:     platform,
		shortVideoID: shortVideoID,
	}
}

type ShareShortVideoLink struct {
	platform     Platform
	shortVideoID string
}

func (s *ShareShortVideoLink) Build() (string, error) {
	config := PlatformConfigs[s.platform]

	baseURL, err := url.Parse(config.BaseURL)
	if err != nil {
		return "", fmt.Errorf("invalid base URL: %w", err)
	}

	deeplinkPath := fmt.Sprintf(string(ShortVideoValue), s.shortVideoID)
	deeplinkURL := config.URLScheme + deeplinkPath
	webDeeplinkURL := config.WebBaseURL + "/short_videos/" + s.shortVideoID

	params := url.Values{}
	params.Add("c", string(ShareShortVideoCampaign))
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
