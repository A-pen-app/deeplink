package deeplink

import (
	"net/url"
	"testing"
)

func TestActivityPostLinkBuild(t *testing.T) {
	const id = "0f1e2d3c-4b5a-6978-8a9b-0c1d2e3f4a5b"

	cases := []struct {
		name         string
		platform     Platform
		wantBase     string
		wantDeepLink string
		wantWebLink  string
	}{
		{"apen", PlatformApen, "https://apen.penpeer.co/sJck", "apen://posts/" + id, "https://www.a-pen.co/post/" + id},
		{"phar", PlatformPhar, "https://phar.penpeer.co/9db5", "phar://posts/" + id, "https://phar-web.penpeer.co/post/" + id},
		{"nurse", PlatformNurse, "https://nurse.penpeer.co/cLnc", "nstation://posts/" + id, "https://nurse-web.penpeer.co/post/" + id},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			raw, err := NewActivityPostLink(tc.platform, id).Build()
			if err != nil {
				t.Fatalf("Build() error: %v", err)
			}
			u, err := url.Parse(raw)
			if err != nil {
				t.Fatalf("url.Parse(%q) error: %v", raw, err)
			}
			if gotBase := u.Scheme + "://" + u.Host + u.Path; gotBase != tc.wantBase {
				t.Errorf("base URL = %q, want %q", gotBase, tc.wantBase)
			}
			assertQueryParams(t, u.Query(), map[string]string{
				"c":                 "活動貼文",
				"pid":               "Email",
				"af_xp":             "email",
				"deep_link_value":   tc.wantDeepLink,
				"af_dp":             tc.wantDeepLink,
				"af_web_dp":         tc.wantWebLink,
				"af_force_deeplink": "true",
			})
		})
	}
}

func TestActivityPostLinkRejectsInvalidBaseURL(t *testing.T) {
	platform := registerTemporaryPlatform(t, PlatformConfig{BaseURL: "://missing-scheme", URLScheme: "example://"})
	if _, err := NewActivityPostLink(platform, "abc").Build(); err == nil {
		t.Fatal("Build() error = nil, want invalid base URL error")
	}
}
