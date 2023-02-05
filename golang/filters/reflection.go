package filters

import (
	"reflect"
)

// FilterReflection filters a slice of any using reflection
func FilterReflection(items any, pred any) any {
	contentValue := reflect.ValueOf(items)
	pFunc := reflect.ValueOf(pred)
	newContent := reflect.MakeSlice(reflect.TypeOf(items), 0, 0)
	for i := 0; i < contentValue.Len(); i++ {
		item := contentValue.Index(i)
		if pFunc.Call([]reflect.Value{item, reflect.ValueOf(i)})[0].Bool() {
			newContent = reflect.Append(newContent, item)
		}
	}
	return newContent.Interface()
}
