package core

const jsonapi_version = "1.0" // Version of JSON:API implementation

type Implementation struct {
	Version string                 `json:"version,omitempty"`
	Meta    map[string]interface{} `json:"meta,omitempty"`
}
