package entity

type Person struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	Age       string `json:"age" binding:"gte=1,lte=130"`
}
type Video struct {
	Title       string `json:"title" binding:"min=2,max=20" validate:"is-cool"`
	Description string `json:"description" binding:"max=100"`
	URL         string `json:"url" binding:"required,url"`
	Author      Person `json:"author" binding:"required"`
}
