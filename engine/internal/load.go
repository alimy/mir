// Copyright 2025 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package internal

import (
	"bytes"
	"embed"
	"errors"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"go/types"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"reflect"
	"slices"
	"sort"
	"strconv"
	"strings"
	"text/template"
	"time"

	"github.com/alimy/mir/v5"

	"golang.org/x/tools/go/ast/astutil"
	"golang.org/x/tools/go/packages"
)

type (
	// A SchemaSpec holds a serializable version of an mir.Schema
	// and its Go package and module information.
	SchemaSpec struct {
		PkgPath  string
		PkgAlias string
		Module   *packages.Module
		Names    []string
	}

	// Config holds the configuration for loading an mir/schema package.
	Config struct {
		// InDebug whether in debug mode
		InDebug bool
		// InitOpts origin init options
		InitOpts string
		// SchemaPath is the path list for the schema package.
		SchemaPath []string
		// BuildFlags are forwarded to the package.Config when
		// loading the schema package.
		BuildFlags []string
		// AssertTypeImports AssertType function params package path.
		AssertTypeImports []string
		// AssertTypeSpec AssertType function spec.
		AssertTypeSpec string
	}
)

// Load loads the schemas package and build the Go plugin with this info.
func (c *Config) Load() error {
	specs, err := c.load()
	if err != nil {
		return fmt.Errorf("mirc/load: parse schema dir: %w", err)
	}
	if len(specs) == 0 {
		return fmt.Errorf("mirc/load: no schema found in: %+v", c.SchemaPath)
	}
	var b bytes.Buffer
	err = buildTmpl.ExecuteTemplate(&b, "main", struct {
		*Config
		Schemas []*SchemaSpec
	}{
		Config:  c,
		Schemas: specs,
	})
	if err != nil {
		return fmt.Errorf("mirc/load: execute template: %w", err)
	}
	buf, err := format.Source(b.Bytes())
	if err != nil {
		return fmt.Errorf("mirc/load: format template: %w", err)
	}

	targetDir := ".mirc"
	if err := os.MkdirAll(targetDir, os.ModePerm); err != nil {
		return err
	}
	target := path.Join(targetDir, fmt.Sprintf("main_t%d.go", time.Now().Unix()))
	if err := os.WriteFile(target, buf, 0644); err != nil {
		return fmt.Errorf("mirc/load: write file %s: %w", target, err)
	}
	// cleanup when not in debug mode
	if !c.InDebug {
		defer os.RemoveAll(targetDir)
	}

	if _, err = gorun(target, c.BuildFlags); err != nil {
		return err
	}
	return nil
}

// load the ent/schema info.
func (c *Config) load() ([]*SchemaSpec, error) {
	// mirInterface holds the reflect.Type of mir.Schema interface.
	mirInterface := reflect.TypeOf(struct{ mir.Schema }{}).Field(0).Type

	patterns := anyPathPatterns(c.SchemaPath)
	pkgs, err := packages.Load(&packages.Config{
		BuildFlags: c.BuildFlags,
		Mode:       packages.NeedName | packages.NeedTypes | packages.NeedTypesInfo | packages.NeedModule,
	}, append(patterns, mirInterface.PkgPath())...)
	if err != nil {
		return nil, fmt.Errorf("loading package: %w", err)
	}
	if len(pkgs) < 2 {
		// Check if the package loading failed due to Go-related
		// errors, such as 'missing go.sum entry'.
		if err := golist(c.BuildFlags, patterns...); err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("missing package information for: %+v", patterns)
	}

	var (
		mirPkg    *packages.Package
		mirPkgIdx int
	)
	for idx, pkg := range pkgs {
		if len(pkg.Errors) != 0 {
			return nil, c.loadError(pkg.Errors[0])
		}
		if pkg.PkgPath == mirInterface.PkgPath() {
			mirPkg = pkg
			mirPkgIdx = idx
		}
	}
	if mirPkg == nil {
		return nil, fmt.Errorf("missing mir package information")
	}
	pkgs = slices.Delete(pkgs, mirPkgIdx, mirPkgIdx+1)

	iface := mirPkg.Types.Scope().Lookup(mirInterface.Name()).Type().Underlying().(*types.Interface)
	specs := make(map[string]*SchemaSpec, len(pkgs))
	for _, pkg := range pkgs {
		schema, exist := specs[pkg.PkgPath]
		if !exist {
			schema = &SchemaSpec{
				PkgPath: pkg.PkgPath,
				Module:  pkg.Module,
			}
			specs[pkg.PkgPath] = schema
		}
		for k, v := range pkg.TypesInfo.Defs {
			typ, ok := v.(*types.TypeName)
			if !ok || !k.IsExported() || !types.Implements(typ.Type(), iface) {
				continue
			}
			spec, ok := k.Obj.Decl.(*ast.TypeSpec)
			if !ok {
				return nil, fmt.Errorf("invalid declaration %T for %s", k.Obj.Decl, k.Name)
			}
			if _, ok := spec.Type.(*ast.StructType); !ok {
				return nil, fmt.Errorf("invalid spec type %T for %s", spec.Type, k.Name)
			}
			schema.Names = append(schema.Names, k.Name)
		}
	}

	var (
		schemas []*SchemaSpec
		pkgIdx  int
	)
	for _, it := range specs {
		if len(it.Names) > 0 {
			pkgIdx++
			it.PkgAlias = fmt.Sprintf("s%d", pkgIdx)
			sort.Strings(it.Names)
			schemas = append(schemas, it)
		}
	}
	return schemas, nil
}

func (c *Config) loadError(perr packages.Error) (err error) {
	if strings.Contains(perr.Msg, "import cycle not allowed") {
		for _, it := range c.SchemaPath {
			if cause := c.cycleCause(it); cause != "" {
				perr.Msg += "\n" + cause
			}
		}
	}
	err = perr
	if perr.Pos == "" {
		// Strip "-:" prefix in case of empty position.
		err = errors.New(perr.Msg)
	}
	return err
}

func (c *Config) cycleCause(path string) (cause string) {
	dir, err := parser.ParseDir(token.NewFileSet(), path, nil, 0)
	// Ignore reporting in case of parsing
	// error, or there no packages to parse.
	if err != nil || len(dir) == 0 {
		return
	}
	// Find the package that contains the schema, or
	// extract the first package if there is only one.
	pkg := dir[filepath.Base(path)]
	if pkg == nil {
		for _, v := range dir {
			pkg = v
			break
		}
	}
	// Package local declarations used by schema fields.
	locals := make(map[string]bool)
	for _, f := range pkg.Files {
		for _, d := range f.Decls {
			g, ok := d.(*ast.GenDecl)
			if !ok || g.Tok != token.TYPE {
				continue
			}
			for _, s := range g.Specs {
				ts, ok := s.(*ast.TypeSpec)
				if !ok || !ts.Name.IsExported() {
					continue
				}
				// Non-struct types such as "type Role int".
				st, ok := ts.Type.(*ast.StructType)
				if !ok {
					locals[ts.Name.Name] = true
					continue
				}
				var embedSchema bool
				astutil.Apply(st.Fields, func(c *astutil.Cursor) bool {
					f, ok := c.Node().(*ast.Field)
					if ok {
						switch x := f.Type.(type) {
						case *ast.SelectorExpr:
							if x.Sel.Name == "Schema" {
								embedSchema = true
							}
						case *ast.Ident:
							// A common pattern is to create local base schema to be embedded by other schemas.
							if name := strings.ToLower(x.Name); name == "schema" {
								embedSchema = true
							}
						}
					}
					// Stop traversing the AST in case an ~ent.Schema is embedded.
					return !embedSchema
				}, nil)
				if !embedSchema {
					locals[ts.Name.Name] = true
				}
			}
		}
	}
	// No local declarations to report.
	if len(locals) == 0 {
		return
	}
	// Usage of local declarations by schema fields.
	goTypes := make(map[string]bool)
	for _, f := range pkg.Files {
		for _, d := range f.Decls {
			f, ok := d.(*ast.FuncDecl)
			if !ok || f.Name.Name != "Fields" || f.Type.Params.NumFields() != 0 || f.Type.Results.NumFields() != 1 {
				continue
			}
			astutil.Apply(f.Body, func(cursor *astutil.Cursor) bool {
				i, ok := cursor.Node().(*ast.Ident)
				if ok && locals[i.Name] {
					goTypes[i.Name] = true
				}
				return true
			}, nil)
		}
	}
	names := make([]string, 0, len(goTypes))
	for k := range goTypes {
		names = append(names, strconv.Quote(k))
	}
	sort.Strings(names)
	if len(names) > 0 {
		cause = fmt.Sprintf("To resolve this issue, move the custom types used by the generated code to a separate package: %s", strings.Join(names, ", "))
	}
	return
}

var (
	//go:embed template/main.tmpl
	files     embed.FS
	buildTmpl = templates()
)

func templates() *template.Template {
	tmpl := template.Must(template.New("templates").
		ParseFS(files, "template/main.tmpl"))
	return tmpl
}

// run 'go run' command and return its output.
func gorun(target string, buildFlags []string) (string, error) {
	s, err := gocmd("run", buildFlags, target)
	if err != nil {
		return "", fmt.Errorf("mirc/load: %s", err)
	}
	return s, nil
}

// golist checks if 'go list' can be executed on the given target.
func golist(buildFlags []string, targets ...string) error {
	_, err := gocmd("list", buildFlags, targets...)
	return err
}

// goCmd runs a go command and returns its output.
func gocmd(command string, buildFlags []string, targets ...string) (string, error) {
	args := []string{command}
	args = append(args, buildFlags...)
	args = append(args, targets...)
	cmd := exec.Command("go", args...)
	stderr := bytes.NewBuffer(nil)
	stdout := bytes.NewBuffer(nil)
	cmd.Stderr = stderr
	cmd.Stdout = stdout
	if err := cmd.Run(); err != nil {
		return "", errors.New(strings.TrimSpace(stderr.String()))
	}
	return stdout.String(), nil
}
