package web

type PostResponse struct {
	Id   int    `json:"id"`
	Title string `json:"title"`
	Category string `json:"category"`
	Content string `json:"content"`
}
