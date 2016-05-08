package gen

//go:generate go-bindata -o assets/assets.go -pkg assets -prefix _assets _assets/...

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"gopkg.in/yaml.v2"

	"golang.org/x/tools/imports"

	"github.com/frm-adiputra/qspec/internal/gen/assets"
)

var modelSourceCodeTmpl = template.New("root")

func init() {
	a, err := assets.AssetDir("templates")
	if err != nil {
		panic(err)
	}

	for _, n := range a {
		t := modelSourceCodeTmpl.New(n)
		_, err = t.Parse(string(assets.MustAsset("templates/" + n)))
		if err != nil {
			panic(err)
		}
	}
}

// GenerateSourceCodeFromYAML generates model source code from a YAML spec.
func GenerateSourceCodeFromYAML(specFile string) error {
	pkg := packageFromPath(specFile)

	r, err := os.Open(specFile)
	if err != nil {
		return err
	}
	defer r.Close()

	err = os.MkdirAll(pkg, 0775)
	if err != nil {
		return err
	}

	w, err := os.Create(filepath.Join(pkg, pkg+".go"))
	if err != nil {
		return err
	}
	defer w.Close()

	m, err := NewModelSpecFromYAML(pkg, specFile, r)
	if err != nil {
		return err
	}

	b, err := m.GenerateSourceCode()
	if err != nil {
		return err
	}

	w.Write(b)
	return nil
}

// NewModelSpecFromYAML creates a new ModelSpec from a YAML spec.
func NewModelSpecFromYAML(pkg, specFile string, r io.Reader) (*ModelSpec, error) {
	m := &ModelSpec{
		pkg:      pkg,
		specFile: specFile,
	}

	b, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(b, m)
	if err != nil {
		return nil, err
	}

	m.init()

	err = m.Validate()
	if err != nil {
		return nil, fmt.Errorf("validation error (%s): %s", specFile, err)
	}

	return m, nil
}

// GenerateSourceCode generates model source code.
func (m ModelSpec) GenerateSourceCode() ([]byte, error) {
	var b bytes.Buffer
	err := modelSourceCodeTmpl.ExecuteTemplate(&b, "MODEL", m)
	if err != nil {
		return nil, err
	}

	// code, err := format.Source(b.Bytes())
	// if err != nil {
	// 	return b.Bytes(), fmt.Errorf("format error: %s", err)
	// }

	opt := imports.Options{FormatOnly: true, Comments: true, TabIndent: true}
	code, err := imports.Process(filepath.Join(m.pkg, m.pkg+".go"), b.Bytes(), &opt)
	if err != nil {
		return b.Bytes(), fmt.Errorf("format imports error: %s", err)
	}

	return code, nil
}

func packageFromPath(tomlPath string) string {
	base := filepath.Base(tomlPath)
	idx := strings.LastIndex(base, ".")
	if idx == -1 {
		return base
	}

	return base[:idx]
}
