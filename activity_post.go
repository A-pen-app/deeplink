package deeplink

import (
	"fmt"
	"net/url"
)

func NewActivityPostLink(platform Platform, postID string) Deeplink {
	return &ActivityPostLink{platform: platform, postID: postID}
}

type ActivityPostLink struct {
	platform Platform
	postID   string
}

func (p *ActivityPostLink) Build() (string, error) {
	config := PlatformConfigs[p.platform]

	baseURL, err := url.Parse(config.BaseURL)
	if err != nil {
		return "", fmt.Errorf("invalid base URL: %w", err)
	}

	deeplinkURL := config.URLScheme + fmt.Sprintf(string(PostValue), p.postID)

	params := url.Values{}
	params.Add("af_xp", "email")
	params.Add("pid", "Email")
	params.Add("c", string(ActivityPostCampaign))
	params.Add("deep_link_value", deeplinkURL)
	params.Add("af_dp", deeplinkURL)
	params.Add("af_force_deeplink", "true")
	if config.WebBaseURL != "" {
		params.Add("af_web_dp", config.WebBaseURL+"/post/"+p.postID)
	}

	baseURL.RawQuery = params.Encode()
	return baseURL.String(), nil
}
