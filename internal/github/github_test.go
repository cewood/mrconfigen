package github

import (
	"fmt"
	"testing"

	"github.com/cewood/mrconfigen/internal"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-github/v32/github"
)

func TestConvertRepos(t *testing.T) {
	var name = "input-name"
	var url = "input-url"

	var tt = struct {
		input []*github.Repository
		want  []internal.Repo
	}{
		[]*github.Repository{{Name: &name, SSHURL: &url}},
		[]internal.Repo{{Name: name, GitURL: url}},
	}

	testname := fmt.Sprintf("input: '%s', want: '%s'", tt.input, tt.want)
	t.Run(testname, func(t *testing.T) {
		ans := convertRepos(tt.input)
		if !cmp.Equal(ans, tt.want) {
			t.Errorf("got '%s', want '%s'", ans, tt.want)
		}
	})
}

func TestNewClient(t *testing.T) {
	var tt = struct {
		input string
		want  *github.Client
	}{
		"sometoken",
		&github.Client{},
	}

	testname := fmt.Sprintf("input: '%v', want: '%v'", tt.input, tt.want)
	t.Run(testname, func(t *testing.T) {
		ans := newClient(tt.input)
		if fmt.Sprintf("%T", ans) != "*github.Client" {
			t.Errorf("got '%v', want '%v'", ans, tt.want)
		}
	})
}

func TestOrgQueryFunc(t *testing.T) {
	var tt = struct {
		input *github.RepositoryListByOrgOptions
		want  string
	}{
		&github.RepositoryListByOrgOptions{},
		"func(*github.Client, string, int) ([]*github.Repository, *github.Response, error)",
	}

	testname := fmt.Sprintf("input: '%v'", tt.input)
	t.Run(testname, func(t *testing.T) {
		ans := orgQueryFunc(tt.input)
		if fmt.Sprintf("%T", ans) != tt.want {
			t.Errorf("got '%T', want '%v'", ans, tt.want)
		}
	})
}

func TestUserQueryFunc(t *testing.T) {
	var tt = struct {
		input *github.RepositoryListOptions
		want  string
	}{
		&github.RepositoryListOptions{},
		"func(*github.Client, string, int) ([]*github.Repository, *github.Response, error)",
	}

	testname := fmt.Sprintf("input: '%v'", tt.input)
	t.Run(testname, func(t *testing.T) {
		ans := userQueryFunc(tt.input)
		if fmt.Sprintf("%T", ans) != tt.want {
			t.Errorf("got '%T', want '%v'", ans, tt.want)
		}
	})
}
