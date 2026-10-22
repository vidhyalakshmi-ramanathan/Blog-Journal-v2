package entities

import "github.com/golang-jwt/jwt/v5"

type SignupData struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	Email       string `json:"email"`
	Phonenumber string `json:"phonenumber"`
}

type PostResponse struct {
	ID       uint   `json:"id"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	CreateAt string `json:"createdAt"`
	AuthorID uint   `json:"authorid"`
}

type LoginData struct {
	Identifier string `json:"identifier"`
	Password   string `json:"password"`
}

type PostData struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type ClaimsData struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

type ApiResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}
