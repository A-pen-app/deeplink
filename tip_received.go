package deeplink

// NewTipReceivedLink is the tipped post's link in the tip_received EDM
// (PS-4021).
func NewTipReceivedLink(platform Platform, postID string) Deeplink {
	return &emailPostLink{platform: platform, postID: postID, campaign: TipReceivedCampaign}
}
