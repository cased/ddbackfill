package auditevents

import (
	"fmt"
	"reflect"
	"strings"
)

const Delimiter = "\u0000"

func Flatten(input map[string]interface{}) map[string][]string {
	f := flattened{
		input: input,
		table: map[string]map[string]bool{},
	}
	return f.process()
}

type flattened struct {
	input map[string]interface{}
	table map[string]map[string]bool
}

func (f flattened) process() map[string][]string {
	v := reflect.ValueOf(f.input)
	f.walk(v, "")

	// To ensure we don't return duplicate elements
	flat := map[string][]string{}
	for key, values := range f.table {
		for value := range values {
			flat[key] = append(flat[key], value)
		}
	}

	return flat
}

func (f flattened) walk(v reflect.Value, path string) {
	for v.Kind() == reflect.Ptr || v.Kind() == reflect.Interface {
		v = v.Elem()
	}
	switch v.Kind() {
	case reflect.Array, reflect.Slice:
		for i := 0; i < v.Len(); i++ {
			f.walk(v.Index(i), path)
		}
	case reflect.Map:
		iter := v.MapRange()
		for iter.Next() {
			key := iter.Key().String()
			nestedPath := key
			if path != "" {
				nestedPath = fmt.Sprintf("%s%s%s", path, Delimiter, key)
			}
			f.walk(iter.Value(), nestedPath)
		}
	case reflect.Invalid:
		// nil value
	default:
		f.buildPath(path, fmt.Sprintf("%v", v))
	}
}

func (f flattened) buildPath(path, value string) {
	// Haven't seen this path before
	if f.table[path] == nil {
		f.table[path] = map[string]bool{}
	}

	// We've already seen this path and value, skip...
	if f.table[path][value] {
		return
	}

	f.table[path][value] = true

	if !strings.Contains(path, Delimiter) {
		return
	}

	// issue/user/repository/name
	el := strings.Split(path, Delimiter)

	// name
	suffix := el[len(el)-1]

	// issue/user
	prefix := el[0 : len(el)-2]

	newPathParts := append(prefix, suffix)
	newPath := strings.Join(newPathParts, Delimiter)

	f.buildPath(newPath, value)
}
