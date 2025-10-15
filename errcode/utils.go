package errcode

import "errors"

func Is(err, target error) bool             { return errors.Is(err, target) }
func As(err error, target interface{}) bool { return errors.As(err, target) }
func Unwrap(err error) error {
	return errors.Unwrap(err)
}


func ToCodeErr(err error, defaultError ... error) error {
	var e *errCode
	if As(err, &e){
		return e
	}
	if len(defaultError) > 0 {
		return defaultError[0]
	}
	return defaultErr
}