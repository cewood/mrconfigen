package github

import (
	"context"
	"os"

	"github.com/cewood/mrconfigen/internal"
	"github.com/google/go-github/v32/github"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
)

// Run is responsible for all control, error handling, configuration, everything
func Run() {
	optsUser := &github.RepositoryListOptions{
		Affiliation: viper.GetString("affiliation"),
		Direction:   viper.GetString("direction"),
		ListOptions: github.ListOptions{PerPage: viper.GetInt("count")},
		Sort:        viper.GetString("sort"),
		Type:        viper.GetString("type"),
		Visibility:  viper.GetString("visibility"),
	}

	optsOrg := &github.RepositoryListByOrgOptions{
		Direction:   viper.GetString("direction"),
		ListOptions: github.ListOptions{PerPage: viper.GetInt("count")},
		Sort:        viper.GetString("sort"),
		Type:        viper.GetString("type"),
	}

	var queryFunc func(client *github.Client, name string, page int) ([]*github.Repository, *github.Response, error)

	if viper.GetBool("user") {
		queryFunc = userQueryFunc(optsUser)
	} else {
		queryFunc = orgQueryFunc(optsOrg)
	}

	repos := queryRepos(
		newClient(viper.GetString("token")),
		viper.GetString("prefix"),
		viper.GetString("name"),
		queryFunc,
	)

	internal.RenderTemplate(
		viper.GetString("prefix"),
		viper.GetString("name"),
		convertRepos(repos),
		viper.GetString("template"),
		os.Stdout,
	)
}

// newClient creates a new github client using oauth2 for authentication
func newClient(token string) *github.Client {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(context.Background(), ts)

	return github.NewClient(tc)
}

// orgQueryFunc is responsible for determining the query type and returning a func to perform that query
func orgQueryFunc(opts *github.RepositoryListByOrgOptions) func(client *github.Client, name string, page int) ([]*github.Repository, *github.Response, error) {
	return func(client *github.Client, name string, page int) ([]*github.Repository, *github.Response, error) {
		opts.Page = page

		return client.Repositories.ListByOrg(context.Background(), name, opts)
	}
}

// userQueryFunc is responsible for determining the query type and returning a func to perform that query
func userQueryFunc(opts *github.RepositoryListOptions) func(client *github.Client, name string, page int) ([]*github.Repository, *github.Response, error) {
	return func(client *github.Client, name string, page int) ([]*github.Repository, *github.Response, error) {
		opts.Page = page

		return client.Repositories.List(context.Background(), name, opts)
	}
}

// queryRepos is responsible for determining the query type and performing the query
func queryRepos(client *github.Client, prefix string, name string, queryFunc func(client *github.Client, name string, page int) ([]*github.Repository, *github.Response, error)) []*github.Repository {
	var allRepos []*github.Repository

	var page int

	for {
		repos, resp, err := queryFunc(client, name, page)

		log.WithFields(log.Fields{
			"repos": repos,
			"resp":  resp,
			"err":   err,
		}).Trace("queryFunc outputs")

		if err != nil {
			log.WithFields(log.Fields{
				"err": err,
			}).Panic("encountered error during request")
		}

		allRepos = append(allRepos, repos...)

		if resp.NextPage == 0 {
			break
		}

		page = resp.NextPage
	}

	return allRepos
}

// convertRepos converts a []*github.Repository into a []internal.Repo which is
//  what the internal.RenderTemplate func accepts as input
func convertRepos(repos []*github.Repository) []internal.Repo {
	newRepos := make([]internal.Repo, 0)

	for _, v := range repos {
		newRepos = append(newRepos,
			internal.Repo{
				Name:   *v.Name,
				GitURL: *v.SSHURL,
			})
	}

	return newRepos
}
