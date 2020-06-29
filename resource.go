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

func (resource Resource) Identify() Resource {

	return Resource{
		Type:       resource.Type,
		Identifier: resource.Identifier,
		Meta:       resource.Meta,
	}
}
