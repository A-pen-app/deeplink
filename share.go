package deeplink

import (
	"fmt"
	"net/url"
)

func NewShareJoinLink(inviteType InvitationType, code string) Deeplink {
	p := ShareJoinLink{
		invitationType: inviteType,
		invitationCode: code,
	}
	return &p
}

type ShareJoinLink struct {
	invitationType InvitationType
	invitationCode string
}

func (p *ShareJoinLink) Build() (string, error) {
	v := fmt.Sprintf(string(LoginDeeplinkValue), p.invitationType, p.invitationCode)
	encodedValue := url.QueryEscape(v)
	link := fmt.Sprintf("%s?af_xp=custom&pid=Apen_dev&c=%s&deep_link_value=%s&af_dp=%s&af_force_deeplink=true", baseUrl, ShareJoinCampaign, encodedValue, encodedValue)
	return link, nil
}
