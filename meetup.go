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

	// 解析 base URL
	baseURL, err := url.Parse(config.BaseURL)
	if err != nil {
		return "", fmt.Errorf("invalid base URL: %w", err)
	}

	// 組合 deeplink URL
	deeplinkPath := fmt.Sprintf(string(MeetupAttendValue), p.meetupID)
	deeplinkURL := config.URLScheme + deeplinkPath
	encodedDeeplinkValue := url.QueryEscape(deeplinkURL)

	// 設定查詢參數
	params := url.Values{}
	params.Add("af_xp", "email")
	params.Add("pid", "Email")
	params.Add("c", string(MeetupAttendCampaign))
	params.Add("deep_link_value", encodedDeeplinkValue)
	params.Add("af_dp", encodedDeeplinkValue)
	params.Add("af_force_deeplink", "true")

	// 使用 url.URL 組建最終 URL
	baseURL.RawQuery = params.Encode()

	return baseURL.String(), nil
}
