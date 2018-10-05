package sitemap

import (
	"strconv"
	"testing"
)

func TestAddNum(t *testing.T) {
	var exp string
	s := "sample"
	i := 99
	exp = s + "_" + strconv.Itoa(i)
	if addNum(s, i) != exp {
		t.Fatalf("addNum(%s,%v) is expected to return %s.", s, i, exp)
	}
}

func TestMin(t *testing.T) {
	var exp *int
	x := 1
	y := 2
	if x > y {
		exp = &y
	} else {
		exp = &x
	}
	if min(x, y) != *exp {
		t.Fatalf("min(%v,%v) is expected to return %v.", x, y, *exp)
	}
}
