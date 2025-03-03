package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	// cfgFile      string
	// userLicense  string
	ghInstallCmd = &cobra.Command{
		Use:   "gh-install",
		Short: "Install binaries from GitHub releases",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Installing gh")
		},
	}
)

func init() {

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cobra.yaml)")
	// rootCmd.PersistentFlags().StringP("author", "a", "YOUR NAME", "author name for copyright attribution")
	// rootCmd.PersistentFlags().StringVarP(&userLicense, "license", "l", "", "name of license for the project")
	// rootCmd.PersistentFlags().Bool("viper", true, "use Viper for configuration")

}
