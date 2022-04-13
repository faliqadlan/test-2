package movie

import (
	"be/entities"
	"strconv"

	"github.com/google/uuid"
	"github.com/labstack/gommon/log"
	"github.com/lithammer/shortuuid"
	"gorm.io/gorm"
)

type Repo struct {
	db *gorm.DB
}

func New(db *gorm.DB) *Repo {
	return &Repo{
		db: db,
	}
}

func (r *Repo) Create(req entities.Movie) (entities.Movie, error) {

	var uid string

	for {
		uid = strconv.Itoa(int(uuid.New().ID()))
		var res = r.db.Model(&entities.Movie{}).Where("movie_uid = ?", uid).Scan(&entities.Movie{})
		if res.RowsAffected == 0 {
			break
		}
	}

	req.Movie_uid = uid

	var res = r.db.Model(&entities.Movie{}).Create(&req)

	if res.Error != nil {
		return entities.Movie{}, res.Error
	}

	return req, nil
}

func (r *Repo) Delete(movie_uid string) (entities.Movie, error) {

	var resInit entities.Movie

	var res = r.db.Model(&entities.Movie{}).Where("movie_uid = ?", movie_uid).Delete(&resInit)

	if res.RowsAffected == 0 {
		return entities.Movie{}, gorm.ErrRecordNotFound
	}

	return resInit, nil
}

func (r *Repo) Update(movie_uid string, req entities.Movie) (entities.Movie, error) {

	var res = r.db.Model(&entities.Movie{}).Where("movie_uid = ?", movie_uid).Updates(entities.Movie{
		Title:       req.Title,
		Description: req.Description,
		Duration:    req.Duration,
		Artist:      req.Artist,
		Genres:      req.Genres,
		Image:       req.Image,
	})

	if res.RowsAffected == 0 {
		return entities.Movie{}, gorm.ErrRecordNotFound
	}

	return req, nil
}

func (r *Repo) Get(title, description, artist, genres, movie_uid string, limit, page int) (GetResponses, error) {

	switch {
	case title != "":
		title = "title LIKE '%" + title + "%'"
	case title == "":
		title = "title != '" + shortuuid.New() + "'"
	}

	switch {
	case description != "":
		description = "description LIKE '%" + description + "%'"
	case description == "":
		description = "description != '" + shortuuid.New() + "'"
	}

	switch {
	case artist != "":
		artist = "artist LIKE '%" + artist + "%'"
	case artist == "":
		artist = "artist != '" + shortuuid.New() + "'"
	}

	switch {
	case genres != "":
		genres = "genres LIKE '%" + genres + "%'"
	case genres == "":
		genres = "genres != '" + shortuuid.New() + "'"
	}
	log.Info(movie_uid)
	switch {
	case movie_uid != "":
		movie_uid = "movie_uid = '" + movie_uid + "'"
	case movie_uid == "":
		movie_uid = "movie_uid != '" + shortuuid.New() + "'"
	}
	// log.Info(title)
	var condition = title + " AND " + description + " AND " + artist + " AND " + genres + " AND " + movie_uid

	var response = GetResponses{}

	var res = r.db.Model(&entities.Movie{}).Where(condition).Select("movie_uid as Movie_uid, title as Title, description as Description, duration as Duration, artist as Artist, genres as Genres, image as Image").Limit(limit).Offset((page - 1) * limit).Scan(&response.Responses)

	if res.Error != nil {
		return GetResponses{}, res.Error
	}

	return response, nil
}
