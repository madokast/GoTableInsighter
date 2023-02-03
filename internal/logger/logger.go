/*
日志
*/

package logger

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"runtime"
	"strconv"
	"strings"
	"time"
)

type writer interface {
	WriteString(string) (int, error)
	Flush() error
}

var stdout writer = bufio.NewWriter(os.Stderr)

func Debug(messages ...interface{}) { _log(stdout, 2, _debug, messages...) }
func Info(messages ...interface{})  { _log(stdout, 2, _info, messages...) }
func Warn(messages ...interface{})  { _log(stdout, 2, _warn, messages...) }
func Error(messages ...interface{}) { _log(stdout, 2, _error, messages...) }

const (
	cNORMAL       = "\033[m"
	cRED          = "\033[0;32;31m"
	cLIGHT_RED    = "\033[1;31m"
	cGREEN        = "\033[0;32;32m"
	cLIGHT_GREEN  = "\033[1;32m"
	cBLUE         = "\033[0;32;34m"
	cLIGHT_BLUE   = "\033[1;34m"
	cDARY_GRAY    = "\033[1;30m"
	cCYAN         = "\033[0;36m"
	cLIGHT_CYAN   = "\033[1;36m"
	cPURPLE       = "\033[0;35m"
	cLIGHT_PURPLE = "\033[1;35m"
	cBROWN        = "\033[0;33m"
	cYELLOW       = "\033[1;33m"
	cLIGHT_GRAY   = "\033[0;37m"
	cWHITE        = "\033[1;37m"
)

const (
	_debug = " [" + "DEBUG" + "] "
	_info  = " [" + cGREEN + "INFO" + cNORMAL + "] "
	_warn  = " [" + cRED + "WARN" + cNORMAL + "] "
	_error = " [" + cLIGHT_RED + "ERROR" + cNORMAL + "] "
)

func _log(writer writer, caller_back int, level string, messages ...interface{}) {
	date := time.Now().Format("01-02 15:04:05")
	_, file, lineNo, ok := runtime.Caller(caller_back)
	if !ok {
		file = "unknown"
		lineNo = 0
	} else {
		_, file = path.Split(file)
		i := strings.LastIndex(file, ".go")
		if i > 0 {
			file = file[:i]
		}
	}

	writer.WriteString(date)
	writer.WriteString(level)
	writer.WriteString(file)
	writer.WriteString(":")
	writer.WriteString(strconv.Itoa(lineNo))
	for _, msg := range messages {
		writer.WriteString(" ")
		writer.WriteString(fmt.Sprintf("%v", msg))
	}
	writer.WriteString("\n")
	writer.Flush()
}
