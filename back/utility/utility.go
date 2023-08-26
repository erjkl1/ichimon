package utility

import "reflect"

func CopyFields(source interface{}, target interface{}) {
    sourceValue := reflect.ValueOf(source)
    targetValue := reflect.ValueOf(target).Elem()

    for i := 0; i < sourceValue.NumField(); i++ {
        sourceField := sourceValue.Field(i)
        fieldName := sourceValue.Type().Field(i).Name

        targetField := targetValue.FieldByName(fieldName)
        if targetField.IsValid() && targetField.Type() == sourceField.Type() {
            targetField.Set(sourceField)
        }
    }
}