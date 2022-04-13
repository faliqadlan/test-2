package movie

import "be/entities"

type Req struct {
	Title 	 string `json:"title" form:"title" validate:"required"`
	Description string `json:"description" form:"description" validate:"required"`
	Duration string `json:"duration" form:"duration" validate:"required"`
	Artist string `json:"artist" form:"artist" validate:"required"`
	Genres string `json:"genres" form:"genres" validate:"required"`
	Image string `json:"image" form:"image"`
}

func (r *Req) ToMovie() *entities.Movie {
	return &entities.Movie{
		Title:       r.Title,
		Description: r.Description,
		Duration:    r.Duration,
		Artist:      r.Artist,
		Genres:      r.Genres,
		Image:       r.Image,
	}
}
