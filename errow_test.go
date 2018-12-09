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
	errw := Wrap(err, "wrapped error")
	errww := Wrap(errw)
	log.Println(errw)
	log.Println(errww)
	fmt.Print("\n--- stack trace of the error ---\n\n")
	log.Printf("%+v", errww)
	fmt.Print("\n--- end of stack trace of the error ---\n\n")
	ok := strings.Contains(fmt.Sprint(errww), "errow/errow_test.go:")
	if !ok {
		t.Error(errw)
	}
}

func TestNew(t *testing.T) {
	errw := New("error: ", "err text")
	fmt.Println(errw)
	ok := strings.Contains(fmt.Sprint(errw), "errow/errow_test.go:")
	if !ok {
		t.Error(errw)
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
