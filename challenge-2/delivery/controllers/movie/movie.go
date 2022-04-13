package movie

import (
	"be/api/aws/s3"
	logic "be/delivery/logic/movie"
	"be/delivery/templates"
	"be/repository/movie"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type Controller struct {
	r  movie.Movie
	s3 s3.TaskS3M
	l  logic.Movie
}

func New(r movie.Movie, s3 s3.TaskS3M, l logic.Movie) *Controller {
	return &Controller{
		r:  r,
		s3: s3,
		l:  l,
	}
}

func (cont *Controller) Create() echo.HandlerFunc {
	return func(c echo.Context) error {

		var req logic.Req

		if err := c.Bind(&req); err != nil {
			switch {
			case strings.Contains(err.Error(), "title"):
				err = errors.New("invalid title format")
			case strings.Contains(err.Error(), "description"):
				err = errors.New("invalid description format")
			case strings.Contains(err.Error(), "duration"):
				err = errors.New("invalid duration format")
			case strings.Contains(err.Error(), "artist"):
				err = errors.New("invalid artist format")
			case strings.Contains(err.Error(), "genres"):
				err = errors.New("invalid genres format")
			default:
				err = errors.New("invalid input")
			}
			return c.JSON(http.StatusBadRequest, templates.BadRequest(nil, err.Error(), nil))
		}

		// validation struct

		if err := cont.l.ValidationStruct(req); err != nil {
			log.Warn(err)
			return c.JSON(http.StatusBadRequest, templates.BadRequest(nil, err.Error(), nil))
		}

		// aws s3

		file, err := c.FormFile("file")
		if err != nil {
			log.Warn(err)
		}
		if err == nil {
			link, err := cont.s3.UploadFileToS3(*file)
			if err != nil {
				log.Warn(err)
				return c.JSON(http.StatusInternalServerError, templates.InternalServerError(nil, errors.New("there's some problem is server"), nil))
			}

			req.Image = link
		}

		// create product

		_, err = cont.r.Create(*req.ToMovie())

		if err != nil {
			log.Warn(err)
			switch {
			default:
				err = errors.New("there's some problem is server")
			}
			return c.JSON(http.StatusInternalServerError, templates.InternalServerError(nil, err.Error(), nil))
		}

		return c.JSON(http.StatusCreated, templates.Success(http.StatusCreated, "success add movie", nil))
	}
}

func (cont *Controller) Update() echo.HandlerFunc {
	return func(c echo.Context) error {

		var movie_uid = c.QueryParam("movie_uid")
		var req logic.Req

		if err := c.Bind(&req); err != nil {
			switch {
			case strings.Contains(err.Error(), "title"):
				err = errors.New("invalid title format")
			case strings.Contains(err.Error(), "description"):
				err = errors.New("invalid description format")
			case strings.Contains(err.Error(), "duration"):
				err = errors.New("invalid duration format")
			case strings.Contains(err.Error(), "artist"):
				err = errors.New("invalid artist format")
			case strings.Contains(err.Error(), "genres"):
				err = errors.New("invalid genres format")
			default:
				err = errors.New("invalid input")
			}
			return c.JSON(http.StatusBadRequest, templates.BadRequest(nil, err.Error(), nil))
		}

		// aws s3

		file, err := c.FormFile("file")
		if err != nil {
			log.Warn(err)
		}
		if err == nil {
			res1, err := cont.r.Get("", "", "", "", movie_uid, 1, 1)
			if err != nil {
				log.Warn(err)
				return c.JSON(http.StatusInternalServerError, templates.InternalServerError(nil, errors.New("there's some problem is server"), nil))
			}
			if res1.Responses[0].Image != "https://www.teralogistics.com/wp-content/uploads/2020/12/default.png" {
				var nameFile = res1.Responses[0].Image

				nameFile = strings.Replace(nameFile, "https://karen-givi-bucket.s3.ap-southeast-1.amazonaws.com/", "", -1)

				var res = cont.s3.UpdateFileS3(nameFile, *file)
				log.Info(res)
				if res != "success" {
					log.Warn(res)
					return c.JSON(http.StatusInternalServerError, templates.InternalServerError(nil, errors.New("there's some problem is server"), nil))
				}
			} else {
				var link, err = cont.s3.UploadFileToS3(*file)
				if err != nil {
					log.Warn(err)
					return c.JSON(http.StatusInternalServerError, templates.InternalServerError(nil, errors.New("there's some problem is server"), nil))
				}

				req.Image = link
			}
		}

		// create product

		_, err = cont.r.Update(movie_uid, *req.ToMovie())

		if err != nil {
			log.Warn(err)
			switch {
			default:
				err = errors.New("there's some problem is server")
			}
			return c.JSON(http.StatusInternalServerError, templates.InternalServerError(nil, err.Error(), nil))
		}

		return c.JSON(http.StatusAccepted, templates.Success(http.StatusAccepted, "success Update movie", nil))
	}
}

func (cont *Controller) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {

		var movie_uid = c.QueryParam("movie_uid")

		// aws s3

		res1, err := cont.r.Get("", "", "", "", movie_uid, 1, 1)
		if err != nil {
			log.Error(err)
		}

		if res1.Responses[0].Image != "https://www.teralogistics.com/wp-content/uploads/2020/12/default.png" {

			var nameFile = res1.Responses[0].Image

			nameFile = strings.Replace(nameFile, "https://karen-givi-bucket.s3.ap-southeast-1.amazonaws.com/", "", -1)
			res := cont.s3.DeleteFileS3(nameFile)
			log.Info(res)
			if res != "success" {
				log.Warn(res)
			}
		}

		// create product

		_, err = cont.r.Delete(movie_uid)

		if err != nil {
			log.Warn(err)
			switch {
			default:
				err = errors.New("there's some problem is server")
			}
			return c.JSON(http.StatusInternalServerError, templates.InternalServerError(nil, err.Error(), nil))
		}

		return c.JSON(http.StatusAccepted, templates.Success(http.StatusAccepted, "success delete movie", nil))
	}
}

func (cont *Controller) Get() echo.HandlerFunc {
	return func(c echo.Context) error {

		var title = c.QueryParam("title")
		var description = c.QueryParam("description")
		var artist = c.QueryParam("artist")
		var genres = c.QueryParam("genres")
		var movie_uid = c.QueryParam("movie_uid")
		var limitString = c.QueryParam("limit")
		limit, err := strconv.Atoi(limitString)
		if err != nil && limitString != "" {
			log.Warn(err)
			return c.JSON(http.StatusBadRequest, templates.BadRequest(nil, "invalid limit format", nil))
		}
		var pageString = c.QueryParam("page")
		page, err := strconv.Atoi(pageString)
		if err != nil && pageString != "" {
			log.Warn(err)
			return c.JSON(http.StatusBadRequest, templates.BadRequest(nil, "invalid page format", nil))
		}

		if limitString == "" {
			limit = 0
		}
		if pageString == "" {
			page = 1
		}
		// log.Info(title)
		res, err := cont.r.Get(title, description, artist, genres, movie_uid, limit, page)

		if err != nil {
			log.Warn(err)
			switch {
			default:
				err = errors.New("there's some problem is server")
			}
			return c.JSON(http.StatusInternalServerError, templates.InternalServerError(nil, err.Error(), nil))
		}

		return c.JSON(http.StatusOK, templates.Success(http.StatusOK, "success get movie", res))
	}
}
