package gen

// ResultType constants
const (
	ResultTypeNone = "None"
	ResultTypeRow  = "Row"
	ResultTypeRows = "Rows"
)

// ModelSpec represents a model specification.
type ModelSpec struct {
	Description string
	Imports     []string
	Structs     map[string]StructSpec
	Queries     map[string]QuerySpec

	pkg      string
	specFile string
	structs  []StructSpec
	queries  []QuerySpec
}

// QuerySpec represents a query specification.
type QuerySpec struct {
	Description     string
	Statement       string
	Prepared        bool
	ParamsStruct    []FieldSpec
	ParamsStructRef string
	Result          ResultSpec

	name             string
	cleanedStatement string
	params           ParamsSpec
}

// ResultSpec represents a query result specification.
type ResultSpec struct {
	Type   string
	Fields []FieldSpec
	Struct ResultStructSpec
}

// ResultStructSpec represents a query result struct specification.
type ResultStructSpec struct {
	Name   string
	Fields []string
}

// StructSpec represents a struct specification.
type StructSpec struct {
	Description string
	Fields      []FieldSpec

	name string
}

// FieldSpec represents a field specification.
type FieldSpec struct {
	Name        string
	Type        string
	Tag         string
	Validations map[string]interface{}
}

// ParamsSpec represents parameters specification from an SQL statement.
type ParamsSpec struct {
	Declaration []string
	Usage       []string
}
