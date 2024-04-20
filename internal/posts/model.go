package posts

type User struct {
	ID       int     `json:"id"`
	Name     *string `json:"name"`
	Login    string  `json:"login"`
	Password string  `json:"password"`
	Type     string  `json:"type"`
}

type ID struct {
	ID int `json:"id"`
}

type Post struct {
	ID     int    `json:"id"`
	Header string `json:"header"`
	Text   string `json:"text"`
}
