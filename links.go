package core

type LinksObject map[string]interface{}

type Link struct {
	Href string                 `json:"href,omitempty"`
	Meta map[string]interface{} `json:"meta,omitempty"`
}
