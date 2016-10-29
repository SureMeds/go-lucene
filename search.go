package lucene

// Search ...
//****************************************************************************//
// This structure represents a stratio lucene filter object										//
//****************************************************************************//
type Search struct {
	Filters []Filter `json:"filter,omitempty"`
	Refresh bool     `json:"refresh,omitempty"`
	Query   Query    `json:"query,omitempty"`
	Sort    []Basis  `json:"sort,omitempty"`
}
