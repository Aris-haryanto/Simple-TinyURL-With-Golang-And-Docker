package api

type ResStore struct {
	HttpCode    int    `json:"-"`
	Description string `json:"description,omitempty"`
	Shortcode   string `json:"shortcode,omitempty"`
}
type ResGet struct {
	HttpCode    int    `json:"-"`
	Description string `json:"description,omitempty"`
	Location    string `json:"location,omitempty"`
}
type ResStats struct {
	HttpCode      int    `json:"-"`
	Description   string `json:"description,omitempty"`
	StartDate     string `json:"startDate,omitempty"`
	LastSeenDate  string `json:"lastSeenDate,omitempty"`
	RedirectCount int    `json:"redirectCount"`
}
