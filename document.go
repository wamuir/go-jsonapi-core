package core

import (
	"bytes"
	"encoding/json"
	"fmt"
)

var AssertionFailure = fmt.Errorf("type assertion failed")

type Document struct {
	JSONAPI  *Implementation        `json:"jsonapi,omitempty"`
	Data     interface{}            `json:"data,omitempty"`
	Meta     map[string]interface{} `json:"meta,omitempty"`
	Links    LinksObject            `json:"links,omitempty"`
	Included Included               `json:"included,omitempty"`
	Errors   []Error                `json:"errors,omitempty"`
}

func (d *Document) AssertDataType() error {

	j, err := json.Marshal(d.Data)
	if err != nil {
		return err
	}

	var r Resource
	if err = json.Unmarshal(j, &r); err == nil {
		d.Data = r
		return nil
	}

	var c Collection
	if err = json.Unmarshal(j, &c); err == nil {
		d.Data = c
		return nil
	}

	return AssertionFailure
}

func (d *Document) ContentLength() int {

	buf := &bytes.Buffer{}
	encoder := json.NewEncoder(buf)
	encoder.SetEscapeHTML(false)
	encoder.SetIndent("", "\t")
	encoder.Encode(d)

	return buf.Len()
}

func (d *Document) PopIncluded() Included {

	i := d.Included
	d.Included = Included{}

	return i
}

func (d *Document) Version() {

	d.JSONAPI = &Implementation{
		Version: jsonapi_version,
		Meta: map[string]interface{}{
			"links": map[string]interface{}{
				"self": Link{
					Href: "http://jsonapi.org/format/1.0/",
				},
			},
		},
	}

	return
}

func New() Document {
	var d Document
	d.Version()
	return d
}
