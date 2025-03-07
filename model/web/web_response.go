package web

type WebResponse struct {
	Code   int         `json:"code"`
	Status bool        `json:"status"`
	Data   interface{} `json:"data"`
}
