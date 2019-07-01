package code

import (
	"fmt"
	"reflect"
)

func RefInp() {
	i := 0
	v := reflect.ValueOf(&i)
	t := reflect.TypeOf(i)
	v = v.Elem()
	fmt.Printf("%v\n", v.CanSet())
	fmt.Printf("%v\n%v\n", v, t)
}
