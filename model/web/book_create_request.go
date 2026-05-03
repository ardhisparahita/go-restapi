package web

type BookCreateRequest struct {
	Title  string `validate:"required,max=255,min=1" json:"title"`
	Author string `validate:"max=100,min=1" json:"author"`
}
