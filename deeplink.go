package deeplink

type Deeplink interface {
	Build() (string, error)
}

const baseUrl string = "https://apen.penpeer.co/sJck"

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
	LoginValue        DeeplinkValue = "apen://login/?type=%v&code=%s"
	PostValue         DeeplinkValue = "apen://posts/%s"
	RewardValue       DeeplinkValue = "apen://shop/wallet_history/?tab=%s"
	MissionsValue     DeeplinkValue = "apen://missions/mission_list"
	MeetupAttendValue DeeplinkValue = "apen://meetups/%s"
)
