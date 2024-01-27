package dto

type LoginDTO struct {
  Mail string `json:"mail"`
  Password string `json:"password"`
}

type LoginResponse struct {
  Name string `json:"name"`
  Token string `json:"token"`
}
