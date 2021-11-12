package models

type User struct {
	ID   string `json:"userID"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Post struct {
	ID      string `json:"postID"`
	UserID  string `json:"userID"`
	Title   string `json:"title"`
	Message string `json:"message"`
}

type PostRead struct {
	ID      string `json:"postID"`
	Title   string `json:"title"`
	Message string `json:"message"`
}

type Read struct {
	User     User
	PostRead PostRead
}

type UserPosts struct {
	User  User
	Posts []PostRead
}
