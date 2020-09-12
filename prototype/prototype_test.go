package prototype

import (
	"fmt"
	"testing"
)

func TestNewPrototype(t *testing.T) {
	manager := NewPrototype()
	t1 := &Type1{name: "type1"}
	manager.Set("t1", t1)

	t11 := manager.Get("t1")
	t12 := t11.Clone()	// 复制

	if t11 == t12 {
		fmt.Println("copy")
	}else{
		fmt.Println("deep copy")
	}

}
