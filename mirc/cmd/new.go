package cmd

import (
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var (
	dstPath string
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
				if  err != nil {
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

	if err = genProject(path, tmpls); err != nil {
		log.Fatal(err)
	}
}

func genProject(dstPath string, tmpls map[string]string) error {
	var (
		err               error
		filePath, dirPath string
		file              *os.File
	)

	for fileName, assetName := range tmpls {
		filePath = filepath.Join(dstPath, fileName)
		dirPath = filepath.Dir(filePath)
		if err = os.MkdirAll(dirPath, 0755); err != nil {
			break
		}
		file, err = os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
		if err != nil {
			break
		}
		if _, err = file.Write(MustAsset(assetName)); err != nil {
			break
		}
	}
	return err
}
