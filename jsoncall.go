package anycall

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func JsonCall(fn interface{}, jsonParams ...string) []interface{} {
	fnValue := reflect.ValueOf(fn)
	fnType := fnValue.Type()

	if fnType.Kind() != reflect.Func {
		panic("fn must be a function")
	}

	if fnType.NumIn() != len(jsonParams) {
		panic(fmt.Sprintf("fn's params length %d, but jsonParams length %d", fnType.NumIn(), len(jsonParams)))
	}

	var params []reflect.Value
	for i := 0; i < fnType.NumIn(); i++ {
		paramValue := reflect.New(fnType.In(i))
		err := json.Unmarshal([]byte(jsonParams[i]), paramValue.Interface())
		if err != nil {
			panic(fmt.Sprintf("%d param unmarshal err:%v", i, err))
		}
		params = append(params, paramValue.Elem())
	}

	var returnResult []interface{}
	outputValues := fnValue.Call(params)
	for _, value := range outputValues {
		returnResult = append(returnResult, value.Interface())
	}
	return returnResult
}
