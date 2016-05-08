package gen

// Name returns the struct's name.
func (s StructSpec) Name() string { return s.name }

// IsUsingValidations returns true if the struct's needs to be validated.
func (s StructSpec) IsUsingValidations() bool {
	for _, f := range s.Fields {
		if len(f.Validations) != 0 {
			return true
		}
	}
	return false
}
