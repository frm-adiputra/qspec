package gen

import "sort"

func (m *ModelSpec) init() {
	m.initStructs()
	m.initQueries()
}

func (m *ModelSpec) initStructs() {
	var keys []string
	for k := range m.Structs {
		keys = append(keys, k)
	}

	sort.Strings(keys)
	var a []StructSpec
	for _, name := range keys {
		v := m.Structs[name]
		v.name = name
		a = append(a, v)
	}
	m.structs = a
}

func (m *ModelSpec) initQueries() {
	var keys []string
	for k := range m.Queries {
		keys = append(keys, k)
	}

	sort.Strings(keys)
	var a []QuerySpec
	for _, name := range keys {
		v := m.Queries[name]
		v.name = name
		v.cleanedStatement = removeParamsName(v.Statement)
		a = append(a, v)
	}
	m.queries = a
}

// Package returns the package name for this model.
func (m ModelSpec) Package() string { return m.pkg }

// SpecFile returns the spec file for this model.
func (m ModelSpec) SpecFile() string { return m.specFile }

// StructsSequence returns the structs for this model in alphabetical order.
func (m ModelSpec) StructsSequence() []StructSpec { return m.structs }

// QueriesSequence returns the queries for this model in alphabetical order.
func (m ModelSpec) QueriesSequence() []QuerySpec { return m.queries }
