package bithooks

import (
	"fmt"
	"testing"

	"github.com/kr/pretty"
	"github.com/stretchr/testify/assert"
)

type testcaseDecode struct {
	raw     string
	hooks   Hooks
	errline int
	errtext string
}

func TestDecode(t *testing.T) {
	test := assert.New(t)

	testcases := []testcaseDecode{
		{
			raw: `
1@1
`,
			hooks: Hooks{
				{
					"1", "1", nil,
				},
			},
		},
		{
			raw: `2@2
`,
			hooks: Hooks{
				{
					"2", "2", nil,
				},
			},
		},
		{
			raw: `
3@3
`,
			hooks: Hooks{
				{
					"3", "3", nil,
				},
			},
		},
		// :spaces
		{
			raw: "4_1@4_1\n4_2@4_2",
			hooks: Hooks{
				{
					"4_1", "4_1", nil,
				},
				{
					"4_2", "4_2", nil,
				},
			},
		},
		{
			raw: "\n5_1@5_1\n5_2@5_2",
			hooks: Hooks{
				{
					"5_1", "5_1", nil,
				},
				{
					"5_2", "5_2", nil,
				},
			},
		},
		{
			raw: "6_1@6_1\n6_2@6_2\n",
			hooks: Hooks{
				{
					"6_1", "6_1", nil,
				},
				{
					"6_2", "6_2", nil,
				},
			},
		},
		{
			raw: "\n\n7_1@7_1\n\n\n7_2@7_2\n\n",
			hooks: Hooks{
				{
					"7_1", "7_1", nil,
				},
				{
					"7_2", "7_2", nil,
				},
			},
		},
		{
			raw: `
8_1@8_1
`,
			hooks: Hooks{
				{
					"8_1", "8_1", nil,
				},
			},
		},
		{
			raw: `
9_1@9_1
 args 9_1 9_1
 argsblah 9_1 9_1

9_2@9_2
`,
			hooks: Hooks{
				{
					"9_1", "9_1",
					[]string{"args 9_1 9_1", "argsblah 9_1 9_1"},
				},
				{
					"9_2", "9_2", nil,
				},
			},
		},
		{
			raw: `
10_1@10_1
 args 10_1 10_1
 argsblah 10_1 10_1

10_2@10_2
 xxxx10_2
 yyyyy10_2
 ` + `
 zzzzy10_2`,
			hooks: Hooks{
				{
					"10_1", "10_1",
					[]string{"args 10_1 10_1", "argsblah 10_1 10_1"},
				},
				{
					"10_2", "10_2",
					[]string{"xxxx10_2", "yyyyy10_2", "", "zzzzy10_2"},
				},
			},
		},
		// :errors
		// :redefine
		{
			raw: `
X@X
X@Y
`,
			hooks: Hooks{
				{"X", "X", nil},
				{"X", "Y", nil},
			},
		},
		{
			raw: `
W@W
E@W
`,
			hooks: Hooks{
				{"W", "W", nil},
				{"E", "W", nil},
			},
		},
		{
			raw: `
Q@Q
Q@Q
`,
			errline: 3,
			errtext: errSyntaxRedefine,
		},
		{
			raw: `
T@T
 args1

T@T
`,
			errline: 5,
			errtext: errSyntaxRedefine,
		},
		{
			raw: `
S@S
 args1
args_without_space

N@N
`,
			errline: 4,
			errtext: errSyntaxDefine,
		},
		{
			raw: `
h@f
 jjj
 kkk

 l
`,
			errline: 6,
			errtext: errSyntaxUnexpectedHookArgs,
		},
	}

	for _, testcase := range testcases {
		testcaseIdentifier := fmt.Sprintf(
			"\ntestcase:\n%# v\n",
			pretty.Formatter(testcase),
		)

		actualHooks, err := Decode(testcase.raw)
		if err != nil {
			if testcase.errline == 0 {
				test.NoError(err, testcaseIdentifier)
			}

			if test.IsType(syntaxError{}, err, testcaseIdentifier) {
				if test.Equal(
					testcase.errtext,
					err.(syntaxError).text,
					testcaseIdentifier,
				) {
					test.Equal(
						testcase.errline,
						err.(syntaxError).line,
						testcaseIdentifier,
					)
				}
			}

			continue
		} else {
			if testcase.errtext != "" && !test.Error(err, testcaseIdentifier) {
				continue
			}
		}

		if test.Len(actualHooks, len(testcase.hooks), testcaseIdentifier) {
			for index, expectedHook := range testcase.hooks {
				testcaseHookIdentifier := fmt.Sprintf(
					"%s\nindex: %d\nexpected hook: %# v",
					testcaseIdentifier, index, pretty.Formatter(expectedHook),
				)
				actualHook := actualHooks[index]

				test.Equal(
					expectedHook.Name, actualHook.Name, testcaseHookIdentifier,
				)
				test.Equal(
					expectedHook.ID, actualHook.ID, testcaseHookIdentifier,
				)
				test.Equal(
					expectedHook.ID, actualHook.ID, testcaseHookIdentifier,
				)
				test.Equal(
					expectedHook.Args, actualHook.Args, testcaseHookIdentifier,
				)
			}
		}
	}
}
