package route

import (
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

type Route struct {
	Package string
	Path    string
	Name    string
	Methods []string
	Func    string
	Filters *Filters
}

type Filters []*Filter

type Filter struct {
	Package string
	Path    string
	Func    string
	Order   int
}

func (fs *Filters) Add(e *Filter) {
	*fs = append(*fs, e)
}

func (fs *Filters) Len() int {
	return len(*fs)
}

func (fs *Filters) Less(i, j int) bool {
	return (*fs)[i].Order > (*fs)[j].Order
}

func (fs *Filters) Swap(i, j int) {
	tmp := (*fs)[i]
	(*fs)[i] = (*fs)[j]
	(*fs)[j] = tmp
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
	r := &Route{
		Filters: &Filters{},
	}
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
