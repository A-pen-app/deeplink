package deeplink

import (
	"fmt"
	"net/url"
)

// NewTipReceivedLink is the tipped post's link in the tip_received EDM
// (PS-4021): it opens the post in the app, or on the web without the app.
func NewTipReceivedLink(platform Platform, postID string) Deeplink {
	return &TipReceivedLink{platform: platform, postID: postID}
}

type TipReceivedLink struct {
	platform Platform
	postID   string
}

func (p *TipReceivedLink) Build() (string, error) {
	config := PlatformConfigs[p.platform]

	baseURL, err := url.Parse(config.BaseURL)
	if err != nil {
		return "", fmt.Errorf("invalid base URL: %w", err)
	}

	deeplinkURL := config.URLScheme + fmt.Sprintf(string(PostValue), p.postID)

	params := url.Values{}
	params.Add("af_xp", "email")
	params.Add("pid", "Email")
	params.Add("c", string(TipReceivedCampaign))
	params.Add("deep_link_value", deeplinkURL)
	params.Add("af_dp", deeplinkURL)
	params.Add("af_force_deeplink", "true")
	if config.WebBaseURL != "" {
		params.Add("af_web_dp", config.WebBaseURL+"/post/"+p.postID)
	}

	baseURL.RawQuery = params.Encode()
	return baseURL.String(), nil
}
