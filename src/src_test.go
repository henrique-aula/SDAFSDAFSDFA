package src

import (
	"testing"
)

func Test_Soma(t *testing.T) {
	cases := []struct {
		a, b, expected int
	}{
		{1, 2, 3},
		{4, 5, 9},
	}

	for _, c := range cases {
		res := Soma(c.a, c.b)
		expected := c.expected
		if res != expected {
			t.Errorf("Soma retornou %v, experado : %v\n", res, expected)
		}
	}

}
