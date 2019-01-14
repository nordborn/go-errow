[![Build Status](https://travis-ci.org/nordborn/go-errow.svg?branch=master)](https://travis-ci.org/nordborn/go-errow)
[![Code Coverage](https://codecov.io/gh/nordborn/go-errow/branch/master/graph/badge.svg)](https://codecov.io/gh/nordborn/go-errow/branch/master/graph/badge.svg)
[![GoDoc](https://godoc.org/github.com/nordborn/go-errow?status.svg)](https://godoc.org/github.com/nordborn/go-errow)

**Package errow is a wrapper for errors in Go (Golang) that makes them more informative**.


Package errow creates error with context and adds line number as a prefix
of the text representation of the given error
`<folder/file.go:15>: error message`.

Example:

```Go
// regular error
err := errors.New("my error")
// wrapped error
errW := Wrap(err, "wrapped error")
fmt.Println(errW)
// Output: <mypackagefolder/main.go:47> wrapped error: my error
```

Also, it puts the error to a stack of errors thanks to github.com/pkg/errors
and you can also print the stack via formatted output `"%+v"`.
Names of methods are similar to redefined methods of github.com/pkg/errors
and provide similar (but in some cases not the same) interface.
You can replace github.com/pkg/errors to github.com/nordborn/go-errow but not
in another direction

Provided functions:

```Go

// New creates new error from text representation
// of given args and adds line no in the beginning of the message.
// New also records the stack trace at the point it was called.
// Example:
//
//   if err != nil {
// 	   return errow.New("my context: ", err)
//   }
//
func New(v ...interface{}) error {}

// Newf creates new error from formatted text representation
// of given args and adds line no in the beginning of the message.
// Newf also records the stack trace at the point it was called.
// Example:
//
//   if err != nil {
// 	   return errow.Newf("val1=%v and val2=%v", val1, val2)
//   }
//
func Newf(format string, v ...interface{}) error {}

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
func Wrap(err error, v ...interface{}) error {}

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
func Wrapf(err error, format string, v ...interface{}) error {}
```

Example output of `go test`:

```
2018/12/09 15:55:55 <go-errow/errow_test.go:13> wrapped error: my error
2018/12/09 15:55:55 <go-errow/errow_test.go:14>: <go-errow/errow_test.go:13> wrapped error: my error

--- stack trace of the error ---

2018/12/09 15:55:55 my error
<go-errow/errow_test.go:13> wrapped error
github.com/nordborn/go-errow.Wrap
        /media/.../go-errow/errow.go:95
github.com/nordborn/go-errow.TestWrap
        /media/.../go-errow/errow_test.go:13
testing.tRunner
        /home/.../testing/testing.go:827
runtime.goexit
        /home/.../runtime/asm_amd64.s:1333
<go-errow/errow_test.go:14>
github.com/nordborn/go-errow.Wrap
        /media/.../go-errow/errow.go:95
github.com/nordborn/go-errow.TestWrap
        /media/.../go-errow/errow_test.go:14
testing.tRunner
        /home/.../testing/testing.go:827
runtime.goexit
        /home/.../runtime/asm_amd64.s:1333
        
--- end of stack trace of the error ---

<go-errow/errow_test.go:26> error: err text
<go-errow/errow_test.go:36> error: [1 2 3]
PASS
ok      github.com/nordborn/go-errow    0.003s

```

Godoc: https://godoc.org/github.com/nordborn/go-errow
