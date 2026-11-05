//go:build !go1.27

package compat

import (
	"net/url"
	"slices"
)

func CloneURL(u *url.URL) *url.URL {
	if u == nil {
		return nil
	}
	cloned := *u
	if u.User != nil {
		user := *u.User
		cloned.User = &user
	}
	return &cloned
}

func CloneValues(v url.Values) url.Values {
	if v == nil {
		return nil
	}
	cloned := make(url.Values, len(v))
	for key, values := range v {
		cloned[key] = slices.Clone(values)
	}
	return cloned
}
