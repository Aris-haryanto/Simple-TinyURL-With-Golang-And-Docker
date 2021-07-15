package test

import (
	"testing"
	"tinyurl/api"
	"tinyurl/services"
)

var (
	ts        services.TinyService
	shortcode = "20ACad"
	url       = "www.google.com"
)

func TestCreateShortcodeWithoutURL(t *testing.T) {
	got := ts.Store(api.TinyStore{
		Shortcode: shortcode,
	})
	want := 400
	if got.HttpCode != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestCreateShortcodeWithInvalidText(t *testing.T) {
	got := ts.Store(api.TinyStore{
		Url:       url,
		Shortcode: "#dT4&@324",
	})
	want := 422
	if got.HttpCode != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

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

// create shortcode
func TestCreateTinyWithoutShortcode(t *testing.T) {
	got := ts.Store(api.TinyStore{
		Url: url,
	})
	want := 200
	if got.HttpCode != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestCreateDuplicateShortcode(t *testing.T) {
	got := ts.Store(api.TinyStore{
		Url:       url,
		Shortcode: shortcode,
	})
	want := 409
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
