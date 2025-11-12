package view

type Category struct {
	ID          int           `json:"id"`
	Name        string        `json:"name"`
	Subcategory []Subcategory `json:"subcategory"`
}

type Subcategory struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
