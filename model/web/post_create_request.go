package web

type PostCreateRequest struct {
	Title string `validate:"required,min=1,max=100" json:"title"`
	Category string `validate:"required,min=1,max=100" json:"category"`
	Content string `validate:"required,min=1,max=5000" json:"content"`
}
