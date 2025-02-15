package User

type LoginRequest struct {
	Username    string `json:"username" binding:"required"`
	Password    string `json:"password" binding:"required"`
	IsTimeBased bool   `json:"isTimeBased" binding:"required"`
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

type ClaimRequest struct {
	NewPassword string `json:"newPassword" binding:"required"`
}

type Error struct {
	Message string `json:"message"`
}
