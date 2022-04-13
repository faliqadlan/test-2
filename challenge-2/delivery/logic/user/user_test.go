package user

import (
	"testing"

	"github.com/labstack/gommon/log"
	"github.com/stretchr/testify/assert"
)

func TestValidationStruct(t *testing.T) {
	t.Run("validator userName", func(t *testing.T) {
		var req = Req{
			UserName: "",
			Email:    "email",
			Password: "password",
			Name:     "name",
		}
		var l = New()
		var err = l.ValidationStruct(req)
		assert.NotNil(t, err)
		log.Info(err)
	})

	t.Run("validator email", func(t *testing.T) {
		var req = Req{
			UserName: "userName",
			Email:    "",
			Password: "password",
			Name:     "name",
		}
		var l = New()
		var err = l.ValidationStruct(req)
		assert.NotNil(t, err)
		log.Info(err)
	})

	t.Run("validator password", func(t *testing.T) {
		var req = Req{
			UserName: "userName",
			Email:    "email",
			Password: "",
			Name:     "name",
		}
		var l = New()
		var err = l.ValidationStruct(req)
		assert.NotNil(t, err)
		log.Info(err)
	})

	t.Run("validator name", func(t *testing.T) {
		var req = Req{
			UserName: "userName",
			Email:    "email",
			Password: "password",
			Name:     "",
		}
		var l = New()
		var err = l.ValidationStruct(req)
		assert.NotNil(t, err)
		log.Info(err)
	})
}
