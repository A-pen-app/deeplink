package deeplink

import (
	"net/url"
	"testing"
)

func TestShareShortVideoLinkBuild(t *testing.T) {
	const id = "0f1e2d3c-4b5a-6978-8a9b-0c1d2e3f4a5b"

	cases := []struct {
		name     string
		platform Platform
		base     string
		deepLink string
		webLink  string
	}{
		{
			name:     "apen",
			platform: PlatformApen,
			base:     "https://apen.penpeer.co/sJck",
			deepLink: "apen://short_videos/" + id,
			webLink:  "https://www.a-pen.co/short_videos/" + id,
		},
		{
			name:     "phar",
			platform: PlatformPhar,
			base:     "https://phar.penpeer.co/9db5",
			deepLink: "phar://short_videos/" + id,
			webLink:  "https://phar-web.penpeer.co/short_videos/" + id,
		},
		{
			name:     "nurse",
			platform: PlatformNurse,
			base:     "https://nurse.penpeer.co/cLnc",
			deepLink: "nstation://short_videos/" + id,
			webLink:  "https://nurse-web.penpeer.co/short_videos/" + id,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			raw, err := NewShareShortVideoLink(tc.platform, id).Build()
			if err != nil {
				t.Fatalf("Build() error: %v", err)
			}

			u, err := url.Parse(raw)
			if err != nil {
				t.Fatalf("url.Parse(%q) error: %v", raw, err)
			}

			gotBase := u.Scheme + "://" + u.Host + u.Path
			if gotBase != tc.base {
				t.Errorf("base URL = %q, want %q", gotBase, tc.base)
			}

			q := u.Query()
			want := map[string]string{
				"c":                 "share_short_video",
				"deep_link_value":   tc.deepLink,
				"af_dp":             tc.deepLink,
				"af_web_dp":         tc.webLink,
				"af_xp":             "custom",
				"af_force_deeplink": "true",
			}
			for key, val := range want {
				if got := q.Get(key); got != val {
					t.Errorf("query %q = %q, want %q", key, got, val)
				}
			}
		})
	}
}

func TestShareShortVideoLinkOmitsWebDeeplinkWithoutWebBaseURL(t *testing.T) {
	const platform Platform = 999
	original, existed := PlatformConfigs[platform]
	PlatformConfigs[platform] = PlatformConfig{
		BaseURL:   "https://example.penpeer.co/x",
		URLScheme: "example://",
		Name:      "Example",
	}
	t.Cleanup(func() {
		if existed {
			PlatformConfigs[platform] = original
			return
		}
		delete(PlatformConfigs, platform)
	})

	raw, err := NewShareShortVideoLink(platform, "abc").Build()
	if err != nil {
		t.Fatalf("Build() error: %v", err)
	}
	u, err := url.Parse(raw)
	if err != nil {
		t.Fatalf("url.Parse(%q) error: %v", raw, err)
	}
	if _, present := u.Query()["af_web_dp"]; present {
		t.Errorf("af_web_dp should be omitted when WebBaseURL is empty, got %q", raw)
	}
	if got := u.Query().Get("deep_link_value"); got != "example://short_videos/abc" {
		t.Errorf("deep_link_value = %q, want %q", got, "example://short_videos/abc")
	}
}
