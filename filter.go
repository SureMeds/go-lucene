package lucene

// Filter ...
//****************************************************************************//
// This structure represents a stratio lucene filter object										//
//****************************************************************************//
type Filter struct {
	Type         string `json:"type,omitempty"`
	Field        string `json:"field,omitempty"`
	Value        string `json:"value,omitempty"`
	Lower        string `json:"lower,omitempty"`
	Upper        string `json:"upper,omitempty"`
	Latitude     string `json:"latitude,omitempty"`
	Longitude    string `json:"longitude,omitempty"`
	MinLatitude  string `json:"min_latitude,omitempty"`
	MinLongitude string `json:"min_longitude,omitempty"`
	MaxLatitude  string `json:"max_latitude,omitempty"`
	MaxLongitude string `json:"max_longitude,omitempty"`
	MaxDistance  string `json:"max_distance,omitempty"`
	MinDistance  string `json:"min_distance,omitempty"`
	Reverse      bool   `json:"reverse,omitempty"`
	Slop         int    `json:"slop,omitempty"`
}
