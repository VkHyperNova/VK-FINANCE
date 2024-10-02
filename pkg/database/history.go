package database

/* Database Functions */

type HistoryItem struct {
	DATE    string  `json:"date"`
	TIME    string  `json:"time"`
	COMMENT string  `json:"comment"`
	VALUE   float64 `json:"value"`
}

type History struct {
	History []HistoryItem `json:"history"` // Slice containing multiple Quote instances.
}
