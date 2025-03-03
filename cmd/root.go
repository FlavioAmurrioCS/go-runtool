package cmd

import (
	"fmt"
	"os"

	"github.com/FlavioAmurrioCS/go-runtool/cmd/sub1"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   os.Args[0],
	Short: "Hugo is a very fast static site generator",
	Long: `A Fast and Flexible Static Site Generator built with
	love by spf13 and friends in Go.
	Complete documentation is available at https://gohugo.io/documentation/`,
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println("Running rootCmd")
		fmt.Println(os.Args)
	},
	Version: "v0.0.1",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(scrapeCmd)
	rootCmd.AddCommand(filterCmd)
	rootCmd.AddCommand(linkInstallCmd)
	rootCmd.AddCommand(ghInstallCmd)
	rootCmd.AddCommand(runCmd)

	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(sub1.Sub1Cmd)
	rootCmd.AddCommand(docCmd)
}
