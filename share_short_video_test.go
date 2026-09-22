package deeplink

import (
	"net/url"
	"testing"
)

func TestShareShortVideoLinkBuild(t *testing.T) {
	const id = "0f1e2d3c-4b5a-6978-8a9b-0c1d2e3f4a5b"

	cases := []struct {
		name         string
		platform     Platform
		wantBase     string
		wantDeepLink string
		wantWebLink  string
	}{
		{
			name:         "apen",
			platform:     PlatformApen,
			wantBase:     "https://apen.penpeer.co/sJck",
			wantDeepLink: "apen://short_videos/" + id,
			wantWebLink:  "https://www.a-pen.co/short_videos/" + id,
		},
		{
			name:         "phar",
			platform:     PlatformPhar,
			wantBase:     "https://phar.penpeer.co/9db5",
			wantDeepLink: "phar://short_videos/" + id,
			wantWebLink:  "https://phar-web.penpeer.co/short_videos/" + id,
		},
		{
			name:         "nurse",
			platform:     PlatformNurse,
			wantBase:     "https://nurse.penpeer.co/cLnc",
			wantDeepLink: "nstation://short_videos/" + id,
			wantWebLink:  "https://nurse-web.penpeer.co/short_videos/" + id,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			u := buildShareShortVideoURL(t, tc.platform, id)

			if gotBase := u.Scheme + "://" + u.Host + u.Path; gotBase != tc.wantBase {
				t.Errorf("base URL = %q, want %q", gotBase, tc.wantBase)
			}

			assertQueryParams(t, u.Query(), map[string]string{
				"c":                 "share_short_video",
				"deep_link_value":   tc.wantDeepLink,
				"af_dp":             tc.wantDeepLink,
				"af_web_dp":         tc.wantWebLink,
				"af_xp":             "custom",
				"af_force_deeplink": "true",
			})
		})
	}
}

func TestShareShortVideoLinkOmitsWebDeeplinkWithoutWebBaseURL(t *testing.T) {
	platform := registerTemporaryPlatform(t, PlatformConfig{
		BaseURL:   "https://example.penpeer.co/x",
		URLScheme: "example://",
		Name:      "Example",
	})

	u := buildShareShortVideoURL(t, platform, "abc")

	if _, present := u.Query()["af_web_dp"]; present {
		t.Errorf("af_web_dp should be omitted when WebBaseURL is empty, got %q", u.String())
	}
	assertQueryParams(t, u.Query(), map[string]string{
		"deep_link_value": "example://short_videos/abc",
	})
}

func TestShareShortVideoLinkRejectsInvalidBaseURL(t *testing.T) {
	platform := registerTemporaryPlatform(t, PlatformConfig{
		BaseURL:   "://missing-scheme",
		URLScheme: "example://",
		Name:      "Example",
	})

	if _, err := NewShareShortVideoLink(platform, "abc").Build(); err == nil {
		t.Fatal("Build() error = nil, want invalid base URL error")
	}
}

// buildShareShortVideoURL builds the share link and parses it so tests can
// assert on individual query parameters instead of the encoded string.
func buildShareShortVideoURL(t *testing.T, platform Platform, shortVideoID string) *url.URL {
	t.Helper()

	raw, err := NewShareShortVideoLink(platform, shortVideoID).Build()
	if err != nil {
		t.Fatalf("Build() error: %v", err)
	}
	u, err := url.Parse(raw)
	if err != nil {
		t.Fatalf("url.Parse(%q) error: %v", raw, err)
	}
	return u
}

func assertQueryParams(t *testing.T, got url.Values, want map[string]string) {
	t.Helper()

	for key, val := range want {
		if got.Get(key) != val {
			t.Errorf("query %q = %q, want %q", key, got.Get(key), val)
		}
	}
}

// registerTemporaryPlatform installs a PlatformConfig under an unused
// Platform key for the duration of the test and removes it afterwards.
func registerTemporaryPlatform(t *testing.T, config PlatformConfig) Platform {
	t.Helper()

	const platform Platform = 999
	if _, exists := PlatformConfigs[platform]; exists {
		t.Fatalf("Platform %d is already registered; pick another test key", platform)
	}
	PlatformConfigs[platform] = config
	t.Cleanup(func() { delete(PlatformConfigs, platform) })
	return platform
}
