package api

type Response struct {
	HttpCode    int         `json:"code"`
	Description string      `json:"description,omitempty"`
	Data        interface{} `json:"data"`
}
