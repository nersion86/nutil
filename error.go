package nutil

import (
	"errors"
	"fmt"
)

// NewBasicError = Create New Error.
func NewBasicError(msg string, args ...any) error {
	errMsg := fmt.Sprintf(msg, args...)
	return errors.New(errMsg)
}

// RecoverPanic = Panic Recovery Func.
func RecoverPanic() {
	if p := recover(); p != nil {
		err, ok := p.(error)
		if !ok {
			err = NewBasicError("%v", p)
		}
		fmt.Println(err)
	}
}

func SafeGo(f func()) {
	go func() {
		defer RecoverPanic()
		f()
	}()
}
