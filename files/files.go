package files

import (
	"github.com/dm0275/go-utils/errors"
	"io"
	"io/ioutil"
	"os"
)

/*
Read a file and return a string
 */
func ReadFile(filePath string) string {
	text,err := ioutil.ReadFile(filePath)
	errors.CheckError(err)
	return string(text)
}

func OverwriteFile(filePath string, content string, filePermissions os.FileMode) error {
	file, err := os.OpenFile(filePath, os.O_RDWR, filePermissions)
	errors.CheckError(err)

	_, err = io.WriteString(file, content)
	errors.CheckError(err)

	file.Sync()
	return err
}