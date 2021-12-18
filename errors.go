package onesignal

import "fmt"

type ApiError struct {
	Status        int    `json:"status"`
	InternalError string `json:"error"`
}

func (e ApiError) Error() string {
	return fmt.Sprintf("status: %d; error: %s", e.Status, e.InternalError)
}
