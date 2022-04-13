package movie

import (
	"testing"

	"github.com/labstack/gommon/log"
	"github.com/stretchr/testify/assert"
)

func TestValidationStruct(t *testing.T) {

	t.Run("validator title", func(t *testing.T) {
		var req = Req{
			Title:       "",
			Description: "test",
			Duration:    "test",
			Artist:      "test",
			Genres:      "test",
		}
		var l = New()
		var err = l.ValidationStruct(req)
		assert.NotNil(t, err)
		log.Info(err)
	})

	t.Run("validator description", func(t *testing.T) {
		var req = Req{
			Title:       "test",
			Description: "",
			Duration:    "test",
			Artist:      "test",
			Genres:      "test",
		}
		var l = New()
		var err = l.ValidationStruct(req)
		assert.NotNil(t, err)
		log.Info(err)
	})

	t.Run("validator duration", func(t *testing.T) {
		var req = Req{
			Title:       "test",
			Description: "test",
			Duration:    "",
			Artist:      "test",
			Genres:      "test",
		}
		var l = New()
		var err = l.ValidationStruct(req)
		assert.NotNil(t, err)
		log.Info(err)
	})

	t.Run("validator artist", func(t *testing.T) {
		var req = Req{
			Title:       "test",
			Description: "test",
			Duration:    "test",
			Artist:      "",
			Genres:      "test",
		}
		var l = New()
		var err = l.ValidationStruct(req)
		assert.NotNil(t, err)
		log.Info(err)
	})

	t.Run("validator genres", func(t *testing.T) {
		var req = Req{
			Title:       "test",
			Description: "test",
			Duration:    "test",
			Artist:      "genres",
			Genres:      "",
		}
		var l = New()
		var err = l.ValidationStruct(req)
		assert.NotNil(t, err)
		log.Info(err)
	})

	t.Run("validator success", func(t *testing.T) {
		var req = Req{
			Title:       "test",
			Description: "test",
			Duration:    "test",
			Artist:      "genres",
			Genres:      "test",
		}
		var l = New()
		var err = l.ValidationStruct(req)
		assert.Nil(t, err)
		log.Info(err)
	})

}
