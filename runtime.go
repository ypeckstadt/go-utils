package utils

import (
	"reflect"
	"time"
)

// HasNilValue checks if the provided value structs contains a nil value
func HasNilValue(values ...interface{}) bool {
	for i := 0; i < len(values); i++ {
		value := reflect.ValueOf(values[i])
		for j := 0; j < value.NumField(); j++ {
			field := value.Field(j)
			switch field.Kind() {
			case reflect.Ptr, reflect.Map, reflect.Array, reflect.Chan, reflect.Slice:
				if field.IsNil() {
					return true
				}
			}
		}
	}
	return false
}

// KeepRunning blocks the main thread from closing
func KeepRunning() {
	done := make(chan bool)
	go forever()
	<-done // Block forever
}

func forever() {
	for {
		time.Sleep(time.Second)
	}
}
