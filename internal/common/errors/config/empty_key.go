package configErrors

import "fmt"

type EmptyKeyError struct {
	Key	string
}

func (e *EmptyKeyError) Error() string {
	return fmt.Sprintf("required key '%s", e.Key)
}
