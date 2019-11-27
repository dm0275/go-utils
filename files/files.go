package files

import (
	"github.com/dm0275/go-utils/errors"
	"io/ioutil"
)

/*
Read a file and return a string
 */
func ReadFile(filePath string) string {
	text,err := ioutil.ReadFile(filePath)
	errors.CheckError(err)
	return string(text)
}