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
			name: "not pr repo",
			text: `
			Fixes https://github.com/your/repo/issues/42
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
			name: "valid issue urls coma",
			text: `
	Fixes https://github.com/my/re-po/issues/13 #14, https://github.com/my/re-po/issues/15,https://github.com/my/re-po/issues/16,
`,
			expectedNumbers: []int{13, 14, 15, 16},
		},
		{
			name: "valid issue urls space",
			text: `
			Fixes #13 #14 #15 #16
		`,
			expectedNumbers: []int{13, 14, 15, 16},
		},
		{
			name: "valid issue mixed numbers/urls space",
			text: `
			Fixes https://github.com/my/re-po/issues/13 #14 https://github.com/my/re-po/issues/15 #16 https://github.com/your/re-po/issues/17
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
			name: "invalid pattern urls",
			text: `
			Fixes https://github.com/my/re-po/issues/13https://github.com/my/re-po/issues/14,https://github.com/my/re-po/issues/15,https://github.com/my/re-po/issues/16,
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
			name: "french style with urls",
			text: `
			Fixes : https://github.com/my/re-po/issues/13,https://github.com/my/re-po/issues/14,https://github.com/my/re-po/issues/15,https://github.com/my/re-po/issues/16,
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
			name: "valid issue urls coma and :",
			text: `
			Fixes: https://github.com/my/re-po/issues/13,https://github.com/my/re-po/issues/14,https://github.com/my/re-po/issues/15,https://github.com/my/re-po/issues/16,
		`,
			expectedNumbers: []int{13, 14, 15, 16},
		},
		{
			name: "multiple lines:",
			text: `
			Fixes: #13,#14
			Fixes: #15,#16
			Fixes: https://github.com/your/repo/issues/17
		`,
			expectedNumbers: []int{13, 14, 15, 16},
		},
		{
			name: "multiple lines with urls:",
			text: `
			Fixes: https://github.com/my/re-po/issues/13,https://github.com/my/re-po/issues/14
			Fixes: https://github.com/my/re-po/issues/15,https://github.com/my/re-po/issues/16
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
			name: "max int 64 with urls",
			text: `
			Fixes: https://github.com/my/re-po/issues/9223372036854775807
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
			name: "valid issue urls ends with a dot",
			text: `
			Fixes https://github.com/my/re-po/issues/13 https://github.com/my/re-po/issues/14, https://github.com/my/re-po/issues/15,https://github.com/my/re-po/issues/16.
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
		{
			name: "multiple line urls end with a dot",
			text: `
			Fixes: https://github.com/my/re-po/issues/13,https://github.com/my/re-po/issues/14.
			Fixes: https://github.com/my/re-po/issues/15,https://github.com/my/re-po/issues/16.
		`,
			expectedNumbers: []int{13, 14, 15, 16},
		},
	}

	for _, test := range testCases {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			issueNumbers := parseIssueFixes(test.text, "my", "re-po")

			if (len(issueNumbers) != 0 || len(test.expectedNumbers) != 0) && !reflect.DeepEqual(issueNumbers, test.expectedNumbers) {
				t.Errorf("Got %v, expected %v", issueNumbers, test.expectedNumbers)
			}
		})
	}
}
