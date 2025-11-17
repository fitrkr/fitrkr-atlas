package view

type Equipment struct {
	Name        string        `json:"name"`
	Description string        `json:"description"`
	Type        string        `json:"type"`
	Attachment  []*Attachment `json:"attachment"`
}

type Attachment struct {
	Name string `json:"name"`
	Type string `json:"type"`
}
