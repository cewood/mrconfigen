package internal

import (
	"bytes"
	"fmt"
	"os"
	"testing"
)

func TestGetTemplate(t *testing.T) {
	basepath, _ := os.Getwd()

	var tests = []struct {
		name  string
		input string
		want  string
	}{
		{
			"test #1: default template",
			"",
			`##
# {{ .Name }} repositories
#
{{- $prefix := .Prefix }}
{{- $owner := .Name }}
{{ range .ReposList }}
[{{ $prefix }}/{{ $owner }}/{{ .Name }}]
checkout = git clone '{{ .GitURL }}' '{{ .Name }}'
{{ end }}
`,
		},
		{
			"test #2: custom template",
			fmt.Sprintf("%s/%s", basepath, "internal_test_tpl.txt"),
			`##
# {{ .Name }} repositories - custom template
#
{{- $prefix := .Prefix }}
{{- $owner := .Name }}
{{ range .ReposList }}
[{{ $prefix }}/{{ $owner }}/{{ .Name }}]
checkout = git clone '{{ .GitURL }}' '{{ .Name }}'
{{ end }}
`,
		},
		{
			"test #3: broken custom template",
			fmt.Sprintf("%s/%s", basepath, "internal_test_tpl_broken.txt"),
			`##
# {{ .Name }} repositories
#
{{- $prefix := .Prefix }}
{{- $owner := .Name }}
{{ range .ReposList }}
[{{ $prefix }}/{{ $owner }}/{{ .Name }}]
checkout = git clone '{{ .GitURL }}' '{{ .Name }}'
{{ end }}
`,
		},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("name: '%s'", tt.name)
		t.Run(testname, func(t *testing.T) {
			fmt.Printf("input is: %v\n", tt.input)
			output := getTemplate(tt.input)

			if output != tt.want {
				t.Errorf("got %v, want %v", output, tt.want)
			}
		})
	}
}

func TestRenderTemplate(t *testing.T) {
	var tests = []struct {
		name   string
		owner  string
		prefix string
		repos  []Repo
		want   string
	}{
		{
			"test #1",
			"owner",
			"prefix",
			[]Repo{{Name: "repo", GitURL: "git@gnowhere:owner/repo.git"}},
			`##
# owner repositories
#

[prefix/owner/repo]
checkout = git clone 'git@gnowhere:owner/repo.git' 'repo'

`,
		},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("owner: '%s', name: '%s', prefix: '%s', repos: '%v'", tt.owner, tt.name, tt.prefix, tt.repos)
		t.Run(testname, func(t *testing.T) {
			buf := new(bytes.Buffer)

			RenderTemplate(tt.prefix, tt.owner, tt.repos, DefaultTemplate, buf)

			if buf.String() != tt.want {
				t.Errorf("got %v, want %v", buf.String(), tt.want)
			}
		})
	}
}
