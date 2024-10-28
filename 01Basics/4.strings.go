package main
import (
	"fmt"
	"reflect"
)

func main() {
	var str = "master"
	fmt.Println(len(str))
	fmt.Println(str)
	fmt.Println(reflect.TypeOf(str))

	var s1 = "hello "
	var s2 = "world"
	fmt.Println(s1 + s2)
}