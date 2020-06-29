package core

type Resource struct {
	Identifier    string                 `json:"id,omitempty"`
	Type          string                 `json:"type"`
	Attributes    map[string]interface{} `json:"attributes,omitempty"`
	Meta          map[string]interface{} `json:"meta,omitempty"`
	Relationships map[string]Document    `json:"relationships,omitempty"`
	// Links         map[string]string      `json:"links,omitempty"`
	Links LinksObject `json:"links,omitempty"`
}

func (r Resource) Identify() Resource {

	return Resource{
		Type:       r.Type,
		Identifier: r.Identifier,
		Meta:       r.Meta,
	}
}
