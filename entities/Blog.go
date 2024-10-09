package entities

import "time"

type Blog struct {
	Id       int       `json:"id"`
	Title    string    `json:"title"`
	Author   string    `json:"author"`
	Content  string    `json:"content"`
	CreateAt time.Time `json:"created_at"`
	Image    string    `json:"image"`
}

type FormBlog struct {
	Title   string `json:"title"`
	Author  string `json:"author"`
	Content string `json:"content"`
	Image   string `json:"image"`
}
