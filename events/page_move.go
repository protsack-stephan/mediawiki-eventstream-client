package events

// PageMove event scheme struct
type PageMove struct {
	baseEvent
	Data struct {
		Schema         string    `json:"$schema"`
		Meta           Meta      `json:"meta"`
		Database       string    `json:"database"`
		Performer      Performer `json:"performer"`
		PageID         int       `json:"page_id"`
		PageTitle      string    `json:"page_title"`
		PageNamespace  int       `json:"page_namespace"`
		PageIsRedirect bool      `json:"page_is_redirect"`
		RevID          int       `json:"rev_id"`
		PriorState     struct {
			PageTitle     string `json:"page_title"`
			PageNamespace int    `json:"page_namespace"`
			RevID         int    `json:"rev_id"`
		} `json:"prior_state"`
		Comment       string `json:"comment"`
		Parsedcomment string `json:"parsedcomment"`
	}
}
