package user

import (
	"be/api/aws/s3"
	logic "be/delivery/logic/user"
	"be/delivery/templates"
	"be/repository/user"
	"errors"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type Controller struct {
	r  user.User
	s3 s3.TaskS3M
	l  logic.User
}

func New(r user.User, s3 s3.TaskS3M, l logic.User) *Controller {
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
			case strings.Contains(err.Error(), "userName"):
				err = errors.New("invalid user name formar")
			case strings.Contains(err.Error(), "email"):
				err = errors.New("invalid email format")
			case strings.Contains(err.Error(), "password"):
				err = errors.New("invalid password format")
			case strings.Contains(err.Error(), "name"):
				err = errors.New("invalid name format")
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

		_, err = cont.r.Create(*req.ToUser())

		if err != nil {
			log.Warn(err)
			switch {
			case err.Error() == errors.New("user name is already exist").Error():
				err = errors.New("user name is already exist")
			case err.Error() == errors.New("email is already exist").Error():
				err = errors.New("email is already exist")
			default:
				err = errors.New("there's some problem is server")
			}
			return c.JSON(http.StatusInternalServerError, templates.InternalServerError(nil, err.Error(), nil))
		}

		return c.JSON(http.StatusCreated, templates.Success(http.StatusCreated, "success add user", nil))
	}
}
