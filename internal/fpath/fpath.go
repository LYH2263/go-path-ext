package fpath

import (
	"path"
	"strings"
)

func Ext(name string) string {
	base := path.Base(name)
	if strings.HasSuffix(strings.ToLower(base), ".tar.gz") {
		return ".tar.gz"
	}
	return path.Ext(base)
}
