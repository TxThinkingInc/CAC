package cac

import (
	"strings"
	"testing"
)

func TestParse(t *testing.T) {
	tt := []struct {
		name    string
		content string
		pass    func(list []string) bool
	}{
		{
			name:    "Just space",
			content: "command1 command2 command3",
			pass: func(list []string) bool {
				return len(list) == 3
			},
		},
		{
			name: "Just line",
			content: `
Monday
Tuesday

Thursday
Fri\
day`,
			pass: func(list []string) bool {
				return len(list) == 4
			},
		},
		{
			name: "With command",
			content: `
frog
			# Frog also called "the chicken in the filed" in Chinese
# beacuse it taste like chicken.
chicken`,
			pass: func(list []string) bool {
				return len(list) == 2
			},
		},
		{
			name: "Mix",
			content: `
			# some command
			fun today
			rig\
			ht

			lsll`,
			pass: func(list []string) bool {
				return len(list) == 4
			},
		},
	}

	for _, test := range tt {
		result := Parse(test.content)
		if !test.pass(result) {
			t.Errorf("Failed to parse %s, got: %v", test.name, strings.Join(result, " "))
			continue
		}
		t.Logf("%s passed", test.name)
	}
}
