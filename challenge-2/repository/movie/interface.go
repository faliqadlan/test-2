package movie

import "be/entities"

type Movie interface {
	Create(req entities.Movie) (entities.Movie, error)
	Delete(movie_uid string) (entities.Movie, error)
	Update(movie_uid string, req entities.Movie) (entities.Movie, error)
	Get(title, description, artist, genres, movie_id string, limit, page int) (GetResponses, error)
}
