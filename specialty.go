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

	// 解析 base URL
	baseURL, err := url.Parse(config.BaseURL)
	if err != nil {
		return "", fmt.Errorf("invalid base URL: %w", err)
	}

	// 組合 deeplink URL
	deeplinkPath := fmt.Sprintf(string(PostValue), p.postID)
	deeplinkURL := config.URLScheme + deeplinkPath
	encodedValue := url.QueryEscape(deeplinkURL)

	// 設定查詢參數
	params := url.Values{}
	params.Add("af_xp", "email")
	params.Add("pid", "Email")
	params.Add("c", string(SpecialtyCampaign))
	params.Add("deep_link_value", encodedValue)
	params.Add("af_dp", encodedValue)
	params.Add("af_force_deeplink", "true")

	// 使用 url.URL 組建最終 URL
	baseURL.RawQuery = params.Encode()

	return baseURL.String(), nil
}
