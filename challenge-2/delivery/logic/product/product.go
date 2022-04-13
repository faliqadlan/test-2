package product

import (
	"errors"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/gommon/log"
)

type Logic struct{}

func New() *Logic {
	return &Logic{}
}

func (l *Logic) ValidationStruct(req Req) error {
	var v = validator.New()
	if err := v.Struct(req); err != nil {
		log.Warn(err)
		switch {
		case strings.Contains(err.Error(), "Name"):
			err = errors.New("invalid name format")
		case strings.Contains(err.Error(), "Price"):
			err = errors.New("invalid price format")
		case strings.Contains(err.Error(), "Stock"):
			err = errors.New("invalid stock format")
		default:
			err = errors.New("invalid input")
		}
		return err
	}
	return nil
}
