package funcutils

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestMapFunc(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		assert.Equal(t, []string{}, MapFunc([]string{}, strings.ToUpper))
	})

	t.Run("map", func(t *testing.T) {
		assert.Equal(t, []string{"ASDF", "UIOP"}, MapFunc([]string{"asdf", "UIOP"}, strings.ToUpper))
	})
}

func TestIndexFunc(t *testing.T) {
	f := func(s string) bool {
		return strings.ToUpper(s) == "UIOP"
	}

	t.Run("empty", func(t *testing.T) {
		i, found := IndexFunc([]string{}, f)
		assert.Equal(t, 0, i)
		assert.False(t, found)
	})

	t.Run("not found", func(t *testing.T) {
		i, found := IndexFunc([]string{"ABCD", "efgh"}, f)
		assert.Equal(t, 0, i)
		assert.False(t, found)
	})

	t.Run("found", func(t *testing.T) {
		i, found := IndexFunc([]string{"asdf", "UIOP"}, f)
		assert.Equal(t, 1, i)
		assert.True(t, found)
	})
}
