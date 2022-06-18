package util

import "reflect"

func Merge[T any](base, patch T) (merged T) {
	structFields := reflect.VisibleFields(reflect.TypeOf(base))
	patchValue := reflect.ValueOf(patch)

	merged = base
	mergedValue := reflect.ValueOf(&merged).Elem()

	patchFields(structFields, patchValue, mergedValue)
	return merged
}

func patchFields(structFields []reflect.StructField, patchValue reflect.Value, mergedValue reflect.Value) {
	for _, field := range structFields {
		patchField := patchValue.FieldByIndex(field.Index)
		mergedField := mergedValue.FieldByIndex(field.Index)
		patchIfNonZero(patchField, mergedField)
	}
}

func patchIfNonZero(patchField reflect.Value, mergedField reflect.Value) {
	if !patchField.IsZero() {
		switch patchField.Kind() {
		case reflect.String:
			mergedField.SetString(patchField.String())
			break
		case reflect.Int:
			mergedField.SetInt(patchField.Int())
			break
		case reflect.Bool:
			mergedField.SetBool(patchField.Bool())
			break
		default:
			break
		}
	}
}
