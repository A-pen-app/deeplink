package deeplink

import (
	"fmt"
	"net/url"
)

func NewMeetupAttendLink(platform Platform, meetupID string) Deeplink {
	p := MeetupAttendLink{
		platform: platform,
		meetupID: meetupID,
	}
	return &p
}

type MeetupAttendLink struct {
	platform Platform
	meetupID string
}

func (p *MeetupAttendLink) Build() (string, error) {
	config := platformConfigs[p.platform]

	deeplinkPath := fmt.Sprintf(string(MeetupAttendValue), p.meetupID)
	deeplinkURL := config.URLScheme + deeplinkPath
	encodedDeeplinkValue := url.QueryEscape(deeplinkURL)

	params := url.Values{}
	params.Add("af_xp", "email")
	params.Add("pid", "Email")
	params.Add("c", string(MeetupAttendCampaign))
	params.Add("deep_link_value", encodedDeeplinkValue)
	params.Add("af_dp", encodedDeeplinkValue)
	params.Add("af_force_deeplink", "true")

	return config.BaseURL + "?" + params.Encode(), nil
}
