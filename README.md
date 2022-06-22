# jsoncall

像动态语言（Python、PHP...）一样，在无法提前预知函数类型和参数类型的情况下，调用任意函数，并且传入任意参数类型  
例子 [jsoncall_test.go](https://github.com/HammerMax/jsoncall/blob/main/jsoncall_test.go)

```go
go get github.com/HammerMax/jsoncall
```

```go
func AnyFunc(a RequestA, b RequestB) (Response, error)

func Example_jsonCall() {
	returnResult := JsonCall(AnyFunc, `{"name": "hellp", "age": 20}`, `json param`)
	fmt.Printf("response:%v, err:%v", returnResult[0], returnResult[1])
}
```