package AdapterPattern

import (
	"fmt"
	"testing"
)

var EXPECT = "adapter method"

func TestAdapter(t *testing.T) {
	target := NewAdaptee()
	adapter := NewAdapter(target)
	ret := adapter.Request()
	if ret != EXPECT {
		t.Fatalf("Expect: %s, Error: %s\n", EXPECT, ret)
	}
	fmt.Println(ret)
}
