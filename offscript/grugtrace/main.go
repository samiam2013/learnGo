package main

import (
	"bytes"
	"errors"
	"fmt"
	"runtime/debug"

	"github.com/sirupsen/logrus"
)

func main() {
	if err := f(); err != nil {
		logrus.WithError(err).Error("failed to f()")
		fmt.Printf("%v\n", err)
	}
}

func f() error {
	err := errors.New("big fail. bad")
	return ErrorWithStackTrace(err)
}

func ErrorWithStackTrace(err error) error {
	b := debug.Stack()
	b = bytes.ReplaceAll(b, []byte("\t"), []byte(" "))
	lines := bytes.Split(b, []byte("\n"))
	lines = lines[5:]
	b = bytes.Join(lines, []byte(" \n "))
	fmt.Printf("%+v", string(b))
	return errors.Join(err, fmt.Errorf("%s", string(b)))
}
