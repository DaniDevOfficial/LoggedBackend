package jwt

type JWTUser struct {
	Username string
	UserId   string
}

type JWTPayload struct {
	UserId       string
	Username     string
	IsClaimToken bool
	Exp          int64
}

type JWTTokenResponse struct {
	Token string `json:"token"`
}
