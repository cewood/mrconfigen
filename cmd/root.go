package cmd

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "mrconfigen",
	Short: "A tool to generate myrepos configuration from various SCM systems",
	Long:  `Generate a myrepos configuration file from your SCM system of choice.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		viper.BindPFlag("debug", cmd.Flags().Lookup("debug"))
		viper.BindPFlag("verbose", cmd.Flags().Lookup("verbose"))
		viper.BindPFlag("template", cmd.Flags().Lookup("template"))
		viper.BindPFlag("prefix", cmd.Flags().Lookup("prefix"))

		// Output to stdout instead of the default stderr
		log.SetOutput(os.Stdout)

		log.SetReportCaller(true)

		if viper.GetBool("debug") {
			log.SetLevel(log.TraceLevel)
		} else if viper.GetBool("verbose") {
			log.SetLevel(log.DebugLevel)
		} else {
			log.SetLevel(log.WarnLevel)
		}
	},
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	// Default for prefix, if not explicitly set
	basepath, _ := os.Getwd()

	// Define our flags and configuration settings.
	rootCmd.PersistentFlags().BoolP("debug", "d", false, "enable debug output, defaults to false")
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "enable verbose output, defaults to false")
	rootCmd.PersistentFlags().StringP("prefix", "p", basepath, "base path to use in the mrconfig, is CWD if not specified")
	rootCmd.PersistentFlags().StringP("template", "t", "", "provide the path to a custom template, to override the default inbuilt template")
}
