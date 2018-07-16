package main

import (
	"encoding/json"
	"fmt"
)

type object struct {
	A string
	B string
	C string
}

type req struct {
	Array   []int
	Boolean bool
	Null    string
	Number  int
	Object  object
	String  string
}

func main() {

	r := req{
		[]int{1, 2, 3},
		true,
		"null",
		123,
		object{"c", "d", "e"},
		"Hello",
	}

	fmt.Println("Hello, playground ")

	m, err := jsonMarshal(r)

	if err != nil {

		fmt.Println("Json Marshal error", err)
		return
	}

	fmt.Println("The marshalled output is", string(m))

	um := new(req)

	if err := JsonUnmarshal(m, um); err != nil {
		fmt.Println("Json UnMarshal error", err)
		return

	}

}

func jsonUnmarshal(m []byte, r interface{}) (err error) {

	if err = json.Unmarshal(m, r); err != nil {

		return
	}

	fmt.Println("The unmarshalled output is", r)

	return
}

func jsonMarshal(r interface{}) (res []byte, err error) {

	res, err = json.Marshal(r)

	if err != nil {

		return
	}

	return
}
