package view

type MuscleGroup struct {
	ID          int      `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Muscle      []Muscle `json:"muscle"`
}

type Muscle struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Activation string `json:"activation"`
}
