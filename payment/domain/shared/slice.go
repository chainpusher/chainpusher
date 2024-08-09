package shared

import (
	"reflect"
)

type Slice[T any] []T

type Sliceable[T any] interface {
}

func (s Slice[T]) Paginate(size int) []Slice[T] {
	var pages []Slice[T]
	for i := 0; i < len(s); i += size {
		end := i + size
		if end > len(s) {
			end = len(s)
		}
		pages = append(pages, s[i:end])
	}
	return pages
}

func IsPointerType[T any](v T) bool {
	k := reflect.TypeOf(v).Kind()
	return k == reflect.Ptr
}

func IsNil[T any](v T) bool {
	if !IsPointerType(v) {
		return false
	}
	return reflect.ValueOf(v).IsNil()
}
func (s Slice[T]) Add(v T) Slice[T] {

	if IsPointerType(v) && IsNil(v) {
		return s
	}
	return append(s, v)
}

func Map[T any, U any](slice Slice[T], fn func(T) U) Slice[U] {

	var mapped Slice[U]
	for _, v := range slice {
		mapped = append(mapped, fn(v))
	}
	return mapped
}

func (s Slice[T]) FlatMap(fn func(T) Slice[T]) Slice[T] {
	var flat Slice[T]
	for _, v := range s {
		flat = append(flat, fn(v)...)
	}
	return flat
}

func (s Slice[T]) Filter(fn func(T) bool) Slice[T] {
	var filtered Slice[T]
	for _, v := range s {
		if fn(v) {
			filtered = append(filtered, v)
		}
	}
	return filtered
}

func (s Slice[T]) GroupBy(fn func(T) interface{}) map[interface{}]Slice[T] {
	groups := make(map[interface{}]Slice[T])
	for _, v := range s {
		key := fn(v)
		groups[key] = append(groups[key], v)
	}
	return groups
}
