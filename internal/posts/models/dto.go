package models

type LoginDTO struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type UserDTO struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Login    string `json:"login"`
	Password string `json:"password"`
	Type     string `json:"type"`
}

type PostDTO struct {
	ID     int    `json:"id"`
	Header string `json:"header"`
	Text   string `json:"text"`
}
