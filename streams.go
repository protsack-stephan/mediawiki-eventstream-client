package eventstream

const baseURL = "https://stream.wikimedia.org/v2/stream/"

// All the available streams
const (
	pageDelete     = baseURL + "page-delete"
	pageMove       = baseURL + "page-move"
	revisionCreate = baseURL + "revision-create"
	revisionScore  = baseURL + "revision-score"
)
