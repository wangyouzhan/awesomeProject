package stringutil

import (
	"fmt"
	"testing"
)

func TestReverse(t *testing.T) {
	cases := []struct{in, want string}{
		{"Hello, world", "dlrow, olleH"},
		{"Hello, 世界", "世界 , olleH"},
		{"", ""},
	}
	for _, c := range cases{
		got := Reverse(c.in)
		if got != c.want{
			t.Errorf("Reverse(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

/**
	测试append
 */
func TestWang(t *testing.T)  {
	x := []int{1,2,3}
	x = append(x, 4, 5, 6)
	fmt.Println(x)
}