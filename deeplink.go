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
)

type DeeplinkCampaign string

const (
	ShareJoinCampaign DeeplinkCampaign = "share_and_join"
	SpecialtyCampaign DeeplinkCampaign = "專科專區"
	ReferralCampaign  DeeplinkCampaign = "referral"
)

type DeeplinkValue string

const (
	LoginDeeplinkValue    DeeplinkValue = "apen://login/?type=%v&code=%s"
	PostDeeplinkValue     DeeplinkValue = "apen://posts/%s"
	RewardDeeplinkValue   DeeplinkValue = "apen://shop/wallet_history/?tab=%s"
	MissionsDeeplinkValue DeeplinkValue = "apen://missions/mission_list"
)
