package view

type Category struct {
	ID   int      `json:"id"`
	Type string   `json:"type"`
	Name []string `json:"name"`
}
