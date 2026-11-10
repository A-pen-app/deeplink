package deeplink

// NewActivityPostLink is the activity post's link in the activity EDMs
// (PS-4399).
func NewActivityPostLink(platform Platform, postID string) Deeplink {
	return &emailPostLink{platform: platform, postID: postID, campaign: ActivityPostCampaign}
}
