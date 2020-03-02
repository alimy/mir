package cmd

import (
	"log"
	"os"
	"path/filepath"
	"text/template"

	"github.com/spf13/cobra"
)

var (
	dstPath string
	pkgName string
	style   string
)

func init() {
	newCmd := &cobra.Command{
		Use:   "new",
		Short: "create template project",
		Long:  "create template project",
		Run:   newRun,
	}

	// parse flags for agentCmd
	newCmd.Flags().StringVarP(&dstPath, "dst", "d", ".", "genereted destination target directory")
	newCmd.Flags().StringVarP(&pkgName, "pkg", "p", "github.com/alimy/mir-example", "project's package name")
	newCmd.Flags().StringVarP(&style, "type", "t", "gin", "generated engine type style(eg: gin,chi,mux,httprout)")

	// register agentCmd as sub-command
	register(newCmd)
}

// newRun run new command
func newRun(_cmd *cobra.Command, _args []string) {
	path, err := filepath.EvalSymlinks(dstPath)
	if err != nil {
		if os.IsNotExist(err) {
			if !filepath.IsAbs(dstPath) {
				cwd, err := os.Getwd()
				if err != nil {
					log.Fatal(err)
				}
				path = filepath.Join(cwd, dstPath)
			} else {
				path = dstPath
			}
		} else {
			log.Fatal(err)
		}
	}

	tmpls, exist := tmplFiles[style]
	if !exist {
		log.Fatal("not exist style engine")
	}

	ctx := &tmplCtx{
		PkgName: pkgName,
	}
	if err = genProject(ctx, path, tmpls); err != nil {
		log.Fatal(err)
	}
}

func genProject(ctx *tmplCtx, dstPath string, tmpls map[string]tmplInfo) error {
	var (
		err               error
		filePath, dirPath string
		file              *os.File
	)

	tmpl := template.New("mirc")
	for fileName, assetInfo := range tmpls {
		filePath = filepath.Join(dstPath, fileName)
		dirPath = filepath.Dir(filePath)
		if err = os.MkdirAll(dirPath, 0755); err != nil {
			break
		}
		file, err = os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
		if err != nil {
			break
		}
		if assetInfo.isTmpl {
			t, err := tmpl.Parse(MustAssetString(assetInfo.name))
			if err != nil {
				break
			}
			if err = t.Execute(file, ctx); err != nil {
				break
			}
		} else {
			if _, err = file.Write(MustAsset(assetInfo.name)); err != nil {
				break
			}
		}
	}
	return err
}
