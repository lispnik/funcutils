package funcutils

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"slices"
	"strings"
	"testing"
	"unicode"
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

func TestGroupByFunc(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		grouped := GroupByFunc([]string{}, Identity[string])
		require.Empty(t, grouped)
	})

	t.Run("group by", func(t *testing.T) {
		groupByFunc := func(s string) string {
			if len(s) > 0 {
				// the most retarded thing in Go, maybe
				r := []rune(s)
				return string(unicode.ToUpper(r[0]))
			}
			return ""
		}

		names := []string{"Alice", "Bob", "", "Brenda", "Charlie", "Christina", "", "Derek", "Eva", "Ethan", ""}
		expected := map[string][]string{
			"A": {"Alice"},
			"B": {"Bob", "Brenda"},
			"C": {"Charlie", "Christina"},
			"D": {"Derek"},
			"E": {"Ethan", "Eva"},
			"":  {"", "", ""},
		}
		phonebook := GroupByFunc(names, groupByFunc)
		for k, v := range phonebook {
			w := slices.Clone(v)
			slices.Sort(w)
			require.Equal(t, expected[k], w)
		}
		require.Equal(t, len(expected), len(phonebook))
	})
}
