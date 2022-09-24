package models

type Category struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	
}

type CategoriesResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (CategoriesResponse) TableName() string {
	return "categories"
}
