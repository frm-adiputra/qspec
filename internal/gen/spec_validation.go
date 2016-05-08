package gen

import (
	"errors"
	"fmt"
)

// Validate validates model spec.
func (m *ModelSpec) Validate() error {
	var err error
	for i, q := range m.queries {
		q.params, err = CreateParamsSpec(q.Statement)
		if err != nil {
			return fmt.Errorf("query '%s': %s", q.name, err)
		}

		err = q.Validate()
		if err != nil {
			return fmt.Errorf("query '%s': %s", q.name, err)
		}

		m.queries[i] = q
	}

	for _, s := range m.structs {
		err = s.Validate()
		if err != nil {
			return fmt.Errorf("struct '%s': %s", s.name, err)
		}
	}
	return nil
}

// Validate validates query spec.
func (q QuerySpec) Validate() error {
	// Statement is required
	if q.Statement == "" {
		return errors.New("statement is required")
	}

	// Check result type is having correct value
	switch q.Result.Type {
	case
		ResultTypeNone,
		ResultTypeRow,
		ResultTypeRows,
		"string",
		"int64",
		"float64",
		"bool",
		"sql.NullString",
		"sql.NullInt64",
		"sql.NullFloat64",
		"sql.NullBool":
	default:
		return errors.New("invalid result type: " + q.Result.Type)
	}

	// Struct name and struct fields must be set both or none
	if q.Result.Struct.Name == "" && len(q.Result.Struct.Fields) != 0 {
		return errors.New("result struct name is required")
	}
	if q.Result.Struct.Name != "" && len(q.Result.Struct.Fields) == 0 {
		return errors.New("result struct fields is required")
	}

	// Result struct and result fields are mutually exclusive (cannot be
	// set both)
	if (q.Result.Struct.Name != "" || len(q.Result.Struct.Fields) != 0) && len(q.Result.Fields) != 0 {
		return errors.New("result struct must not be set if result fields is set")
	}

	// Check relations between result type and result struct/result fields
	switch q.Result.Type {
	case ResultTypeNone:
		if q.Result.Struct.Name != "" || len(q.Result.Struct.Fields) != 0 {
			return errors.New("cannot set result struct while result type is 'None'")
		}
		if len(q.Result.Fields) != 0 {
			return errors.New("cannot set result fields while result type is 'None'")
		}
	case
		"string",
		"int64",
		"float64",
		"bool",
		"sql.NullString",
		"sql.NullInt64",
		"sql.NullFloat64",
		"sql.NullBool":
		if q.Result.Struct.Name != "" || len(q.Result.Struct.Fields) != 0 {
			return fmt.Errorf("cannot set result struct while result type is '%s'", q.Result.Type)
		}
		if len(q.Result.Fields) != 0 {
			return fmt.Errorf("cannot set result fields while result type is '%s'", q.Result.Type)
		}
	}

	// Validate result fields
	for _, f := range q.Result.Fields {
		err := f.Validate()
		if err != nil {
			return err
		}
	}

	if q.ParamsStructRef != "" && len(q.ParamsStructFields) != 0 {
		return errors.New("cannot set ParamsStructRef while ParamsStructFields is also set")
	}

	return nil
}

// Validate validates the struct spec.
func (s StructSpec) Validate() error {
	for _, f := range s.Fields {
		err := f.Validate()
		if err != nil {
			return err
		}
	}

	return nil
}

// Validate validates the field spec.
func (f FieldSpec) Validate() error {
	if f.Name == "" {
		return errors.New("Field name required")
	}

	if f.Type == "" {
		return errors.New("Field type required")
	}

	return nil
}
