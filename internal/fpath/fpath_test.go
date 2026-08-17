package fpath

import "testing"

func TestExtTarGz(t *testing.T) {
	if Ext("a.tar.gz") != ".tar.gz" {
		t.Fatalf("%q", Ext("a.tar.gz"))
	}
}
