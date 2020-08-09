# mrconfigen

A small utility to generate [myrepos](https://myrepos.branchable.com/) config files from the repostories in your SCM system of choice.


## Example usage

Using Mrconfigen to query a GitHub Organisation:

```shell
$ mrconfigen github --token 1234567890 --prefix /home/jdoe/code/src/github.com --name acmecorp
##
# acmecorp repositories
#

[/home/jdoe/code/src/github.com/acmecorp/front-end]
checkout = git clone 'git@github.com:acmecorp/front-end.git' 'front-end'

[/home/jdoe/code/src/github.com/acmecorp/payments]
checkout = git clone 'git@github.com:acmecorp/payments.git' 'payments'

[/home/jdoe/code/src/github.com/acmecorp/shipping]
checkout = git clone 'git@github.com:acmecorp/shipping.git' 'shipping'

```


## Provider specific options
### GitHub

```
Query the GitHub repositories of an Org/User and generate mr config for them

Usage:
  mrconfigen github [flags]

Flags:
      --affiliation string   list repos of given affiliation[s]. comma-separated list, can include: owner, collaborator, organization_member. only used in user mode (default "owner,collaborator,organization_member")
  -c, --count int            the number of records to fetch per-request/pagination (default 100)
      --direction string     direction in which to sort repositories. can be one of asc or desc (default "asc")
  -h, --help                 help for github
  -n, --name string          the org or user to query
      --sort string          how to sort the repository list. can be one of created, updated, pushed, full_name (default "full_name")
      --token string         the token to use for api requests
      --type string          type of repositories to list. can be one of: all, public, private, forks, sources, member (default "all")
  -u, --user                 enable user query mode. default is org query
      --visibility string    visibility of repositories to list. can be one of all, public, or private. this option is only used in user mode (default "all")

Global Flags:
  -d, --debug             enable debug output, defaults to false
  -p, --prefix string     base path to use in the mrconfig, is CWD if not specified (default "/home/cewood/code/sync/src/github.com/cewood/mrconfigen")
  -t, --template string   provide the path to a custom template, to override the default inbuilt template
  -v, --verbose           enable verbose output, defaults to false
```


## Frequently asked questions
### Why would I want to use Myrepos to begin with

Great question, it's a nice tool for dealing with many git repositories en masse. Whether you deal with a lot of repositories at your work and want to update them all every morning in one simple command. Or your company has a lot of repositories that you want to keep up with, as well as discover new ones, which can be tricky/time intensive to do manually. Mrconfigen makes this effortless.
