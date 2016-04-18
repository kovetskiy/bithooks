package bitmultihooks

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
				{"x", "y", ""},
			},
			output: "x@y\n",
		},
		{
			hooks: Hooks{
				{"1", "1", ""},
				{"2", "2", ""},
				{"3", "3", ""},
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
				{"q", "w", "data"},
			},
			output: "" +
				"q@w\n" +
				" data\n",
		},
		{
			hooks: Hooks{
				{"y", "u", "data1 \n data2 "},
			},
			output: "" +
				"y@u\n" +
				" data1 \n" +
				"  data2 \n",
		},
		{
			hooks: Hooks{
				{"a", "s", "data_as\ndataaaaaaaaaa"},
				{"z", "c", "data_zc"},
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
