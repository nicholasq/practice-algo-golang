package arraystring

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestMergeStringsAlternatively(t *testing.T) {
	type testcase struct {
		input [2]string
		want  string
	}

	testCases := []testcase{
		{input: [2]string{"abc", "def"}, want: "adbecf"},
		{input: [2]string{"ac", "efgh"}, want: "aecfgh"},
		{input: [2]string{"", ""}, want: ""},
		{input: [2]string{"a", ""}, want: "a"},
		{input: [2]string{"", "b"}, want: "b"},
	}

	for _, c := range testCases {
		start := time.Now()
		got := MergeAlternately(c.input[0], c.input[1])
		end := time.Now()
		elapsed := end.Sub(start)
		if got != c.want {
			t.Errorf("MergeStringsAlternatively")
		}
		fmt.Printf("MergeStringsAlternatively() - elapsed time: %dns\n", elapsed.Nanoseconds())
	}
}

func TestGcdOfStrings(t *testing.T) {
	type testcase struct {
		input [2]string
		want  string
	}

	testCases := []testcase{
		{input: [2]string{"abcabc", "abc"}, want: "abc"},
		{input: [2]string{"ababab", "ab"}, want: "ab"},
		{input: [2]string{"leet", "code"}, want: ""},
	}

	for _, c := range testCases {
		start := time.Now()
		got := GcdOfStrings(c.input[0], c.input[1])
		end := time.Now()
		elapsed := end.Sub(start)
		if got != c.want {
			t.Errorf("Expected: %s, got: %s", c.want, got)
		}
		fmt.Printf("GcdOfStrings() - elapsed time: %dns\n", elapsed.Nanoseconds())
	}
}

func TestKidsWithCandies(t *testing.T) {

	type args struct {
		candies      []int
		extraCandies int
	}

	type testcase struct {
		input args
		want  []bool
	}

	testCases := []testcase{
		{input: args{[]int{2, 3, 5, 1, 3}, 3}, want: []bool{true, true, true, false, true}},
		{input: args{[]int{4, 2, 1, 1, 2}, 1}, want: []bool{true, false, false, false, false}},
		{input: args{[]int{12, 1, 12}, 10}, want: []bool{true, false, true}},
	}

	for _, c := range testCases {
		start := time.Now()
		got := KidsWithCandies(c.input.candies, c.input.extraCandies)
		end := time.Now()
		elapsed := end.Sub(start)
		if !reflect.DeepEqual(c.want, got) {
			t.Errorf("Expected: %v, got %v", c.want, got)
		}
		fmt.Printf("KidsWithCandies() - elapsed time: %dns\n", elapsed.Nanoseconds())
	}
}
