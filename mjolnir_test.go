package main

import (
	"math"
	"reflect"
	"testing"
)

func Test_parseIssueFixes(t *testing.T) {
	testCases := []struct {
		name            string
		text            string
		expectedNumbers []int
	}{
		{
			name: "only letters",
			text: `
	Fixes dlsqj
`,
			expectedNumbers: []int{},
		},
		{
			name: "valid issue numbers coma",
			text: `
	Fixes #13 #14, #15,#16,
`,
			expectedNumbers: []int{13, 14, 15, 16},
		},
		{
			name: "valid issue numbers space",
			text: `
	Fixes #13 #14 #15 #16
`,
			expectedNumbers: []int{13, 14, 15, 16},
		},
		{
			name: "invalid pattern",
			text: `
	Fixes #13#14,#15,#16,
`,
			expectedNumbers: []int{},
		},
		{
			name: "french style",
			text: `
	Fixes : #13,#14,#15,#16,
`,
			expectedNumbers: []int{13, 14, 15, 16},
		},
		{
			name: "valid issue numbers coma and :",
			text: `
	Fixes: #13,#14,#15,#16,
`,
			expectedNumbers: []int{13, 14, 15, 16},
		},
		{
			name: "multiple lines:",
			text: `
	Fixes: #13,#14
	Fixes: #15,#16
`,
			expectedNumbers: []int{13, 14, 15, 16},
		},
		{
			name: "max int 64",
			text: `
	Fixes: #9223372036854775807
`,
			expectedNumbers: []int{math.MaxInt64},
		},
		{
			name: "valid issue numbers ends with a dot",
			text: `
	Fixes #13 #14, #15,#16.
`,
			expectedNumbers: []int{13, 14, 15, 16},
		},
		{
			name: "multiple lines end with a dot",
			text: `
	Fixes: #13,#14.
	Fixes: #15,#16.
`,
			expectedNumbers: []int{13, 14, 15, 16},
		},
	}

	for _, test := range testCases {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			issueNumbers := parseIssueFixes(test.text)

			if (len(issueNumbers) != 0 || len(test.expectedNumbers) != 0) && !reflect.DeepEqual(issueNumbers, test.expectedNumbers) {
				t.Errorf("Got %v, expected %v", issueNumbers, test.expectedNumbers)
			}
		})
	}
}
