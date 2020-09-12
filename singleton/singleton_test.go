package singleton

import (
	"fmt"
	"testing"
)

func TestNewSingleton(t *testing.T) {
	singleton := GetSingleton()
	singleton.data = 100
	fmt.Println(singleton.data) // 10

	singleton2 := GetSingleton()
	fmt.Println(singleton2.data) // 10

	singleton3 := GetSingleton()
	singleton4 := GetSingleton()
	if singleton3 == singleton4 {
		fmt.Println("equal")
	}else{
		fmt.Println("not equal")
	}
}
