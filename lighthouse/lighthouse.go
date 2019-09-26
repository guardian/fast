package lighthouse

import "encoding/json"

// Category - Type for top-level Lighthouse reports
type Category struct {
	Score float64 `json:"score"`
}

// Audit - information on a specific audit such as TTI
type Audit struct {
	NumericValue float64 `json:"numericValue"`
	Description  string  `json:"description"`
}

// SummaryItem - resource metadata, e.g. for script bytes loaded
type SummaryItem struct {
	ResourceType string  `json:"resourceType"`
	Size         float64 `json:"size"`
}

// Audits - lighthouse audits we care about
type Audits struct {
	Interactive             Audit `json:"interactive"`
	MainthreadWorkBreakdown Audit `json:"mainthread-work-breakdown"`
	ResourceSummary         struct {
		Details struct {
			Items []SummaryItem `json:"items"`
		} `json:"details"`
	} `json:"resource-summary"`
}

// Lighthouse output
type Lighthouse struct {
	Audits     Audits `json:"audits"`
	Categories struct {
		Performance Category `json:"performance"`
	} `json:"categories"`
}

// Unmarshal - convert to JSON string
func (lh *Lighthouse) Unmarshal(data []byte) error {
	return json.Unmarshal(data, lh)
}
