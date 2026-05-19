package validation

import "slices"

func In[T comparable](v T, opts ...T) bool {
	return slices.Contains(opts, v)
}
