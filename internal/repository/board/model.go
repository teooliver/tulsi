package board

type Status struct {
	ID    string   `json:"id"`
	Name  string   `json:"name"`
	Tasks []string `json:"tasks"`
}
