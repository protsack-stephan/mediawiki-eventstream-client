package eventstream

import "encoding/json"

// PageDelete event scheme struct
type PageDelete struct {
	baseSchema
	Data struct {
		baseData
		RevCount      int    `json:"rev_count"`
		Comment       string `json:"comment"`
		Parsedcomment string `json:"parsedcomment"`
	}
}

func (pd *PageDelete) unmarshal(evt *Event) error {
	pd.ID = evt.ID
	return json.Unmarshal(evt.Data, &pd.Data)
}
