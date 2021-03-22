package models

type Relation struct {
	Child  string `json:"child"`
	Parent string `json:"parent"`
	Type   Gender `json:"type"`
}
