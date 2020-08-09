package cmd

import (
	"log"

	"github.com/cewood/mrconfigen/internal/github"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// githubCmd represents the github command
var githubCmd = &cobra.Command{
	Use:   "github",
	Short: "GitHub provider for Mrconfigen",
	Long:  `Query the GitHub repositories of an Org/User and generate mr config for them`,
	Run: func(cmd *cobra.Command, args []string) {
		// Explicitly check our required args, since the inbuilt required
		//  annotation doesn't take into account environment variables
		if viper.GetString("token") == "" {
			log.Fatal("token was not set, please supply a valid token")
		} else if viper.GetString("name") == "" {
			log.Fatal("name was not set, please supply a valid name")
		}

		github.Run()
	},
	PreRun: func(cmd *cobra.Command, args []string) {
		// Explicitly check our required args, since the inbuilt required
		//  annotation doesn't take into account environment variables
		if viper.GetString("token") == "" {
			log.Fatal("token was not set, please supply a valid token")
		} else if viper.GetString("name") == "" {
			log.Fatal("name was not set, please supply a valid name")
		}
	},
}

func init() {
	rootCmd.AddCommand(githubCmd)

	// Define our flags and configuration settings.
	githubCmd.Flags().BoolP("user", "u", false, "enable user query mode. default is org query")
	githubCmd.Flags().IntP("count", "c", 100, "the number of records to fetch per-request/pagination")
	githubCmd.Flags().String("affiliation", "owner,collaborator,organization_member", "list repos of given affiliation[s]. comma-separated list, can include: owner, collaborator, organization_member. only used in user mode")
	githubCmd.Flags().String("direction", "asc", "direction in which to sort repositories. can be one of asc or desc")
	githubCmd.Flags().String("sort", "full_name", "how to sort the repository list. can be one of created, updated, pushed, full_name")
	githubCmd.Flags().String("type", "all", "type of repositories to list. can be one of: all, public, private, forks, sources, member")
	githubCmd.Flags().String("visibility", "all", "visibility of repositories to list. can be one of all, public, or private. this option is only used in user mode")
	githubCmd.Flags().StringP("name", "n", "", "the org or user to query")
	githubCmd.Flags().String("token", "", "the token to use for api requests")

	// Bind flags to viper config, for free environment lookups
	for _, val := range []string{
		"affiliation",
		"count",
		"direction",
		"name",
		"prefix",
		"sort",
		"token",
		"type",
		"user",
		"visibility",
	} {
		if err := viper.BindPFlag(val, githubCmd.Flags().Lookup(val)); err != nil {
			log.Fatalf("unable to bind %s flag: '%v'", val, err)
		}
	}
}
