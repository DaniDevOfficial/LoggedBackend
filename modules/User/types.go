package User

type LoginRequest struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	IsTimeBased bool   `json:"isTimeBased"`
}

type DbUser struct {
	Id        string `json:"id"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	IsClaimed bool   `json:"is_claimed"`
}

type LoginResponse struct {
	IsClaimed bool `json:"is_claimed"`
}

type Error struct {
	Message string `json:"message"`
}
