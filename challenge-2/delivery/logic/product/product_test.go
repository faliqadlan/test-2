package product

import (
	"testing"

	"github.com/labstack/gommon/log"
	"github.com/stretchr/testify/assert"
)

func TestValidationStruct(t *testing.T) {

	t.Run("validator name", func(t *testing.T) {
		var req = Req{
			Name:        "",
			Price:       "1000",
			Stock:       10,
			Description: "description",
		}
		var l = New()
		var err = l.ValidationStruct(req)
		assert.NotNil(t, err)
		log.Info(err)
	})

	t.Run("validator Price", func(t *testing.T) {
		var req = Req{
			Name:        "name",
			Price:       "",
			Stock:       10,
			Description: "description",
		}
		var l = New()
		var err = l.ValidationStruct(req)
		assert.NotNil(t, err)
		log.Info(err)
	})

	t.Run("validator stock", func(t *testing.T) {
		var req = Req{
			Name:        "name",
			Price:       "1000",
			Stock:       0,
			Description: "description",
		}
		var l = New()
		var err = l.ValidationStruct(req)
		assert.NotNil(t, err)
		log.Info(err)
	})
}
