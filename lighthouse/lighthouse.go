package lighthouse

import "encoding/json"

// Category - Type for top-level Lighthouse reports
type Category struct {
	Score float64 `json:"score"`
}

// Lighthouse output
type Lighthouse struct {
	Categories struct {
		Performance Category `json:"performance"`
	} `json:"categories"`
}

func (lh *Lighthouse) Unmarshal(data []byte) error {
	return json.Unmarshal(data, lh)
}
