package deeplink

import (
	"fmt"
	"net/url"
)

func NewSpecialtyPostLink(postID string) Deeplink {
	p := SpecialtyPostLink{
		postID: postID,
	}
	return &p
}

type SpecialtyPostLink struct {
	postID string
}

func (p *SpecialtyPostLink) Build() (string, error) {
	v := fmt.Sprintf(string(PostValue), p.postID)
	encodedValue := url.QueryEscape(v)
	link := fmt.Sprintf("%s?af_xp=email&pid=Email&c=%s&deep_link_value=%s&af_dp=%s&af_force_deeplink=true", baseUrl, SpecialtyCampaign, encodedValue, encodedValue)
	return link, nil
}
