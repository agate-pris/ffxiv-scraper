package main

import (
	"path"
	"testing"
)

func TestPathJoin(t *testing.T) {
	var p string
	p = path.Join(p, "./lodestone/playguide/db/item")
	p = path.Join(p, "1.html")
	if p != "lodestone/playguide/db/item/1.html" {
		t.Error(p)
	}
}
