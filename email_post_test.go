package deeplink

import (
	"net/url"
	"strings"
	"testing"
)

// Both EDM post links share one builder; these cases cover what the
// per-link tests do not.
var emailPostLinks = map[string]func(Platform, string) Deeplink{
	"activity": NewActivityPostLink,
	"tip":      NewTipReceivedLink,
}

func TestEmailPostLinkRejectsUnknownPlatform(t *testing.T) {
	for name, newLink := range emailPostLinks {
		if _, err := newLink(Platform(42), "abc").Build(); err == nil {
			t.Errorf("%s: Build() error = nil for an unregistered platform", name)
		}
	}
}

func TestEmailPostLinkRejectsEmptyPostID(t *testing.T) {
	for name, newLink := range emailPostLinks {
		if _, err := newLink(PlatformApen, "").Build(); err == nil {
			t.Errorf("%s: Build() error = nil for an empty post id", name)
		}
	}
}

// A post id with reserved characters stays inside the path of both links.
func TestEmailPostLinkEscapesPostID(t *testing.T) {
	const id = "abc?x=1#frag/../y"
	raw, err := NewTipReceivedLink(PlatformApen, id).Build()
	if err != nil {
		t.Fatalf("Build() error: %v", err)
	}
	u, err := url.Parse(raw)
	if err != nil {
		t.Fatalf("url.Parse(%q) error: %v", raw, err)
	}
	q := u.Query()
	if len(q["x"]) != 0 || u.Fragment != "" {
		t.Errorf("post id leaked into the OneLink query or fragment: %s", raw)
	}
	want := "apen://posts/" + url.PathEscape(id)
	if q.Get("af_dp") != want || q.Get("deep_link_value") != want {
		t.Errorf("af_dp = %q, deep_link_value = %q, want %q", q.Get("af_dp"), q.Get("deep_link_value"), want)
	}
	if web := q.Get("af_web_dp"); !strings.HasSuffix(web, "/post/"+url.PathEscape(id)) {
		t.Errorf("af_web_dp = %q", web)
	}
}

func TestEmailPostLinkWithoutWebBaseURL(t *testing.T) {
	platform := registerTemporaryPlatform(t, PlatformConfig{BaseURL: "https://example.onelink.me/x", URLScheme: "example://"})
	raw, err := NewTipReceivedLink(platform, "abc").Build()
	if err != nil {
		t.Fatalf("Build() error: %v", err)
	}
	u, err := url.Parse(raw)
	if err != nil {
		t.Fatalf("url.Parse(%q) error: %v", raw, err)
	}
	if _, present := u.Query()["af_web_dp"]; present {
		t.Errorf("af_web_dp present without a WebBaseURL: %s", raw)
	}
	if u.Query().Get("af_dp") != "example://posts/abc" {
		t.Errorf("af_dp = %q", u.Query().Get("af_dp"))
	}
}
