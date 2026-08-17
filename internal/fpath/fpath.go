package fpath

import "path"

func Ext(name string) string {
	// BUG: only last segment
	return path.Ext(name)
}
