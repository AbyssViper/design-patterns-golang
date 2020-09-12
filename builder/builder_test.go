package builder

import (
	"fmt"
	"testing"
)

func TestNewDirector(t *testing.T) {
	builder := &StringBuilder{}
	director := NewDirector(builder)
	director.Generate()
	fmt.Println(builder.GetResult())
}