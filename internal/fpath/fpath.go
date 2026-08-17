package fpath

import (
	"path"
	"strings"
)

// compoundExts lists known multi-part extensions. They are matched as
// suffixes of the file name's final path element, so a name like
// "a.tar.gz" yields ".tar.gz" instead of just ".gz".
var compoundExts = []string{
	".tar.gz",
	".tar.bz2",
	".tar.xz",
	".tar.zst",
	".tar.lz",
	".tar.lzma",
	".tar.lz4",
	".tar.br",
	".tar.sz",
	".tar.Z",
}

// Ext returns the file name extension used by name.
//
// Unlike path.Ext, which only returns the final dot-suffix, Ext recognises
// common compound extensions such as ".tar.gz" and returns them in full.
func Ext(name string) string {
	base := name
	if i := strings.LastIndexByte(name, '/'); i >= 0 {
		base = name[i+1:]
	}
	for _, ext := range compoundExts {
		// require a non-empty base name before the compound extension so
		// that hidden files like ".tar.gz" are not swallowed whole
		if len(base) > len(ext) && strings.HasSuffix(base, ext) {
			return ext
		}
	}
	return path.Ext(name)
}
