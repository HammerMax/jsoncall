package anycall

import (
	"errors"
	"fmt"
)

type RequestA struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type RequestB struct {
	Is bool `json:"is"`
}

type Response struct {
	Code    int
	Message string
}

func Fn(a RequestA, b RequestB) (Response, error) {
	fmt.Printf("request a:%v\n", a)
	fmt.Printf("request b:%v\n", b)

	return Response{Code: 999, Message: "success"}, errors.New("new error")
}

func Example_jsonCall() {
	returnResult := JsonCall(Fn, `{"name": "hellp", "age": 20}`, `{"is": true}`)
	fmt.Printf("response:%v, err:%v", returnResult[0], returnResult[1])

	// Output:
	// request a:{hellp 20}
	// request b:{true}
	// response:{999 success}, err:new error
}
