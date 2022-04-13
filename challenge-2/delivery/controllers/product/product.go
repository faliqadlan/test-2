package product

import (
	"be/api/aws/s3"
	logic "be/delivery/logic/product"
	"be/delivery/middlewares"
	"be/delivery/templates"
	"be/repository/product"
	"errors"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type Controller struct {
	r  product.Product
	s3 s3.TaskS3M
	l  logic.Product
}

func New(r product.Product, s3 s3.TaskS3M, l logic.Product) *Controller {
	return &Controller{
		r:  r,
		s3: s3,
		l:  l,
	}
}

func (cont *Controller) Create() echo.HandlerFunc {
	return func(c echo.Context) error {

		var user_uid = middlewares.ExtractTokenUid(c)
		var req logic.Req

		if err := c.Bind(&req); err != nil {
			switch {
			case strings.Contains(err.Error(), "name"):
				err = errors.New("invalid name format")
			case strings.Contains(err.Error(), "price"):
				err = errors.New("invalid price format")
			case strings.Contains(err.Error(), "stock"):
				err = errors.New("invalid stock format")
			case strings.Contains(err.Error(), "description"):
				err = errors.New("invalid description format")
			case strings.Contains(err.Error(), "strconv.ParseInt"):
				err = errors.New("invalid input stock")
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

		_, err = cont.r.Create(user_uid, *req.ToProduct())

		if err != nil {
			log.Warn(err)
			switch {
			default:
				err = errors.New("there's some problem is server")
			}
			return c.JSON(http.StatusInternalServerError, templates.InternalServerError(nil, err.Error(), nil))
		}

		return c.JSON(http.StatusCreated, templates.Success(http.StatusCreated, "success add product", nil))
	}
}

func (cont *Controller) Update() echo.HandlerFunc {
	return func(c echo.Context) error {

		var user_uid = middlewares.ExtractTokenUid(c)
		var product_uid = c.QueryParam("product_uid")
		var req logic.Req

		if err := c.Bind(&req); err != nil {
			switch {
			case strings.Contains(err.Error(), "name"):
				err = errors.New("invalid name format")
			case strings.Contains(err.Error(), "price"):
				err = errors.New("invalid price format")
			case strings.Contains(err.Error(), "stock"):
				err = errors.New("invalid stock format")
			case strings.Contains(err.Error(), "description"):
				err = errors.New("invalid description format")
			case strings.Contains(err.Error(), "strconv.ParseInt"):
				err = errors.New("invalid input stock")
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
			res1, err := cont.r.Get(user_uid, product_uid)
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

		_, err = cont.r.Update(product_uid, *req.ToProduct())

		if err != nil {
			log.Warn(err)
			switch {
			default:
				err = errors.New("there's some problem is server")
			}
			return c.JSON(http.StatusInternalServerError, templates.InternalServerError(nil, err.Error(), nil))
		}

		return c.JSON(http.StatusAccepted, templates.Success(http.StatusAccepted, "success Update product", nil))
	}
}

func (cont *Controller) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {

		var user_uid = middlewares.ExtractTokenUid(c)
		var product_uid = c.QueryParam("product_uid")

		// aws s3

		res1, err := cont.r.Get(user_uid, product_uid)
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

		_, err = cont.r.Delete(product_uid)

		if err != nil {
			log.Warn(err)
			switch {
			default:
				err = errors.New("there's some problem is server")
			}
			return c.JSON(http.StatusInternalServerError, templates.InternalServerError(nil, err.Error(), nil))
		}

		return c.JSON(http.StatusAccepted, templates.Success(http.StatusAccepted, "success delete product", nil))
	}
}

func (cont *Controller) Get() echo.HandlerFunc {
	return func(c echo.Context) error {

		var user_uid = middlewares.ExtractTokenUid(c)
		var product_uid = c.QueryParam("product_uid")
		var all = c.QueryParam("all")

		if all != "" {
			user_uid = ""
		}
		// get product
		// log.Info(user_uid)
		// log.Info(product_uid)
		res, err := cont.r.Get(user_uid, product_uid)

		if err != nil {
			log.Warn(err)
			switch {
			default:
				err = errors.New("there's some problem is server")
			}
			return c.JSON(http.StatusInternalServerError, templates.InternalServerError(nil, err.Error(), nil))
		}

		return c.JSON(http.StatusOK, templates.Success(http.StatusOK, "success get product", res))
	}
}
