package route

import (
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

type Route struct {
	Path       string
	Name       string
	Methods    []string
	HandleFunc string
	Filters    []*Filter
}

type Filter struct {
	Path  string
	Func  string
	Order int
}

// @Filter(/post/*, order=1)
func ParseFilter(anno string) (*Filter, error) {
	if !strings.HasPrefix(anno, "@Filter") {
		return nil, errors.Errorf("Unrecognized direct %s", anno)
	}
	anno = strings.TrimSuffix(strings.TrimPrefix(anno, "@Filter("), ")")

	splits := strings.Split(anno, ",")
	f := &Filter{}
	f.Path = splits[0]

	for i := 1; i < len(splits); i++ {
		str := strings.ReplaceAll(splits[i], " ", "")
		kv := strings.Split(str, "=")
		switch kv[0] {
		case "order":
			f.Order, _ = strconv.Atoi(kv[1])
		default:
			return nil, errors.Errorf("Unrecognized Filter property %s", kv[0])
		}
	}
	return f, nil
}

func ParseRoute(anno string) (*Route, error) {
	if !strings.HasPrefix(anno, "@Route") {
		return nil, errors.Errorf("Unrecognized direct %s", anno)
	}
	anno = strings.TrimSuffix(strings.TrimPrefix(anno, "@Route("), ")")

	splits := strings.Split(anno, ",")
	r := &Route{}
	r.Path = splits[0]
	for i := 1; i < len(splits); i++ {
		str := strings.ReplaceAll(splits[i], " ", "")
		kv := strings.Split(str, "=")
		switch kv[0] {
		case "name":
			r.Name = kv[1]
		case "methods":
			r.Methods = strToArr(kv[1])
		default:
			return nil, errors.Errorf("Unrecognized Route property %s", kv[0])
		}
	}
	return r, nil
}

// [bar, foo] => []string{"bar", "foo"}
func strToArr(str string) []string {
	str = strings.TrimSuffix(strings.TrimPrefix(str, "["), "]")
	str = strings.ReplaceAll(str, " ", "")
	return strings.Split(str, ",")
}
