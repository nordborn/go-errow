// Package errow creates error with context and adds line number as a prefix
// of the text representation of the given error
// `<package/file.go:15>: error message`.
//
// Also, it puts the error to a stack of errors thanks to github.com/pkg/errors
// and you can also print the stack via formatted output `"%+v"`.
// Names of methods are similar to redefined methods of github.com/pkg/errors
// and provides similar (but in some cases not the same) interface.
// You can replace github.com/pkg/errors to github.com/nordborn/go-errow but not
// in another direction

package errow

import (
	"fmt"
	"github.com/pkg/errors"
	"path/filepath"
	"runtime"
	"strings"
)

// /home/.../.../go/src/very/long/path/mypackage/myfile.go -> mypackage/mypath.go
func shorterFilePath(fPath string) string {
	dirPath, fileName := filepath.Split(filepath.ToSlash(fPath))
	dirs := strings.Split(strings.Trim(dirPath, "/"), "/")
	return filepath.Join(dirs[len(dirs)-1], fileName)
}

func traceInfo(fPath string, line int) string {
	return fmt.Sprintf("<%v:%v>", shorterFilePath(fPath), line)
}

func getFileLine(skip int) (string, int) {
	_, fPath, line, ok := runtime.Caller(skip)
	if !ok {
		fPath = "???"
		line = 0
	}
	return fPath, line
}

func msg(v ...interface{}) string {
	fPath, line := getFileLine(3)
	return strings.Trim(fmt.Sprint(traceInfo(fPath, line), " ", fmt.Sprint(v...)), " ")
}

func msgf(format string, v ...interface{}) string {
	fPath, line := getFileLine(3)
	return strings.Trim(fmt.Sprint(traceInfo(fPath, line), " ", fmt.Sprintf(format, v...)), " ")
}

// New creates new error from text representation
// of given args and adds line no in the beginning of the message.
// New also records the stack trace at the point it was called.
// Example:
//
//   if err != nil {
// 	   return errow.New("my context: ", err)
//   }
//
func New(v ...interface{}) error {
	return errors.New(msg(v...))
}

// Newf creates new error from formatted text representation
// of given args and adds line no in the beginning of the message.
// Newf also records the stack trace at the point it was called.
// Example:
//
//   if err != nil {
// 	   return errow.Newf("val1=%v and val2=%v", val1, val2)
//   }
//
func Newf(format string, v ...interface{}) error {
	return errors.New(msgf(format, v...))
}

// Wrap returns an error annotating err with a stack trace at
// the point Wrap is called, and the supplied message.
// If err is nil, Wrap returns nil.
// Also, it adds file name and line of the point is called to
// the text of the error
// Example:
//
//   if err != nil {
// 	   return errow.Wrap(err)
//   }
//
//   if err2 != nil {
// 	   return errow.Wrap(err2, "important notice")
//   }
//
func Wrap(err error, v ...interface{}) error {
	return errors.Wrap(err, msg(v...))
}

// Wrapf returns an error annotating err with a stack
// trace at the point Wrapf is called, and the format specifier.
// If err is nil, Wrapf returns nil.
// Also, it adds file name and line of the point is called to
// the text of the error
// Example
//
//   if err != nil {
// 	   return errow.Wrap(err, "got err on vals: val1=%v val2=%v", val1, val2)
//   }
//
func Wrapf(err error, format string, v ...interface{}) error {
	return errors.Wrap(err, msgf(format, v...))
}
