package models

type Flowchart struct {
	Nodes []Node `json:"nodes"`
	Edges []Edge `json:"edges"`
}
