package auth

import (
	"be/delivery/middlewares"
	"be/delivery/templates"
	"be/repository/auth"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type Controller struct {
	r auth.Auth
}

func New(r auth.Auth) *Controller {
	return &Controller{r: r}
}

func (cont *Controller) Login() echo.HandlerFunc {
	return func(c echo.Context) error {

		var userLogin = Userlogin{}

		if err := c.Bind(&userLogin); err != nil || userLogin.UserName == "" || userLogin.Password == "" {
			return c.JSON(http.StatusBadRequest, templates.BadRequest(nil, "invalid email or password", err))
		}

		res, err := cont.r.Login(userLogin.UserName, userLogin.Password)

		if err != nil {
			log.Warn(err)
			switch err.Error() {
			case "incorrect password":
				err = errors.New("incorrect password")
			case "record not found":
				err = errors.New("account is not found")
			default:
				err = errors.New("there's some problem is server")
			}

			return c.JSON(http.StatusInternalServerError, templates.InternalServerError(nil, err.Error(), nil))
		}

		token, err := middlewares.GenerateToken(res)

		if err != nil {
			log.Warn(err)
			err = errors.New("there's some problem is server")
			return c.JSON(http.StatusNotAcceptable, templates.BadRequest(http.StatusNotAcceptable, err.Error(), nil))
		}

		return c.JSON(http.StatusOK, templates.Success(http.StatusOK, "success login", map[string]interface{}{"token": token}))

	}
}
