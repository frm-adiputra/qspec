package gen

import (
	"strings"
	"testing"

	"github.com/frm-adiputra/qspec/internal/gen/assets"
	"github.com/stretchr/testify/assert"
)

func TestPackageFromTOMLPath(t *testing.T) {
	assert := assert.New(t)

	var fixtures = []struct{ TOMLPath, Expected string }{
		{"aa/bb/cc", "cc"},
		{"aa/bb/cc.toml", "cc"},
		{"aa/bb/cc.yml", "cc"},
		{"aa/bb/cc.custom", "cc"},
	}

	for _, fx := range fixtures {
		actual := packageFromPath(fx.TOMLPath)
		assert.Equal(fx.Expected, actual)
	}
}

func TestGenerateSourceCodeQueryRow(t *testing.T) {
	assert := assert.New(t)

	r := strings.NewReader(string(assets.MustAsset("samples/sample1.yml")))
	m, err := NewModelSpecFromYAML("sample", "samples/sample.yml", r)
	if !assert.NoError(err) {
		t.FailNow()
	}

	b, err := m.GenerateSourceCode()
	if !assert.NoError(err) {
		// if b != nil {
		// 	t.Logf("Generated code:\n%s", b)
		// }
		t.FailNow()
	}

	expected := strings.TrimSpace(string(assets.MustAsset("samples/_sample1_yml.go")))
	actual := strings.TrimSpace(string(b))
	if !assert.Equal(expected, actual) {
		t.Logf("(expected):\n%s", expected)
		t.Logf("(actual):\n%s", actual)
	}
}

func TestGenerateSourceCodeNoParams(t *testing.T) {
	assert := assert.New(t)

	r := strings.NewReader(string(assets.MustAsset("samples/no_params.yml")))
	m, err := NewModelSpecFromYAML("sample", "samples/sample.yml", r)
	if !assert.NoError(err) {
		t.FailNow()
	}

	b, err := m.GenerateSourceCode()
	if !assert.NoError(err) {
		// if b != nil {
		// 	t.Logf("Generated code:\n%s", b)
		// }
		t.FailNow()
	}

	expected := strings.TrimSpace(string(assets.MustAsset("samples/_no_params.go")))
	actual := strings.TrimSpace(string(b))
	if !assert.Equal(expected, actual) {
		t.Logf("(expected):\n%s", expected)
		t.Logf("(actual):\n%s", actual)
	}
}

func TestGenerateSourceCodeQueryRows(t *testing.T) {
	assert := assert.New(t)

	r := strings.NewReader(string(assets.MustAsset("samples/query_rows.yml")))
	m, err := NewModelSpecFromYAML("sample", "samples/sample.yml", r)
	if !assert.NoError(err) {
		t.FailNow()
	}

	b, err := m.GenerateSourceCode()
	if !assert.NoError(err) {
		// if b != nil {
		// 	t.Logf("Generated code:\n%s", b)
		// }
		t.FailNow()
	}

	expected := strings.TrimSpace(string(assets.MustAsset("samples/_query_rows.go")))
	actual := strings.TrimSpace(string(b))
	if !assert.Equal(expected, actual) {
		t.Logf("(expected):\n%s", expected)
		t.Logf("(actual):\n%s", actual)
	}
}

func TestGenerateSourceCodeQueryExec(t *testing.T) {
	assert := assert.New(t)

	r := strings.NewReader(string(assets.MustAsset("samples/query_exec.yml")))
	m, err := NewModelSpecFromYAML("sample", "samples/sample.yml", r)
	if !assert.NoError(err) {
		t.FailNow()
	}

	b, err := m.GenerateSourceCode()
	if !assert.NoError(err) {
		// if b != nil {
		// 	t.Logf("Generated code:\n%s", b)
		// }
		t.FailNow()
	}

	expected := strings.TrimSpace(string(assets.MustAsset("samples/_query_exec.go")))
	actual := strings.TrimSpace(string(b))
	if !assert.Equal(expected, actual) {
		t.Logf("(expected):\n%s", expected)
		t.Logf("(actual):\n%s", actual)
	}
}

func TestGenerateSourceCodeQueryScalar(t *testing.T) {
	assert := assert.New(t)

	r := strings.NewReader(string(assets.MustAsset("samples/query_scalar.yml")))
	m, err := NewModelSpecFromYAML("sample", "samples/sample.yml", r)
	if !assert.NoError(err) {
		t.FailNow()
	}

	b, err := m.GenerateSourceCode()
	if !assert.NoError(err) {
		// if b != nil {
		// 	t.Logf("Generated code:\n%s", b)
		// }
		t.FailNow()
	}

	expected := strings.TrimSpace(string(assets.MustAsset("samples/_query_scalar.go")))
	actual := strings.TrimSpace(string(b))
	if !assert.Equal(expected, actual) {
		t.Logf("(expected):\n%s", expected)
		t.Logf("(actual):\n%s", actual)
	}
}

func TestGenerateSourceCodeNoPrepared(t *testing.T) {
	assert := assert.New(t)

	r := strings.NewReader(string(assets.MustAsset("samples/no_prepared.yml")))
	m, err := NewModelSpecFromYAML("sample", "samples/sample.yml", r)
	if !assert.NoError(err) {
		t.FailNow()
	}

	b, err := m.GenerateSourceCode()
	if !assert.NoError(err) {
		// if b != nil {
		// 	t.Logf("Generated code:\n%s", b)
		// }
		t.FailNow()
	}

	expected := strings.TrimSpace(string(assets.MustAsset("samples/_no_prepared.go")))
	actual := strings.TrimSpace(string(b))
	if !assert.Equal(expected, actual) {
		t.Logf("(expected):\n%s", expected)
		t.Logf("(actual):\n%s", actual)
	}
}

func TestGenerateSourceCodeImport(t *testing.T) {
	assert := assert.New(t)

	r := strings.NewReader(string(assets.MustAsset("samples/imports.yml")))
	m, err := NewModelSpecFromYAML("sample", "samples/sample.yml", r)
	if !assert.NoError(err) {
		t.FailNow()
	}

	b, err := m.GenerateSourceCode()
	if !assert.NoError(err) {
		// if b != nil {
		// 	t.Logf("Generated code:\n%s", b)
		// }
		t.FailNow()
	}

	expected := strings.TrimSpace(string(assets.MustAsset("samples/_imports.go")))
	actual := strings.TrimSpace(string(b))
	if !assert.Equal(expected, actual) {
		t.Logf("(expected):\n%s", expected)
		t.Logf("(actual):\n%s", actual)
	}
}

func TestGenerateSourceCodeStructAsQueryParams(t *testing.T) {
	assert := assert.New(t)

	r := strings.NewReader(string(assets.MustAsset("samples/struct_as_query_params.yml")))
	m, err := NewModelSpecFromYAML("sample", "samples/sample.yml", r)
	if !assert.NoError(err) {
		t.FailNow()
	}

	b, err := m.GenerateSourceCode()
	if !assert.NoError(err) {
		if b != nil {
			t.Logf("Generated code:\n%s", b)
		}
		t.FailNow()
	}

	expected := strings.TrimSpace(string(assets.MustAsset("samples/_struct_as_query_params.go")))
	actual := strings.TrimSpace(string(b))
	if !assert.Equal(expected, actual) {
		t.Logf("(expected):\n%s", expected)
		t.Logf("(actual):\n%s", actual)
	}
}
