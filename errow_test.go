package errow

import (
	"errors"
	"fmt"
	"log"
	"strings"
	"testing"
)

func TestWrap(t *testing.T) {
	err := errors.New("my error")
	errW := Wrap(err, "wrapped error")
	errWW := Wrap(errW)
	log.Println(errW)
	log.Println(errWW)
	fmt.Print("\n--- stack trace of the error ---\n\n")
	log.Printf("%+v", errWW)
	fmt.Print("\n--- end of stack trace of the error ---\n\n")
	ok := strings.Contains(fmt.Sprint(errWW), "errow/errow_test.go:")
	if !ok {
		t.Error(errW)
	}
}

func TestNew(t *testing.T) {
	errW := New("error: ", "err text")
	fmt.Println(errW)
	ok := strings.Contains(fmt.Sprint(errW), "errow/errow_test.go:")
	if !ok {
		t.Error(errW)
	}
}

func TestNewf(t *testing.T) {
	errPayload := []int{1, 2, 3}
	errw := Newf("error: %v", errPayload)
	fmt.Println(errw)
	ok := strings.Contains(fmt.Sprint(errw), "errow/errow_test.go:")
	if !ok {
		t.Error(errw)
	}
}

func ExampleWrap() {
	err := errors.New("my error")
	errW := Wrap(err, "wrapped error")
	fmt.Println(errW)
	// Output: <go-errow/errow_test.go:47> wrapped error: my error
}
