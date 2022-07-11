package log

import (
	"fmt"
	"time"
)

const (
	INFO  = iota
	DEBUG = iota
)

func Write(filename string, content string) error {
	content = fmt.Sprintf("%s : %s \n", time.Now().String(), content)
	saveFile(filename, content)
	return nil
}
