package cleaner

import "testing"

func TestClean(t *testing.T) {
	// Setup some test cases with inputs and expected results.
	cases := []struct {
		input  string
		output string
		err    error
	}{
		{"Hello", "hello", nil},
		{"F4T417Y", "f4t417y", nil},
		{"   ", "", ErrAllNonAlphanums},
		{"trailingspace ", "trailingspace", nil},
		{" leadingspace", "leadingspace", nil},
		{"1234567890", "1234567890", ErrAllDigits},
		{"1000.01", "100001", ErrAllDigits},
		{"100,001", "100001", ErrAllDigits},
	}

	for _, c := range cases {
		output, err := Clean(c.input)
		if output != c.output || err != c.err {
			t.Errorf("Clean(%q) == %q with %q\n\tExpected %q with %q",
				c.input, output, err, c.output, c.err)
		}
	}
}
