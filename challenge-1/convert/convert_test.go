package convert

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvert(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		var timeFormat string = "24:60:60"

		var output = TimeConv(timeFormat)

		assert.Equal(t, "10.00:100.00:100.00", output)
	})

	t.Run("case 2", func(t *testing.T) {
		var timeFormat string = "00:00:00"

		var output = TimeConv(timeFormat)

		assert.Equal(t, "0.00:0.00:0.00", output)
	})

	t.Run("case 3", func(t *testing.T) {
		var timeFormat string = "05:05:05"

		var output = TimeConv(timeFormat)

		t.Log(output)
	})

	t.Run("case 4", func(t *testing.T) {
		var timeFormat string = "15:15:15"

		var output = TimeConv(timeFormat)

		t.Log(output)
	})
}
