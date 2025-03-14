package deeplink

import (
	"fmt"
	"net/url"
)

func NewReferralLink(inviteType InvitationType, code string) Deeplink {
	p := ReferralLink{
		invitationType: inviteType,
		invitationCode: code,
	}
	return &p
}

type ReferralLink struct {
	invitationType InvitationType
	invitationCode string
}

func (p *ReferralLink) Build() (string, error) {
	v := fmt.Sprintf(string(LoginValue), p.invitationType, p.invitationCode)
	encodedValue := url.QueryEscape(v)
	link := fmt.Sprintf("%s?af_xp=custom&pid=Apen_dev&c=%s&deep_link_value=%s&af_dp=%s&af_force_deeplink=true", baseUrl, ReferralCampaign, encodedValue, encodedValue)
	return link, nil
}
