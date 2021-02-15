package slices

import (
	"reflect"
)

// Index finds index of string in []string, otherwise returns -1
func Index(vs []string, t string) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}

// MapSliceToInt maps a []string to a []int
func MapSliceToInt(vs []string, f func(string, int) int) []int {
	vsm := make([]int, len(vs))
	for i, v := range vs {
		vsm[i] = f(v, i)
	}
	return vsm
}

// Any returns whether f returns true for any value.
func Any(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if f(v) {
			return true
		}
	}
	return false
}

// StringInSlice returns whether a is in list.
func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

// IntInSlice returns whether a is in list.
func IntInSlice(a int, list []int) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

// FilterInterface filters interface{} if function returns true for any value.
func FilterInterface(elements interface{}, cond func(interface{}) bool) interface{} {
	contentType := reflect.TypeOf(elements)
	contentValue := reflect.ValueOf(elements)

	newElements := reflect.MakeSlice(contentType, 0, 0)
	for i := 0; i < contentValue.Len(); i++ {
		if v := contentValue.Index(i); cond(v.Interface()) {
			newElements = reflect.Append(newElements, v)
		}
	}

	return newElements.Interface()
}

// IndexInterface returns index of interface{} in an interface{}, otherwise -1.
func IndexInterface(elements interface{}, t interface{}) int {
	contentValue := reflect.ValueOf(elements)

	for i := 0; i < contentValue.Len(); i++ {
		if i == t {
			return i
		}
	}
	return -1
}

// RemoveInterfaceDuplicates removes duplicates from []BookSearchResult and returns the new array.
func RemoveInterfaceDuplicates(elements interface{}) interface{} {
	contentType := reflect.TypeOf(elements)
	contentValue := reflect.ValueOf(elements)

	encountered := reflect.MakeMap(reflect.MapOf(contentType, reflect.TypeOf("bool")))
	result := reflect.MakeSlice(contentType, 0, 0)

	for i := 0; i < contentValue.Len(); i++ {
		if encountered.MapIndex(contentValue.Index(i)).Bool() {
			// Do not add duplicate.
		} else {
			encountered.SetMapIndex(contentValue.Index(i), reflect.ValueOf(true))
			result = reflect.Append(result, contentValue.Index(i))
		}
	}

	return result
}
