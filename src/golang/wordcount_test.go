package main

import (
	"bufio"
	"strings"
	"testing"
)

func TestOpenIfNot(t *testing.T) {
	// Setup some test cases with inputs and expected results.
	cases := []struct {
		input1    string
		input2    string
		expectNil bool
	}{
		{"wordcount.go", ".", false},
		{".", ".", true},
		{".", "wordcount.go", false},
	}

	for _, c := range cases {
		output := openIfNot(&c.input1, c.input2)
		if output != nil {
			defer output.Close()
			if c.expectNil {
				t.Errorf("openIfNot(%q, %q) == \"valid *os.File\"\n\tExpected \"nil\"",
					c.input1, c.input2)
			}
		} else {
			if !c.expectNil {
				t.Errorf("openIfNot(%q, %q) == \"nil\"\n\tExpected \"valid *os.File\"",
					c.input1, c.input2)
			}
		}
	}
}

func TestReadAndCount(t *testing.T) {
	// Setup some test cases with inputs and expected results.
	cases := []struct {
		input  string
		expect map[string]int64
	}{
		{"", map[string]int64{}},
		{"000\n", map[string]int64{}},
		{"a\n", map[string]int64{"a": 1}},
		{"a b \n b 5 a 1\n", map[string]int64{"a": 2, "b": 2}},
		{"a b\n c c\n100.00\n", map[string]int64{"a": 1, "b": 1, "c": 2}},
	}

	for _, c := range cases {
		output := readAndCount(bufio.NewReader(strings.NewReader(c.input)))
		for word, count := range output {
			if exCount, has := c.expect[word]; has && exCount != count {
				t.Errorf("readAndCount(%q) counted %q %d times expected %d times",
					c.input, word, count, exCount)
			}
		}
		for exWord, exCount := range c.expect {
			if _, has := output[exWord]; !has {
				t.Errorf("readAndCount(%q) did not count %q expected %d times count",
					c.input, exWord, exCount)
			}
		}
	}
}
