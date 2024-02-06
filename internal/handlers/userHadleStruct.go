package handlers

type userReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
