package context

import "reflect"

type Scope struct {
	beans map[string]interface{}
}

func Add[T any](scope *Scope, bean T) {
	scope.beans[reflect.TypeOf(bean).String()] = bean
}

func Get[T any](scope *Scope, kind interface{}) T {
	return scope.beans[reflect.TypeOf(kind).String()].(T)
}

func NewScope() *Scope {
	return &Scope{
		beans: make(map[string]interface{}),
	}
}
