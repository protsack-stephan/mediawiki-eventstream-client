package eventstream

// PageDelete event scheme struct
type PageDelete struct {
	baseEvent
	Data struct {
		baseData
		RevCount      int    `json:"rev_count"`
		Comment       string `json:"comment"`
		Parsedcomment string `json:"parsedcomment"`
	}
}
