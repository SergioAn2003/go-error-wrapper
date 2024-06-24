package errorWrapper

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_PtrWithOP(t *testing.T) {
	t.Run("Should Return", func(t *testing.T) {
		err := errors.New("some error")
		PtrWithOP(&err, "wrap")
		assert.Equal(t, "wrap: some error", err.Error())
	})

	t.Run("Should Not Return", func(t *testing.T) {
		var err error
		PtrWithOP(&err, "wrap")
		assert.Nil(t, err)
	})
}

func Test_WithOP(t *testing.T) {
	t.Run("Should Return", func(t *testing.T) {
		err := WithOP(errors.New("some error"), "wrap")
		assert.Equal(t, "wrap: some error", err.Error())
	})

	t.Run("Should Not Return", func(t *testing.T) {
		err := WithOP(nil, "wrap")
		assert.Nil(t, err)
	})
}
