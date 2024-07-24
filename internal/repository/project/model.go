package project

type Project struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Boards      []string `json:"tables"` // TODO: use ref boards UUIDs in the future
	Text        string   `json:"text"`
	Links       []string `json:"links"`
}
