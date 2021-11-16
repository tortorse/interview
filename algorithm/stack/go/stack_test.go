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

func TestNavigator(t *testing.T) {
	navigator := NewNavigator()
	url1 := "https://idg-daily-report.vercel.app/"
	navigator.Open(url1)

	if navigator.Url != url1 {
		t.Fatalf("Navigator error1")
	}

	url2 := "https://idg-daily-report.vercel.app/reports"
	navigator.Open(url2)

	if navigator.Url != url2 {
		t.Fatalf("Navigator error2")
	}

	url3 := "https://idg-daily-report.vercel.app/exports"
	navigator.Open(url3)

	if navigator.Url != url3 {
		t.Fatalf("Navigator error3")
	}

	navigator.Back()

	if navigator.Url != url2 {
		t.Fatalf("Navigator error4")
	}

	navigator.Back()

	if navigator.Url != url1 {
		t.Fatalf("Navigator error5")
	}

	navigator.Forward()

	if navigator.Url != url2 {
		t.Fatalf("Navigator error6")
	}

	navigator.Back()
	if navigator.Url != url1 {
		t.Fatalf("Navigator error7")
	}

	navigator.Forward()
	navigator.Forward()
	if navigator.Url != url3 {
		t.Fatalf("Navigator error8")
	}
}
