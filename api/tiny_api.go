package api

type TinyStore struct {
	Url           string `json:"url"`
	Shortcode     string `json:"shortcode"`
	StartDate     string `json:"startDate"`
	LastSeenDate  string `json:"lastSeenDate"`
	RedirectCount int    `json:"redirectCount"`
}
