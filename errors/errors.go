package errors

import "github.com/pkg/errors"

func ErrMissingValue(key string) error  {
	return errors.Errorf("missing %s", key)
}

func ErrInvalidSecretEngine(name string) error  {
	return errors.Errorf("invalid secret engine %s", name)
}

func PrepareError(msg string, err error)  error {
	return errors.Wrap(err, msg)
}

func Wrap(msg string) error  {
	return errors.New(msg)
}
