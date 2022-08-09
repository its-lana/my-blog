package web

type PostUpdateRequest struct {
	Id   int    `validate:"required"`
	Title string `validate:"required,max=200,min=1" json:"title"`
	Category string `validate:"required,min=1,max=100" json:"category"`
	Content string `validate:"required,min=1,max=5000" json:"content"`
}
