package events

// PageMove event scheme struct
type PageMove struct {
	baseEvent
	Data struct {
		baseData
		PriorState     struct {
			PageTitle     string `json:"page_title"`
			PageNamespace int    `json:"page_namespace"`
			RevID         int    `json:"rev_id"`
		} `json:"prior_state"`
		Comment       string `json:"comment"`
		Parsedcomment string `json:"parsedcomment"`
	}
}
