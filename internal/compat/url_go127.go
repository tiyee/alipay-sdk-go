//go:build go1.27

package compat

import "net/url"

func CloneURL(u *url.URL) *url.URL        { return u.Clone() }
func CloneValues(v url.Values) url.Values { return v.Clone() }
