package FacadePattern

import (
	"fmt"
	"testing"
)

var EXPECT = "A module running\nB module running\n"

func TestNewAPI(t *testing.T) {
	api := NewAPI()
	ret := api.Test()
	if ret != EXPECT {
		t.Fatalf("Expect: %s, Error: %s\n", EXPECT, ret)
	}
	fmt.Println(api.Test())
}
