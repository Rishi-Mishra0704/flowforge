package models

type Edge struct {
	Source    int    `json:"source"`
	Target    int    `json:"target"`
	Condition string `json:"condition,omitempty"`
}
