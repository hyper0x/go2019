package main

import (
	"errors"
	"fmt"
	"golang.org/x/xerrors"
)

// DetailedError 是一个有错误详情的错误类型。
type DetailedError struct {
	msg string
}

// Error 会返回 error 的信息。
func (de DetailedError) Error() string {
	return de.msg
}

// WrappedError 是一个可包装其他错误的错误类型。
type WrappedError struct {
	msg   string
	inner error
}

// Error 会返回 error 的信息。
func (we WrappedError) Error() string {
	return we.msg
}

// Unwrap 会返回被包装的 error 值。
func (we WrappedError) Unwrap() error {
	return we.inner
}

// FormattedError 是可暴露内部错误信息的错误类型。
type FormattedError struct {
	msg   string
	inner error
}

// Error 会返回 error 的信息。
func (fe FormattedError) Error() string {
	return fe.msg
}

// Unwrap 会返回被包装的 error 值。
func (fe FormattedError) Unwrap() error {
	return fe.inner
}

// Format 会打印格式化后的错误值。
func (fe FormattedError) Format(f fmt.State, c rune) {
	xerrors.FormatError(fe, f, c)
}

// FormatError 会返回错误链中的下一个错误值。
func (fe FormattedError) FormatError(p xerrors.Printer) (next error) {
	p.Print(fe.msg)
	return fe.inner
}

func demoForWrapper() {
	err1_1 := errors.New("unsupported operation")
	err1_2 := WrappedError{
		msg:   "operation failed",
		inner: err1_1,
	}
	fmt.Printf("Message(outer error): %v\n", err1_2)
	fmt.Printf("Message(inner error): %v\n\n", errors.Unwrap(err1_2))
}

func demoForFuncs() {
	err2_1 := DetailedError{
		msg: "unsupported operation",
	}
	err2_2 := WrappedError{
		msg:   "operation failed",
		inner: err2_1,
	}
	fmt.Printf("err2_1: %#v\n", err2_1)
	fmt.Printf("err2_2: %#v\n", err2_2)
	fmt.Printf("errors.Is(err2_2, err2_1): %v\n",
		errors.Is(err2_2, err2_1))
	fmt.Printf("errors.As(err2_2, &DetailedError{}): %v\n\n",
		errors.As(err2_2, &DetailedError{}))
}

func demoForXerrors() {
	err3_1 := DetailedError{
		msg: "unsupported operation",
	}
	err3_2 := WrappedError{
		msg:   "operation failed",
		inner: err3_1,
	}
	err3_3 := FormattedError{
		msg:   "operation error",
		inner: err3_2,
	}
	err3_4 := fmt.Errorf("something wrong: %w", err3_3)
	fmt.Printf("Error: %v\n", err3_4)
}

func main() {
	demoForWrapper()
	demoForFuncs()
	demoForXerrors()
}
