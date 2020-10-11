package errors

import (
	stderrors "errors"

	pkgerrors "github.com/pkg/errors"
)

type Logger interface {
	Debugf(format string, v ...interface{})
}

var log Logger

func WithLogger(logger Logger) {
	if log == nil {
		log = logger
	}
}

func New(message string) error {
	err := pkgerrors.New(message)
	if log != nil {
		log.Debugf("New error: %+v", err)
	}
	return err
}

// Errorf formats according to a format specifier and returns the string
// as a value that satisfies error.
// Errorf also records the stack trace at the point it was called.
func Errorf(format string, args ...interface{}) error {
	err := pkgerrors.Errorf(format, args...)
	if log != nil {
		log.Debugf("Errorf error: %+v", err)
	}
	return err
}

// WithStack annotates err with a stack trace at the point WithStack was called.
// If err is nil, WithStack returns nil.
func WithStack(err error) error {
	if err == nil {
		return nil
	}
	err = pkgerrors.WithStack(err)
	if log != nil {
		log.Debugf("WithStack error: %+v", err)
	}
	return err
}

// Wrap returns an error annotating err with a stack trace
// at the point Wrap is called, and the supplied message.
// If err is nil, Wrap returns nil.
func Wrap(err error, message string) error {
	if err == nil {
		return nil
	}
	err = pkgerrors.Wrap(err, message)
	if log != nil {
		log.Debugf("Wrap error: %+v", err)
	}
	return err
}

// Wrapf returns an error annotating err with a stack trace
// at the point Wrapf is called, and the format specifier.
// If err is nil, Wrapf returns nil.
func Wrapf(err error, format string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	err = pkgerrors.Wrapf(err, format, args...)
	if log != nil {
		log.Debugf("Wrapf error: %+v", err)
	}
	return err
}

// WithMessage annotates err with a new message.
// If err is nil, WithMessage returns nil.
func WithMessage(err error, message string) error {
	if err == nil {
		return nil
	}
	err = pkgerrors.WithMessage(err, message)
	if log != nil {
		log.Debugf("WithMessage error: %+v", err)
	}
	return err
}

// WithMessagef annotates err with the format specifier.
// If err is nil, WithMessagef returns nil.
func WithMessagef(err error, format string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	err = pkgerrors.WithMessagef(err, format, args...)
	if log != nil {
		log.Debugf("WithMessage error: %+v", err)
	}
	return err
}

// Cause returns the underlying cause of the error, if possible.
// An error value has a cause if it implements the following
// interface:
//
//     type causer interface {
//            Cause() error
//     }
//
// If the error does not implement Cause, the original error will
// be returned. If the error is nil, nil will be returned without further
// investigation.
func Cause(err error) error { return pkgerrors.Cause(err) }

// for go 1.13
func Is(err, target error) bool { return stderrors.Is(err, target) }

func As(err error, target interface{}) bool { return stderrors.As(err, target) }

func Unwrap(err error) error { return stderrors.Unwrap(err) }
