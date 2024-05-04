package handler

type UserRequest struct {
	Name      string `json:"name" form:"name"`
	Email     string `json:"email" form:"email"`
	Password  string `json:"password" form:"password"`
	Phone     string `json:"phone" form:"phone"`
	Address   string `json:"address" form:"address"`
	StoreName string `json:"store_name" form:"store_name"`
}

type LoginRequest struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}
