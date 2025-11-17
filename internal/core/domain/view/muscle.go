package view

type MuscleGroup struct {
	Group  string   `json:"group"`
	Muscle []Muscle `json:"muscle"`
}

type Muscle struct {
	Name       string `json:"name"`
	Activation string `json:"activation"`
}
