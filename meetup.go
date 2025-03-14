package deeplink

import (
	"fmt"
	"net/url"
)

func NewMeetupAttendLink(meetupID string) Deeplink {
	p := MeetupAttendLink{
		meetupID: meetupID,
	}
	return &p
}

type MeetupAttendLink struct {
	meetupID string
}

func (p *MeetupAttendLink) Build() (string, error) {
	v := fmt.Sprintf(string(MeetupAttendValue), p.meetupID)
	encodedValue := url.QueryEscape(v)
	link := fmt.Sprintf("%s?af_xp=email&pid=Email&c=%s&deep_link_value=%s&af_dp=%s&af_force_deeplink=true", baseUrl, MeetupAttendCampaign, encodedValue, encodedValue)
	return link, nil
}
