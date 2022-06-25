package main

import (
	"log"

	"fbnoi.com/framework/annotation/direct/route"
)

func compareSlice(arr1, arr2 []string) bool {
	if len(arr1) != len(arr2) {
		return false
	}

	for i := 0; i < len(arr1); i++ {
		if arr1[i] != arr2[2] {
			return false
		}
	}

	return true
}

func main() {
	r, err := route.ParseRoute("@Route(/post/:id, name=post_edit, methods=[POST])")
	if err != nil {
		log.Fatal(err)
	}
	if r.Path != "/post/:id" {
		log.Fatalf("expected %s, get %s", "/post/:id", r.Path)
	}

	if r.GetName() != "post_edit" {
		log.Fatalf("expected %s, get %s", "post_edit", r.GetName())
	}

	if !compareSlice(r.GetMethod(), []string{"POST"}) {
		log.Fatalf("expected %s, get %s", "[POST]", r.GetMethod())
	}
}
