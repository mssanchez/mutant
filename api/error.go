package api

// Error struct returned by API
type Error struct {
	Cause      string            `json:"cause,omitempty"`
	StatusCode int               `json:"status_code"`
	Context    map[string]string `json:"context,omitempty"`
}
