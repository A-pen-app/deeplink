package deeplink

import (
	"errors"
	"fmt"
	"net/url"
)

// emailPostLink is the link to a post in an EDM: an AppsFlyer OneLink that
// opens the post in the app, or on the web without it. Each mail has its own
// campaign.
type emailPostLink struct {
	platform Platform
	postID   string
	campaign DeeplinkCampaign
}

func (p *emailPostLink) Build() (string, error) {
	config, ok := PlatformConfigs[p.platform]
	if !ok {
		return "", fmt.Errorf("unknown platform %d", p.platform)
	}
	if p.postID == "" {
		return "", errors.New("empty post id")
	}

	baseURL, err := url.Parse(config.BaseURL)
	if err != nil {
		return "", fmt.Errorf("invalid base URL: %w", err)
	}

	postID := url.PathEscape(p.postID)
	deeplinkURL := config.URLScheme + fmt.Sprintf(string(PostValue), postID)

	params := url.Values{}
	params.Add("af_xp", "email")
	params.Add("pid", "Email")
	params.Add("c", string(p.campaign))
	params.Add("deep_link_value", deeplinkURL)
	params.Add("af_dp", deeplinkURL)
	params.Add("af_force_deeplink", "true")
	if config.WebBaseURL != "" {
		params.Add("af_web_dp", config.WebBaseURL+"/post/"+postID)
	}

	baseURL.RawQuery = params.Encode()
	return baseURL.String(), nil
}
