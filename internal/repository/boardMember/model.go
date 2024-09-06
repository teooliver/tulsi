package boardmember

type BoardMember struct {
	UserID  string `json:"user_id"`  //FK
	BoardID string `json:"board_id"` //FK
}
