package deeplink

import (
	"regexp"
)

type Deeplink interface {
	Build() (string, error)
}

type Platform int

const (
	PlatformApen Platform = iota
	PlatformPhar
	PlatformNurse
)

type PlatformConfig struct {
	BaseURL   string
	URLScheme string
	Name      string
}

var platformConfigs = map[Platform]PlatformConfig{
	PlatformApen:  {BaseURL: "https://apen.penpeer.co/sJck", URLScheme: "apen://", Name: "Apen"},
	PlatformPhar:  {BaseURL: "https://phar.penpeer.co/9db5", URLScheme: "phar://", Name: "Phar"},
	PlatformNurse: {BaseURL: "https://nurse.penpeer.co/cLnc", URLScheme: "nstation://", Name: "Nurse"},
}

type DeeplinkType int

const (
	ShareJoin DeeplinkType = iota
	SpecialtyPost
	Referral
	MeetupAttend
)

type DeeplinkCampaign string

const (
	ShareJoinCampaign    DeeplinkCampaign = "share_and_join"
	SpecialtyCampaign    DeeplinkCampaign = "專科專區"
	ReferralCampaign     DeeplinkCampaign = "referral"
	MeetupAttendCampaign DeeplinkCampaign = "活動"
)

type DeeplinkValue string

const (
	LoginValue        DeeplinkValue = "login/?type=%v&code=%s"
	PostValue         DeeplinkValue = "posts/%s"
	RewardValue       DeeplinkValue = "shop/wallet_history/?tab=%s"
	MissionsValue     DeeplinkValue = "missions/mission_list"
	MeetupAttendValue DeeplinkValue = "meetups/%s"
)

func isValidUUID(s string) bool {
	matched, _ := regexp.MatchString(`^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$`, s)
	return matched
}

func isValid6DigitCode(s string) bool {
	matched, _ := regexp.MatchString(`^[A-Za-z0-9]{6}$`, s)
	return matched
}

type InvitationType int

const (
	InvitationTypeID InvitationType = iota
	InvitationTypeCode
)
