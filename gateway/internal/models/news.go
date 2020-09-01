package models

type News struct {
	ID     uint   `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
	UserID uint   `json:"-"`
	User   *User  `json:"user"`
}

type User struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}
