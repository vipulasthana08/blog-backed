package structure

import "time"

type NewBlog struct {
	Date        time.Time `json:"date"`
	Title       string    `json:"title"`
	Description string    `json:"desc"`
}
type Error struct {
	Message string `json:"Message,omitempty"`
	Op      string `json:"Operation,omitempty"`
	Data    string `json:"info,omitempty"`
}
type ResBlog struct {
	Title       string `json:"title"`
	Description string `json:"desc"`
}

func (e Error) Error() string {
	return e.Message
}
