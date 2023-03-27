package BuilderPattern

import (
	"fmt"
	"testing"
)

func TestBuilder1(t *testing.T) {
	builder1 := &Builder1{}
	director1 := NewBuilder(builder1)
	director1.Construct()
	result := builder1.GetResult()
	if result != "123" {
		t.Fatalf("Builder1 fail expect 123 acture %s", result)
	}
	fmt.Println(result)
}

func TestBuilder2(t *testing.T) {
	builder2 := &Builder2{}
	director2 := NewBuilder(builder2)
	director2.Construct()
	result := builder2.GetResult()
	if result != 6 {
		t.Fatalf("Builder2 fail expect 6 acture %d", result)
	}
	fmt.Println(result)
}
