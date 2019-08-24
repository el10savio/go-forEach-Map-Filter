package main

import (
	"fmt"
	"reflect"
)

func Filter(slice interface{}, sliceFn interface{}) interface{} {
	reflectedSlice := reflect.ValueOf(slice)
	reflectedsliceFn := reflect.ValueOf(sliceFn)
	returnSlice := make([]interface{}, reflectedSlice.Len())

	for index := 0; index < reflectedSlice.Len(); index++ {
		element := reflectedSlice.Index(index)
		if reflectedsliceFn.Call([]reflect.Value{element})[0].Interface() == true {
			returnSlice = append(returnSlice, element.Interface())
		}
	}

	return returnSlice[reflectedSlice.Len():]
}

func Map(slice interface{}, sliceFn interface{}) interface{} {
	reflectedSlice := reflect.ValueOf(slice)
	reflectedsliceFn := reflect.ValueOf(sliceFn)
	returnSlice := make([]interface{}, reflectedSlice.Len())

	for index := 0; index < reflectedSlice.Len(); index++ {
		element := reflectedSlice.Index(index)
		returnSlice = append(returnSlice, reflectedsliceFn.Call([]reflect.Value{element})[0].Interface())
	}

	return returnSlice[reflectedSlice.Len():]
}

func forEach(slice interface{}, sliceFn interface{}) {
	reflectedSlice := reflect.ValueOf(slice)
	reflectedsliceFn := reflect.ValueOf(sliceFn)

	for index := 0; index < reflectedSlice.Len(); index++ {
		element := reflectedSlice.Index(index)
		reflectedsliceFn.Call([]reflect.Value{element})[0].Interface()
	}
}

func addElem(element int) int {
	return element + 1
}

func isEven(element int) bool {
	if element%2 == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	array := []int{10, 2, 37, 40}

	arrayN := Map(array, addElem)
	fmt.Println(arrayN)

	arrayF := Filter(array, isEven)
	fmt.Println(arrayF)
}
