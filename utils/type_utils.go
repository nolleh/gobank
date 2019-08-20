package utils

import (
	"fmt"
	"reflect"
)

func IsInstanceOf(objectPtr, typePtr interface{}) bool {
	return reflect.TypeOf(objectPtr) == reflect.TypeOf(typePtr)
}

func Stringify(obj interface{}) string {
	return fmt.Sprint(obj)
}