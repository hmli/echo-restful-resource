package main

import (
	"flag"
	"os"
	"fmt"
	"go/parser"
	"go/token"
	"io/ioutil"
	"strings"
)

// resrc-maker -p /path/to/src -n photo -o PhotoResource

var tpl = `package %s

type %s struct {}

func (r *%s) Index(c echo.Context) error {
	panic("implement me")
}

func (r *%s) Create(c echo.Context) error {
	panic("implement me")
}

func (r *%s) Store(c echo.Context) error {
	panic("implement me")
}

func (r *%s) Show(c echo.Context) error {
	panic("implement me")
}

func (r *%s) Edit(c echo.Context) error {
	panic("implement me")
}

func (r *%s) Update(c echo.Context) error {
	panic("implement me")
}

func (r *%s) Destroy(c echo.Context) error {
	panic("implement me")
}

`

func main() {
	path := flag.String("p", "", "path to project file")
	name := flag.String("n", "", "Resource name")
	fullName := flag.String("o", "", "Resource struct name")

	flag.Parse()
	if *path == "" {
		wd, err := os.Getwd()
		path = &wd
		if err != nil {
			panic("Invalid path: " + err.Error())
		}
	}
	if *name == "" {
		*name = "Demo"
	}
	if *fullName == "" {
		*fullName = *name + "Resource"
	}

	var pname string
	files, _ := ioutil.ReadDir(*path)
	for _, fi := range files {
		if strings.HasSuffix(fi.Name(), ".go") {
			name, err := pkgName(*path + "/" + fi.Name())
			if err != nil {
				continue
			}
			pname = name
		}
	}
	if pname == "" {
		pname = pkgNameByPath(*path)
	}
	code := fmt.Sprintf(tpl, pname, *fullName, *fullName, *fullName, *fullName, *fullName, *fullName, *fullName, *fullName)
	outputName := strings.ToLower(*name + "_resource.go")
	f, err := os.Create(*path + "/" + outputName)
	if err != nil {
		panic("Failed to create resource file: " + err.Error())
	}
	defer f.Close()
	f.WriteString(code)
	return
}

func pkgName(file string) (string, error) {
	fset := token.NewFileSet()

	// parse the go soure file, but only the package clause
	astFile, err := parser.ParseFile(fset, file, nil, parser.PackageClauseOnly)
	if err != nil {
		return "", err
	}

	if astFile.Name == nil {
		return "", fmt.Errorf("no package name found")
	}
	return astFile.Name.Name, nil
}

func pkgNameByPath(path string) (pkgName string) {
	return strings.Replace(path, "-", "_", -1)
}