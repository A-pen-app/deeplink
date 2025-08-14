package deeplink

import (
	"fmt"
	"net/url"
)

func NewSpecialtyPostLink(platform Platform, postID string) Deeplink {
	p := SpecialtyPostLink{
		platform: platform,
		postID:   postID,
	}
	return &p
}

type SpecialtyPostLink struct {
	platform Platform
	postID   string
}

func (p *SpecialtyPostLink) Build() (string, error) {
	config := platformConfigs[p.platform]

	deeplinkPath := fmt.Sprintf(string(PostValue), p.postID)
	deeplinkURL := config.URLScheme + deeplinkPath
	encodedValue := url.QueryEscape(deeplinkURL)

	params := url.Values{}
	params.Add("af_xp", "email")
	params.Add("pid", "Email")
	params.Add("c", string(SpecialtyCampaign))
	params.Add("deep_link_value", encodedValue)
	params.Add("af_dp", encodedValue)
	params.Add("af_force_deeplink", "true")

	return config.BaseURL + "?" + params.Encode(), nil
}
