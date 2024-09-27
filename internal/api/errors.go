package api

import "fmt"

type Error struct {
	Code    int
	Message string
	Details string
}

func (e Error) Error() string {
	return fmt.Sprintf("API error[%d]: %s; (%s)", e.Code, e.Message, e.Details)
}
