package errx

import (
	"errors"
	"fmt"
	"runtime"
)

func printCallerName() string {
	pc, _, _, _ := runtime.Caller(1)
	return runtime.FuncForPC(pc).Name()
}

// New err to check nil and locate error pos
func New(err error) error {
	// ignore it if nil
	if err != nil {
		pc, _, line, _ := runtime.Caller(1)
		callerName := runtime.FuncForPC(pc).Name()
		errStr := fmt.Sprintf("[%s-%d]-%v",
			callerName, line, err)
		err = errors.New(errStr)
	}
	return err
}
