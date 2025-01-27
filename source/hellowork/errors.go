package hellowork

import (
	"fmt"
)

var (
	ErrNewRequest = func(err error) error {
		return fmt.Errorf("failed to build HTTP request: %w", err)
	}
	ErrDoRequest = func(err error) error {
		return fmt.Errorf("failed to execute HTTP request: %w", err)
	}
	ErrReadBody = func(method, url string, err error) error {
		return fmt.Errorf("failed to Read HTTP response body '%s:%s': %w", method, url, err)
	}
	ErrEmptyBody = func(method, url string) error {
		return fmt.Errorf("unexpected empty body response '%s:%s'", method, url)
	}
	ErrUnmarshal = func(method, url string, err error) error {
		return fmt.Errorf("error when unmarshaling body '%s:%s': %w", method, url, err)
	}
	ErrServerHTTPError = func(code int) error {
		return fmt.Errorf("error server returned an http error[%d]", code)
	}
)
