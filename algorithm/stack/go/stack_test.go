package stack

import "testing"

func TestCheckMultiBracket(t *testing.T) {
	tests := []struct {
		In       string
		Expected bool
	}{
		{
			In:       "",
			Expected: false,
		},
		{
			In:       "(",
			Expected: false,
		},
		{
			In:       "()",
			Expected: true,
		},
		{
			In:       "(())",
			Expected: true,
		},
		{
			In:       "(([[]]))",
			Expected: true,
		},
		{
			In:       "(([[))]]",
			Expected: false,
		},
		{
			In:       "(([[{{}}]]))",
			Expected: true,
		},
	}

	for i, test := range tests {
		res := CheckMultiBracket(test.In)
		if res != test.Expected {
			t.Fatalf("CheckMultiBracket error%d %t want %t", i, res, test.Expected)
		}
	}
}
