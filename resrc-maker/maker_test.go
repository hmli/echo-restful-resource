package main

import (
	"testing"
	"os"
)

func Test_packageName(t *testing.T) {
	path, err := os.Getwd()
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	name, err := pkgName(path+"/maker.go")
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	if name != "main" {
		t.Log("Wrong package name, expect 'main', got '", name, "'" )
	}
	t.Log(name, err)
}
