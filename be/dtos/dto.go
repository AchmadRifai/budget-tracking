package dtos

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegRequest struct {
	LoginRequest
	Email    string `json:"email"`
	FullName string `json:"full_name"`
}
