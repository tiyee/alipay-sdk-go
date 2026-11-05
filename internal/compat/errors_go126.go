//go:build go1.26

package compat

import "errors"

// AsType extracts a matching error without changing its cause chain.
func AsType[T error](err error) (T, bool) { return errors.AsType[T](err) }
