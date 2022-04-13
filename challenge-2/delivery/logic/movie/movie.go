package movie

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
		case strings.Contains(err.Error(), "Title"):
			err = errors.New("invalid title format")
		case strings.Contains(err.Error(), "Description"):
			err = errors.New("invalid description format")
		case strings.Contains(err.Error(), "Duration"):
			err = errors.New("invalid duration format")
		case strings.Contains(err.Error(), "Artist"):
			err = errors.New("invalid artist format")
		case strings.Contains(err.Error(), "Genres"):
			err = errors.New("invalid genres format")
		default:
			err = errors.New("invalid input")
		}
		return err
	}
	return nil
}
