package configerrors

import "fmt"

//KeyError key doesn't exist
type KeyError struct {
	Key string
}

func (e *KeyError) Error() string {
	return fmt.Sprintf("Key %s does not exist in config", e.Key)
}
