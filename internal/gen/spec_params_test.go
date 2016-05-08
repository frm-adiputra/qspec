package gen

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHasInvalidParams(t *testing.T) {
	assert := assert.New(t)

	var fixtures = []struct {
		Stmt     string
		Expected bool
	}{
		{"SELECT * FROM T WHERE c = ?", true},
		{"SELECT * FROM T WHERE c = ? ", true},
		{"SELECT * FROM T WHERE c = ?1", true},
		{"SELECT * FROM T WHERE c = ?1a", true},
		{"SELECT * FROM T WHERE c = ?1 ", true},
		{"SELECT * FROM T WHERE c = ?1a ", true},
		{"SELECT * FROM T WHERE c = ?a", false},
	}

	for _, fx := range fixtures {
		actual := hasInvalidParams(fx.Stmt)
		assert.Equal(fx.Expected, actual, fx.Stmt)
	}
}

func TestCreateParamsSequence(t *testing.T) {
	assert := assert.New(t)

	var fixtures = []struct {
		Stmt     string
		Expected []string
	}{
		{"SELECT * FROM T WHERE a = ?id", []string{"id"}},
		{"SELECT * FROM T WHERE a = ?id ", []string{"id"}},
		{"SELECT * FROM T WHERE a = ?id AND b = ?23", []string{"id", "23"}},
		{"SELECT * FROM T WHERE a = ?id AND b = ?23 ", []string{"id", "23"}},
		{"SELECT * FROM T WHERE a = ?id AND b = ?23 AND c = ?", []string{"id", "23"}},
		{"SELECT * FROM T WHERE a = ?id AND b = ?23 AND c = ? ", []string{"id", "23"}},
	}

	for _, fx := range fixtures {
		actual := createParamsSequence(fx.Stmt)
		assert.Equal(fx.Expected, actual, fx.Stmt)
	}
}

func TestCreateParamsSet(t *testing.T) {
	assert := assert.New(t)

	var fixtures = []struct {
		Sequence []string
		Expected []string
	}{
		{[]string{"id", "ak", "id", "ak", "ll"}, []string{"id", "ak", "ll"}},
		{[]string{"id", "ak", "ll"}, []string{"id", "ak", "ll"}},
	}

	for _, fx := range fixtures {
		actual := createParamsSet(fx.Sequence)
		assert.Equal(fx.Expected, actual)
	}
}

func TestRemoveParamsName(t *testing.T) {
	assert := assert.New(t)

	var fixtures = []struct {
		Stmt     string
		Expected string
	}{
		{"SELECT * FROM T WHERE a = ?id", "SELECT * FROM T WHERE a = ?"},
		{"SELECT * FROM T WHERE a = ?id ", "SELECT * FROM T WHERE a = ? "},
		{"SELECT * FROM T WHERE a = ?id AND b = ?23", "SELECT * FROM T WHERE a = ? AND b = ?"},
		{"SELECT * FROM T WHERE a = ?id AND b = ?23 ", "SELECT * FROM T WHERE a = ? AND b = ? "},
		{"SELECT * FROM T WHERE a = ?id AND b = ?23 AND c = ?", "SELECT * FROM T WHERE a = ? AND b = ? AND c = ?"},
		{"SELECT * FROM T WHERE a = ?id AND b = ?23 AND c = ? ", "SELECT * FROM T WHERE a = ? AND b = ? AND c = ? "},
	}

	for _, fx := range fixtures {
		actual := removeParamsName(fx.Stmt)
		assert.Equal(fx.Expected, actual)
	}
}
