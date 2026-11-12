package deeplink

import (
	"errors"
	"fmt"
	"net/url"
)

// NewEndorsementInvitationLink links to the user's endorsement tab.
func NewEndorsementInvitationLink(platform Platform, userID string) Deeplink {
	return &endorsementInvitationLink{platform: platform, userID: userID}
}

type endorsementInvitationLink struct {
	platform Platform
	userID   string
}

func (l *endorsementInvitationLink) Build() (string, error) {
	config, ok := PlatformConfigs[l.platform]
	if !ok {
		return "", fmt.Errorf("unknown platform %d", l.platform)
	}
	if l.userID == "" {
		return "", errors.New("empty user id")
	}

	baseURL, err := url.Parse(config.BaseURL)
	if err != nil {
		return "", fmt.Errorf("invalid base URL: %w", err)
	}

	path := fmt.Sprintf(string(EndorsementValue), url.PathEscape(l.userID))
	deeplinkURL := config.URLScheme + path

	params := url.Values{}
	params.Add("c", string(EndorsementInvitationCampaign))
	params.Add("pid", "copy_link")
	params.Add("af_xp", "custom")
	params.Add("deep_link_value", deeplinkURL)
	params.Add("af_dp", deeplinkURL)
	params.Add("af_force_deeplink", "true")
	if config.WebBaseURL != "" {
		params.Add("af_web_dp", config.WebBaseURL+"/"+path)
	}

	baseURL.RawQuery = params.Encode()
	return baseURL.String(), nil
}
