package models

import "time"

type Film struct {
	ID            int                `json:"id"`
	Title         string             `json:"title" form:"title" gorm:"type: varchar(255)"`
	ThumbnailFilm string             `json:"thumbnailfilm" form:"thumbnailfilm" gorm:"type: varchar(255)"`
	Year          string             `json:"year" form:"year" gorm:"type: int"`
	CategoryID    int                `json:"category_id" form:"category_id"`
	Category      CategoriesResponse `json:"category" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Description   string             `json:"description" gorm:"type:text" form:"description"`
	LinkFilm      string             `json:"linkfilm" form:"linkfilm" gorm:"type: varchar(255)"`
	CreatedAt     time.Time          `json:"-"`
	UpdatedAt     time.Time          `json:"-"`
}

type FilmResponse struct {
	ID            int                `json:"id"`
	Title         string             `json:"title"`
	ThumbnailFilm string             `json:"thumbnailfilm"`
	Year          string             `json:"year"`
	CategoryID    int                `json:"category_id"`
	Category      CategoriesResponse `json:"category" form:"category"`
	Description   string             `json:"description"`
	LinkFilm      string             `json:"linkfilm" form:"linkfilm" gorm:"type: varchar(255)"`
}

type FilmEpisodeResponse struct {
	ID            int                `json:"id"`
	Title         string             `json:"title"`
	ThumbnailFilm string             `json:"thumbnailfilm"`
	Year          string             `json:"year"`
	CategoryID    int                `json:"category_id"`
	Category      CategoriesResponse `json:"category"`
}

type FilmCategoryResponse struct {
	ID            int                `json:"id"`
	Title         string             `json:"title"`
	ThumbnailFilm string             `json:"thumbnailfilm"`
	Year          string             `json:"year"`
	CategoryID    int                `json:"category_id"`
	Category      CategoriesResponse `json:"category"`
	Description   string             `json:"description"`
}

func (FilmEpisodeResponse) TableName() string {
	return "films"
}

func (FilmResponse) TableName() string {
	return "films"
}

func (FilmCategoryResponse) TableName() string {
	return "films"
}
