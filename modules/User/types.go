package User

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type DbUser struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	IsClaimed bool   `json:"is_claimed"`
}

type Error struct {
	Message string `json:"message"`
}
