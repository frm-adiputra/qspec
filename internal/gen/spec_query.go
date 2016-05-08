package gen

import (
	"bytes"
	"fmt"
	"strings"
)

// Name returns the query's name.
func (q QuerySpec) Name() string { return q.name }

// CleanedStatement returns the query's statement with params name removed.
func (q QuerySpec) CleanedStatement() string { return q.cleanedStatement }

// ReturnType returns the query's return type.
func (q QuerySpec) ReturnType() string {
	switch q.Result.Type {
	case ResultTypeNone:
		return ""
	case ResultTypeRow, ResultTypeRows:
		if q.Result.Struct.Name != "" {
			return q.Result.Struct.Name
		}
		return q.Name() + "Result"
	}

	return q.Result.Type
}

// ResultStruct returns the query's result struct name.
func (q QuerySpec) ResultStruct() StructSpec {
	if !(q.Result.Type == ResultTypeRow || q.Result.Type == ResultTypeRows) {
		return StructSpec{}
	}

	s := StructSpec{
		Description: fmt.Sprintf("represents the result of %s query.", q.Name()),
		Fields:      q.Result.Fields,
		name:        q.Name() + "Result",
	}

	return s
}

// DeclareResultStruct returns true if the query must declare its result struct.
func (q QuerySpec) DeclareResultStruct() bool {
	switch q.Result.Type {
	case ResultTypeRow, ResultTypeRows:
		if q.Result.Struct.Name == "" {
			return true
		}
	}

	return false
}

// ParamsStructName returns the name of the struct of parameters.
func (q QuerySpec) ParamsStructName() string {
	if q.ParamsStructRef != "" {
		return q.ParamsStructRef
	}

	if len(q.ParamsStructFields) != 0 {
		return q.Name() + "Params"
	}

	return ""
}

// ParamsDeclaration returns the parameters' declaration string.
func (q QuerySpec) ParamsDeclaration() string {
	s := q.ParamsStructName()
	if s != "" {
		return "params " + s
	}

	if len(q.params.Declaration) == 0 {
		return ""
	}

	return strings.Join(q.params.Declaration, ", ") + " interface{}"
}

// ParamsUsage returns the parameters' usage string.
func (q QuerySpec) ParamsUsage() string {
	if !q.IsUsingParamsStruct() {
		return strings.Join(q.params.Usage, ", ")
	}

	if len(q.params.Usage) == 0 {
		return ""
	}

	return "params." + strings.Join(q.params.Usage, ", params.")
}

// IsUsingParamsStruct returns true if the query uses a struct as its
// collection of params.
func (q QuerySpec) IsUsingParamsStruct() bool {
	return q.ParamsStructRef != "" || len(q.ParamsStructFields) != 0
}

// ParamsStruct returns the query's params struct.
func (q QuerySpec) ParamsStruct() StructSpec {
	if len(q.ParamsStructFields) == 0 {
		return StructSpec{}
	}

	s := StructSpec{
		Description: fmt.Sprintf("represents query parameters of %s.", q.Name()),
		Fields:      q.ParamsStructFields,
		name:        q.ParamsStructName(),
	}

	return s
}

// IsUsingParamsStructValidations returns true if the query struct of params
// needs to be validated.
func (q QuerySpec) IsUsingParamsStructValidations() bool {
	return (q.ParamsStructRef != "" && q.ParamsStructRefValidate) ||
		q.ParamsStruct().IsUsingValidations()
}

// RowScan returns the rows scan arguments as string.
func (q QuerySpec) RowScan() string {
	if q.DeclareResultStruct() {
		return rowScanFromFieldSpec(q.Result.Fields)
	}

	return rowScanFromFieldNames(q.Result.Struct.Fields)
}

func rowScanFromFieldSpec(a []FieldSpec) string {
	var b bytes.Buffer
	for i, f := range a {
		if i > 0 {
			b.WriteString(", ")
		}
		b.WriteString("&v.")
		b.WriteString(f.Name)
	}

	return b.String()
}

func rowScanFromFieldNames(a []string) string {
	var b bytes.Buffer
	for i, f := range a {
		if i > 0 {
			b.WriteString(", ")
		}
		b.WriteString("&v.")
		b.WriteString(f)
	}

	return b.String()
}
