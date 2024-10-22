package projectmembers

type ProjectMembers struct {
	UserID    string `json:"user_id"`    //FK
	ProjectID string `json:"project_id"` //FK
}
