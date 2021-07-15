package test

import (
	"testing"
	"tinyurl/api"
	"tinyurl/services"
)

var (
	ts        services.TinyService
	shortcode = "clickbait"
	url       = "www.google.com"
)

// create shortcode
func TestCreateShortcode(t *testing.T) {
	got := ts.Store(api.TinyStore{
		Url:       url,
		Shortcode: shortcode,
	})
	want := 200
	if got.HttpCode != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

// check if shortcode stats valid
func TestShortcodeStats(t *testing.T) {
	got := ts.Stats(shortcode)
	want := 200
	if got.HttpCode != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

// check if stats didn't count any redirect
func TestShortcodeStatsWithNoRedirect(t *testing.T) {
	got := ts.Stats(shortcode)
	want := 0
	if got.RedirectCount > want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

// get the shortcode
func TestGetShortcode(t *testing.T) {
	got := ts.Get(shortcode)
	want := 302
	if got.HttpCode != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

// check if stats redirect is increase
func TestShortcodeStatsWithRedirect(t *testing.T) {
	got := ts.Stats(shortcode)
	want := 1
	if got.RedirectCount < want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
