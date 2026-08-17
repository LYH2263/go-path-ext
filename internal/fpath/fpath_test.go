package fpath

import "testing"

func TestExt(t *testing.T) {
	cases := []struct {
		name string
		want string
	}{
		// compound extensions must be returned in full
		{"a.tar.gz", ".tar.gz"},
		{"a.tar.bz2", ".tar.bz2"},
		{"a.tar.xz", ".tar.xz"},
		{"a.tar.zst", ".tar.zst"},
		{"a.tar.lz", ".tar.lz"},
		{"a.tar.lzma", ".tar.lzma"},
		{"archive.tar.gz", ".tar.gz"},
		{"dir/a.tar.gz", ".tar.gz"},

		// single extensions still go through path.Ext
		{"a.gz", ".gz"},
		{"a.tar", ".tar"},
		{"noext", ""},
	}
	for _, c := range cases {
		if got := Ext(c.name); got != c.want {
			t.Errorf("Ext(%q) = %q, want %q", c.name, got, c.want)
		}
	}
}
