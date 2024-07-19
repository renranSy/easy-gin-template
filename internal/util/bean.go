package util

import (
	"reflect"
)

func BeanCopy[O any, T any](origin O, target T) T {
	originValue := reflect.ValueOf(origin).Elem()
	targetValue := reflect.ValueOf(target).Elem()

	for i := 0; i < originValue.NumField(); i++ {
		originFieldName := originValue.Type().Field(i).Name
		targetFieldValue := targetValue.FieldByName(originFieldName)
		if targetFieldValue.IsValid() {
			targetFieldValue.Set(originValue.Field(i))
		}
	}

	return target
}
