package test

import (
	"testing"

	"fbnoi.com/framework/annotation/direct/route"
)

func compareSlice(arr1, arr2 []string) bool {
	if len(arr1) != len(arr2) {
		return false
	}

	for i := 0; i < len(arr1); i++ {
		if arr1[i] != arr2[i] {
			return false
		}
	}

	return true
}

func TestParseRoute(t *testing.T) {
	r, err := route.ParseRoute("@Route(/post/:id, name=post_edit, methods=[POST])")
	if err != nil {
		t.Error(err)
	}
	if r.Path != "/post/:id" {
		t.Errorf("expected %s, get %s", "/post/:id", r.Path)
	}

	if r.Name != "post_edit" {
		t.Errorf("expected %s, get %s", "post_edit", r.Name)
	}

	if !compareSlice(r.Methods, []string{"POST"}) {
		t.Errorf("expected %s, get %s", "[POST]", r.Methods)
	}
}

func TestParseFilter(t *testing.T) {
	f, err := route.ParseFilter("@Filter(/post/*, order=1)")
	if err != nil {
		t.Error(err)
	}
	if f.Path != "/post/*" {
		t.Errorf("expected %s, get %s", "/post/*", f.Path)
	}

	if f.Order != 1 {
		t.Errorf("expected %d, get %d", 1, f.Order)
	}
}
