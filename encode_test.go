package bithooks

import (
	"fmt"
	"testing"

	"github.com/kr/pretty"
	"github.com/stretchr/testify/assert"
)

type testcaseEncode struct {
	hooks  Hooks
	output string
}

func TestEncode(t *testing.T) {
	test := assert.New(t)

	testcases := []testcaseEncode{
		{
			hooks: Hooks{
				{"x", "y", []string{}},
			},
			output: "x@y\n",
		},
		{
			hooks: Hooks{
				{"1", "1", []string{}},
				{"2", "2", []string{}},
				{"3", "3", []string{}},
			},
			output: "" +
				"1@1\n" +
				"\n" +
				"2@2\n" +
				"\n" +
				"3@3\n",
		},
		{
			hooks: Hooks{
				{"q", "w", []string{"data"}},
			},
			output: "" +
				"q@w\n" +
				" data\n",
		},
		{
			hooks: Hooks{
				{"y", "u", []string{"data1 ", " data2 "}},
			},
			output: "" +
				"y@u\n" +
				" data1 \n" +
				"  data2 \n",
		},
		{
			hooks: Hooks{
				{"a", "s", []string{"data_as", "dataaaaaaaaaa"}},
				{"z", "c", []string{"data_zc"}},
			},
			output: "" +
				"a@s\n" +
				" data_as\n" +
				" dataaaaaaaaaa\n" +
				"\n" +
				"z@c\n" +
				" data_zc\n",
		},
	}

	for _, testcase := range testcases {
		testcaseIdentifier := fmt.Sprintf(
			"\ntestcase:\n%# v\n",
			pretty.Formatter(testcase),
		)

		test.Equal(
			testcase.output, Encode(testcase.hooks), testcaseIdentifier,
		)
	}
}
