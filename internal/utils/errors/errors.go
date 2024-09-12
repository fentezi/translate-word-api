package errors

import "fmt"

var (
	ErrKeyNotFound  = fmt.Errorf("key not found")
	ErrKeyExists    = fmt.Errorf("key already exists")
	ErrKeyInvalid   = fmt.Errorf("key invalid")
	ErrValueInvalid = fmt.Errorf("value invalid")
)
