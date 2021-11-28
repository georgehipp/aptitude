package number

import (
	"fmt"
	"testing"
)

func Test_containsDuplicate(t *testing.T) {
	var tests = []struct {
		a, b, c int
		want    bool
	}{
		{1, 1, 0, true},
		{1, 0, 1, true},
		{0, 1, 1, true},
		{0, 1, 2, false},
	}

	for _, tt := range tests {
		test := fmt.Sprintf("%d, %d, %d, %t", tt.a, tt.b, tt.c, tt.want)
		t.Run(test, func(t *testing.T) {
			rndNums := []int{tt.a, tt.b, tt.c}
			got := containsDuplicate(rndNums)
			if got != tt.want {
				t.Errorf("got %t, want %t", got, tt.want)
			}
		})
	}
}

func Test_equalDistance(t *testing.T) {
	var tests = []struct {
		a, b, c int
		want    bool
	}{
		{10, 5, 0, true},
		{10, 0, 5, true},
		{0, 10, 5, true},
		{10, 5, 1, false},
	}

	for _, tt := range tests {
		test := fmt.Sprintf("%d, %d, %d, %t", tt.a, tt.b, tt.c, tt.want)
		t.Run(test, func(t *testing.T) {
			rndNums := []int{tt.a, tt.b, tt.c}
			got := equalDistance(rndNums)
			if got != tt.want {
				t.Errorf("got %t, want %t", got, tt.want)
			}
		})
	}
}

func Test_rndNumAry(t *testing.T) {
	want := 20
	got := rndNumAry(want)
	for x, tt := range got {
		if want <= tt {
			t.Errorf("rndNumAry(20)[%d] = %d, want < %d", x, got, want)
		}
	}
	if len(got) != 3 {
		t.Errorf("rndNumAry(20) length != 3")
	}
}

func Test_correctAnswer(t *testing.T) {
	var tests = []struct {
		a, b, c int
		want    int
	}{
		{10, 4, 0, 10},
		{10, 5, 1, 10},
		{0, 10, 6, 0},
		{1, 4, 10, 10},
	}

	for _, tt := range tests {
		test := fmt.Sprintf("%d, %d, %d, %d", tt.a, tt.b, tt.c, tt.want)
		t.Run(test, func(t *testing.T) {
			rndNums := []int{tt.a, tt.b, tt.c}
			got := correctAnswer(rndNums)
			if got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}
