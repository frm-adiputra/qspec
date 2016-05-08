package gen

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRowScanFromFieldSpec(t *testing.T) {
	assert := assert.New(t)

	var fixtures = []struct {
		Fields   []FieldSpec
		Expected string
	}{
		{[]FieldSpec{{Name: "Field1"}}, "&v.Field1"},
		{[]FieldSpec{{Name: "Field1"}, {Name: "Field2"}}, "&v.Field1, &v.Field2"},
		{[]FieldSpec{{Name: "Field1"}, {Name: "Field2"}, {Name: "Field3"}}, "&v.Field1, &v.Field2, &v.Field3"},
	}

	for _, fx := range fixtures {
		actual := rowScanFromFieldSpec(fx.Fields)
		assert.Equal(fx.Expected, actual)
	}
}
