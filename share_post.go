package deeplink

import (
	"fmt"
	"net/url"
)

func NewSharePostLink(platform Platform, postID string) (Deeplink, error) {
	p := SharePostLink{
		platform: platform,
		postID:   postID,
	}
	return &p, nil
}

type SharePostLink struct {
	platform Platform
	postID   string
}

func (p *SharePostLink) Build() (string, error) {
	config := platformConfigs[p.platform]

	deeplinkPath := fmt.Sprintf(string(PostValue), p.postID)
	deeplinkURL := config.URLScheme + deeplinkPath
	encodedValue := url.QueryEscape(deeplinkURL)

	params := url.Values{}
	params.Add("af_xp", "custom")
	params.Add("deep_link_value", encodedValue)
	params.Add("af_dp", encodedValue)
	params.Add("af_force_deeplink", "true")

	return config.BaseURL + "?" + params.Encode(), nil
}
