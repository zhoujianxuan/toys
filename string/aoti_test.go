package string

import (
	"fmt"
	"testing"
)

func TestAoti(t *testing.T) {
	s := "   -42"
	r := myAtoi(s)
	fmt.Println(r)
}
