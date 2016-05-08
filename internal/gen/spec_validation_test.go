package gen

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuerySpecValidation(t *testing.T) {
	assert := assert.New(t)

	var fixtures = []struct {
		Q            QuerySpec
		ErrExpected  bool
		PrefixErrMsg string
	}{
		{
			QuerySpec{},
			true, "statement is required",
		},
		{
			QuerySpec{
				Statement: "any",
				Result:    ResultSpec{Type: "any"}},
			true, "invalid result type",
		},
		{
			QuerySpec{
				Statement: "any",
				Result:    ResultSpec{Type: "int64"}},
			false, "",
		},
		{
			QuerySpec{
				Statement: "any",
				Result: ResultSpec{
					Type:   "Row",
					Struct: ResultStructSpec{Name: "any"}}},
			true, "result struct fields is required",
		},
		{
			QuerySpec{
				Statement: "any",
				Result: ResultSpec{
					Type:   "Row",
					Struct: ResultStructSpec{Fields: []string{"any"}}}},
			true, "result struct name is required",
		},
		{
			QuerySpec{
				Statement: "any",
				Result: ResultSpec{
					Type:   "Row",
					Fields: []FieldSpec{{Name: "any"}},
					Struct: ResultStructSpec{Name: "any", Fields: []string{"any"}}}},
			true, "result struct must not be set",
		},
		{
			QuerySpec{
				Statement: "any",
				Result: ResultSpec{
					Type:   "Row",
					Fields: []FieldSpec{{Name: "any", Type: "any"}}}},
			false, "",
		},
		{
			QuerySpec{
				Statement: "any",
				Result: ResultSpec{
					Type:   "Row",
					Struct: ResultStructSpec{Name: "Struct1", Fields: []string{"any"}}}},
			false, "",
		},
		{
			QuerySpec{
				Statement: "any",
				Result: ResultSpec{
					Type:   "None",
					Struct: ResultStructSpec{Name: "Struct1", Fields: []string{"any"}}}},
			true, "cannot set result struct while result type is 'None'",
		},
		{
			QuerySpec{
				Statement: "any",
				Result: ResultSpec{
					Type:   "None",
					Fields: []FieldSpec{{Name: "any", Type: "any"}}}},
			true, "cannot set result fields while result type is 'None'",
		},
		{
			QuerySpec{
				Statement: "any",
				Result: ResultSpec{
					Type:   "int64",
					Struct: ResultStructSpec{Name: "Struct1", Fields: []string{"any"}}}},
			true, "cannot set result struct while result type is",
		},
		{
			QuerySpec{
				Statement: "any",
				Result: ResultSpec{
					Type:   "int64",
					Fields: []FieldSpec{{Name: "any", Type: "any"}}}},
			true, "cannot set result fields while result type is",
		},
		{
			QuerySpec{
				Statement:          "any",
				ParamsStructFields: []FieldSpec{{Name: "any"}},
				ParamsStructRef:    "any",
				Result: ResultSpec{
					Type:   "Row",
					Struct: ResultStructSpec{Name: "Struct1", Fields: []string{"any"}}}},
			true, "cannot set ParamsStructRef while ParamsStructFields is also set",
		},
		{
			QuerySpec{
				Statement:               "any",
				ParamsStructRef:         "any",
				ParamsStructRefValidate: false,
				Result: ResultSpec{
					Type:   "Row",
					Struct: ResultStructSpec{Name: "Struct1", Fields: []string{"any"}}}},
			false, "",
		},
		{
			QuerySpec{
				Statement:          "any",
				ParamsStructFields: []FieldSpec{{Name: "any"}},
				Result: ResultSpec{
					Type:   "Row",
					Struct: ResultStructSpec{Name: "Struct1", Fields: []string{"any"}}}},
			false, "",
		},
	}

	for i, fx := range fixtures {
		err := fx.Q.Validate()
		if !fx.ErrExpected {
			assert.NoError(err, "fixtures: %d", i)
			continue
		}

		assert.Error(err, "fixtures: %d", i)
		assert.True(strings.HasPrefix(err.Error(), fx.PrefixErrMsg), "fixtures: %d: unexpected error message: %s", i, err.Error())
	}
}
