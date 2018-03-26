package main

import (
	"reflect"
	"fmt"
)

func main()  {
	sm1 := struct {
		age int
		m map[string]string
	}{
		age: 11,
		m:map[string]string{
			"a": "1",
		},
	}

	sm2 := struct {
		age int
		m map[string]string
	}{
		age: 11,
		m: map[string]string{
			"a": "1",
		},
	}

	// 不能这样比较，因为拥有map，slice属性的struct不能使用==运算符进行比较
	//if sm1 == sm2 {
	//
	//}

	if reflect.DeepEqual(sm1, sm2){
		fmt.Println("sm1 deep equal sm2")
	} else {
		fmt.Println("sm1 not deep equal sm2")
	}
}
