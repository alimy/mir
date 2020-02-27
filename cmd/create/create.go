package create

import (
	"errors"
	"os"
	"path/filepath"
)

// RunCmd run new command
func RunCmd(dst string, style string) error {
	dstPath, err := filepath.EvalSymlinks(dst)
	if err != nil {
		return err
	}
	tmpls, exist := tmplFiles[style]
	if !exist {
		return errors.New("not exist style engine")
	}
	return genProject(dstPath, tmpls)
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
