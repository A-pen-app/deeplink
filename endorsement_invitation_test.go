package deeplink

import (
	"net/url"
	"testing"
)

func TestEndorsementInvitationLinkBuild(t *testing.T) {
	const id = "0f1e2d3c-4b5a-6978-8a9b-0c1d2e3f4a5b"

	cases := []struct {
		name         string
		platform     Platform
		wantBase     string
		wantDeepLink string
		wantWebLink  string
	}{
		{"apen", PlatformApen, "https://apen.penpeer.co/sJck", "apen://profiles/" + id + "/endorsement", "https://www.a-pen.co/profiles/" + id + "/endorsement"},
		{"phar", PlatformPhar, "https://phar.penpeer.co/9db5", "phar://profiles/" + id + "/endorsement", "https://phar-web.penpeer.co/profiles/" + id + "/endorsement"},
		{"nurse", PlatformNurse, "https://nurse.penpeer.co/cLnc", "nstation://profiles/" + id + "/endorsement", "https://nurse-web.penpeer.co/profiles/" + id + "/endorsement"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			u := buildEndorsementInvitationURL(t, tc.platform, id)

			if gotBase := u.Scheme + "://" + u.Host + u.Path; gotBase != tc.wantBase {
				t.Errorf("base URL = %q, want %q", gotBase, tc.wantBase)
			}
			assertQueryParams(t, u.Query(), map[string]string{
				"c":                 "share_profile",
				"pid":               "copy_link",
				"af_xp":             "custom",
				"deep_link_value":   tc.wantDeepLink,
				"af_dp":             tc.wantDeepLink,
				"af_web_dp":         tc.wantWebLink,
				"af_force_deeplink": "true",
			})
		})
	}
}

func TestEndorsementInvitationLinkOmitsWebDeeplinkWithoutWebBaseURL(t *testing.T) {
	platform := registerTemporaryPlatform(t, PlatformConfig{
		BaseURL:   "https://example.penpeer.co/x",
		URLScheme: "example://",
		Name:      "Example",
	})

	u := buildEndorsementInvitationURL(t, platform, "abc")

	if _, present := u.Query()["af_web_dp"]; present {
		t.Errorf("af_web_dp should be omitted when WebBaseURL is empty, got %q", u.String())
	}
	assertQueryParams(t, u.Query(), map[string]string{
		"deep_link_value": "example://profiles/abc/endorsement",
	})
}

func TestEndorsementInvitationLinkRejectsUnknownPlatform(t *testing.T) {
	if _, err := NewEndorsementInvitationLink(Platform(998), "abc").Build(); err == nil {
		t.Fatal("Build() error = nil, want unknown platform error")
	}
}

func TestEndorsementInvitationLinkRejectsEmptyUserID(t *testing.T) {
	if _, err := NewEndorsementInvitationLink(PlatformApen, "").Build(); err == nil {
		t.Fatal("Build() error = nil, want empty user id error")
	}
}

func TestEndorsementInvitationLinkRejectsInvalidBaseURL(t *testing.T) {
	platform := registerTemporaryPlatform(t, PlatformConfig{
		BaseURL:   "://missing-scheme",
		URLScheme: "example://",
		Name:      "Example",
	})

	if _, err := NewEndorsementInvitationLink(platform, "abc").Build(); err == nil {
		t.Fatal("Build() error = nil, want invalid base URL error")
	}
}

func buildEndorsementInvitationURL(t *testing.T, platform Platform, userID string) *url.URL {
	t.Helper()

	raw, err := NewEndorsementInvitationLink(platform, userID).Build()
	if err != nil {
		t.Fatalf("Build() error: %v", err)
	}
	u, err := url.Parse(raw)
	if err != nil {
		t.Fatalf("url.Parse(%q) error: %v", raw, err)
	}
	return u
}
