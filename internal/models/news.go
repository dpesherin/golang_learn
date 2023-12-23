package models

type News struct {
	ID          int64    `json:"id"`
	Title       string   `json:"title"`
	Text        string   `json:"text"`
	Attachments []Attach `json:"attachments"`
	Links       []Link   `json:"links"`
}

type Attach struct {
	Path string `json:"path"`
}

type Link struct {
	Path string `json:"path"`
}
