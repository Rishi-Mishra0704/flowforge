package models

type Node struct {
	ID    int    `json:"id"`
	Label string `json:"value"`
	Type  string `json:"type"`
}
