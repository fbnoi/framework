package scanner

import (
	"go/ast"
	"go/parser"
	"go/token"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"

	"fbnoi.com/framework/annotation/route"
	"github.com/pkg/errors"
)

var (
	routes     []*route.Route
	filters    []*route.Filter
	direct_rgx = regexp.MustCompile(`^@[a-zA-Z]+\([^\(\)]+\)$`)
	osType     = os.Getenv("GOOS")
)

func ScanFolders(dirs ...string) (err error) {
	for _, dir := range dirs {
		dir, err = filepath.Abs(dir)
		if err != nil {
			return err
		}
		var (
			files []fs.DirEntry
			i     fs.FileInfo
			path  string
		)
		files, err = os.ReadDir(dir)
		if err != nil {
			return err
		}
		for _, f := range files {
			if f.Type().IsRegular() && !f.IsDir() {
				i, err = f.Info()
				if err != nil {
					return
				}
				if strings.HasSuffix(i.Name(), ".go") {
					path, err = concatFilePath(dir, i.Name())
					if err != nil {
						return
					}
					err = Scan(path)
					if err != nil {
						return
					}
				}
			}
		}
	}
	for _, filter := range filters {
		s := strings.Trim(filter.Path, "*")
		lStar, rStar := strings.HasPrefix(filter.Path, "*"), strings.HasSuffix(filter.Path, "*")
		for _, route := range routes {
			if (lStar && rStar && strings.Contains(route.Path, s)) ||
				(lStar && strings.HasSuffix(route.Path, s)) ||
				(rStar && strings.HasPrefix(route.Path, s)) ||
				s == route.Path {
				route.Filters.Add(filter)
			}
		}
	}

	for _, route := range routes {
		sort.Sort(route.Filters)
	}
	return nil
}

func Scan(path string) (err error) {
	fSet := token.NewFileSet()
	f, err := parser.ParseFile(fSet, path, nil, parser.ParseComments)
	if err != nil {
		return
	}
	var _package string
	ast.Inspect(f, func(n ast.Node) bool {
		switch _node := n.(type) {
		case *ast.File:
			_package = _node.Name.Name
		case *ast.FuncDecl:
			if _node.Doc != nil {
				for _, c := range _node.Doc.List {
					direct := trimAnnotation(c.Text)
					switch {
					case strings.HasPrefix(direct, "@Route"):
						var r *route.Route
						r, err = route.ParseRoute(trimAnnotation(c.Text))
						r.Package = _package
						r.Func = _node.Name.Name
						routes = append(routes, r)
					case strings.HasPrefix(direct, "@Filter"):
						var f *route.Filter
						f, err = route.ParseFilter(trimAnnotation(c.Text))
						f.Package = _package
						f.Func = _node.Name.Name
						filters = append(filters, f)
					}
					if err != nil {
						return false
					}
				}
			}
		}
		return true
	})
	return
}

func concatFilePath(dir, baseName string) (string, error) {
	dir, err := filepath.Abs(dir)
	if err != nil {
		return "nil", err
	}
	if osType == "windows" {
		return dir + "\\" + baseName, nil
	} else if osType == "linux" {
		return dir + "/" + baseName, nil
	}

	return "", errors.Errorf("unrecognized os %s", baseName)
}

func trimAnnotation(anno string) string {
	anno = strings.TrimSpace(strings.TrimPrefix(anno, "//"))
	if !direct_rgx.MatchString(anno) {
		return ""
	}
	return anno
}
