package main

import (
	"net/url"
	"testing"
)

func TestUrl(t *testing.T) {
	u, e := url.Parse("https://jp.finalfantasyxiv.com/lodestone/playguide/db/item/?category2=1&page=1")
	if e != nil {
		t.Fatal(e)
	}
	if u.Scheme != "https" {
		t.Error(u.Scheme)
	}
	{
		s := u.String()
		if s != "https://jp.finalfantasyxiv.com/lodestone/playguide/db/item/?category2=1&page=1" {
			t.Error(u)
		}
	}
	if u.Host != "jp.finalfantasyxiv.com" {
		t.Error(u.Host)
	}
	{
		h := u.Hostname()
		if h != "jp.finalfantasyxiv.com" {
			t.Error(h)
		}
	}
	if u.Path != "/lodestone/playguide/db/item/" {
		t.Error(u.Path)
	}
	{
		q := u.Query()
		c2 := q["category2"]
		p := q["page"]
		if len(c2) != 1 || c2[0] != "1" {
			t.Error(c2)
		}
		if len(p) != 1 || p[0] != "1" {
			t.Error(p)
		}
	}
}
