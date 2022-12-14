package templates

import (
	"embed"
	"fmt"
	"io/fs"
	"path"
	"sort"
	"strings"
	"text/template"

	"github.com/alimy/embedx"
)

var (
	//go:embed chi
	//go:embed chi/.gitignore.tmpl
	//go:embed echo
	//go:embed echo/.gitignore.tmpl
	//go:embed fiber
	//go:embed fiber/.gitignore.tmpl
	//go:embed fiber-v2
	//go:embed fiber-v2/.gitignore.tmpl
	//go:embed gin
	//go:embed gin/.gitignore.tmpl
	//go:embed hertz
	//go:embed hertz/.gitignore.tmpl
	//go:embed httprouter
	//go:embed httprouter/.gitignore.tmpl
	//go:embed iris
	//go:embed iris/.gitignore.tmpl
	//go:embed macaron
	//go:embed macaron/.gitignore.tmpl
	//go:embed mux
	//go:embed mux/.gitignore.tmpl
	content embed.FS

	// styles map of style slice to style
	styles = make(map[string]*tmplInfo)
)

// TmplCtx template context for generate project
type TmplCtx struct {
	PkgName    string
	MirPkgName string
	MirVersion string
	EngVersion string
}

// ts style slice alice type
type ts []string

type tmplInfo struct {
	target string
	files  []string
}

func init() {
	for _, s := range []struct {
		styles ts
		info   *tmplInfo
	}{
		{ts{"gin"}, &tmplInfo{
			target: "gin",
			files:  levelStar(5),
		}},
		{ts{"hertz"}, &tmplInfo{
			target: "hertz",
			files:  levelStar(5),
		}},
		{ts{"chi"}, &tmplInfo{
			target: "chi",
			files:  levelStar(5),
		}},
		{ts{"echo"}, &tmplInfo{
			target: "echo",
			files:  levelStar(5),
		}},
		{ts{"fiber"}, &tmplInfo{
			target: "fiber",
			files:  levelStar(5),
		}},
		{ts{"fiber-v2"}, &tmplInfo{
			target: "fiber-v2",
			files:  levelStar(5),
		}},
		{ts{"fiber", "v2"}, &tmplInfo{
			target: "fiber-v2",
			files:  levelStar(5),
		}},
		{ts{"httprouter"}, &tmplInfo{
			target: "httprouter",
			files:  levelStar(5),
		}},
		{ts{"iris"}, &tmplInfo{
			target: "iris",
			files:  levelStar(5),
		}},
		{ts{"macaron"}, &tmplInfo{
			target: "macaron",
			files:  levelStar(5),
		}},
		{ts{"mux"}, &tmplInfo{
			target: "mux",
			files:  levelStar(5),
		}},
	} {
		styles[s.styles.String()] = s.info
	}
}

// VersionOfMir return mir's version
func (c *TmplCtx) VersionOfMir(defVer string) string {
	if c.MirVersion != "" {
		return c.MirVersion
	}
	return defVer
}

// VersionOfEngine return engine's version
func (c *TmplCtx) VersionOfEngine(defVer string) string {
	if c.EngVersion != "" {
		return c.EngVersion
	}
	return defVer
}

func (t *tmplInfo) globFiles(fsys fs.FS) ([]string, error) {
	fps := make([]string, 0)
	for _, p := range t.files {
		list, err := fs.Glob(fsys, p)
		if err != nil {
			return nil, err
		}
		fps = append(fps, list...)
	}
	return fps, nil
}

func (s ts) String() string {
	sort.Strings(s)
	return strings.Join(s, ":")
}

func levelStar(l uint8) []string {
	patterns := make([]string, 0, l)
	for i := uint8(0); i < l; i++ {
		ms := strings.Builder{}
		for j := i; j > 0; j-- {
			ms.WriteString("*/")
		}
		ms.WriteString("*.tmpl")
		patterns = append(patterns, ms.String())
	}
	return patterns
}

func Styles() []string {
	styleNames := make([]string, 0, len(styles))
	for name := range styles {
		styleNames = append(styleNames, name)
	}
	return styleNames
}

func NewTemplate(style []string) (*template.Template, error) {
	sort.Strings(style)
	info, exist := styles[strings.Join(style, ":")]
	if !exist {
		return nil, fmt.Errorf("not exist style(%s) template project", style)
	}
	embedFS, err := fs.Sub(content, info.target)
	if err != nil {
		return nil, err
	}
	embedx.RegisterNamer(embedx.NamerFunc(func(filename string) string {
		ext := path.Ext(filename)
		return filename[:len(filename)-len(ext)]
	}))
	files, err := info.globFiles(embedFS)
	if err != nil {
		return nil, err
	}
	tmpl := template.New("mirc").Funcs(template.FuncMap{
		"notEmptyStr": notEmptyStr,
	})
	return embedx.ParseWith(tmpl, embedFS, files...)
}

func notEmptyStr(s string) bool {
	return s != ""
}
