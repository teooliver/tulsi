package checklist

type CheckList struct {
	ID   string `json:"id"` //PK
	Name string `json:"name"`
	// UserID string   `json:user_id` //FK
	// IsPublic bool `json:"is_public"`
	// CreatedDate string `json:"created_date"` // DATE type in go?
}
