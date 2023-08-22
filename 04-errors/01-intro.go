package main

import (
	"errors"
	"fmt"
)

var DummyError error = errors.New("dummy error")

func main() {

	e := errorFn(false)
	if e == nil {
		fmt.Println("fn executed successfully")
	} else if e == DummyError {
		fmt.Println("Dummy error occurred")
	} else {
		fmt.Printf("unknown error occured in fn. err : %v\n", e)
	}
}

func errorFn(pass bool) error {
	if pass {
		return nil
	}
	// err := DummyError
	err := errors.New("some random error")
	return err
}
