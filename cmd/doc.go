package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

var docCmd = &cobra.Command{
	Use:   "doc",
	Short: "Generate documentation",
	Run: func(cmd *cobra.Command, args []string) {
		println("Generating documentation...")
		os.MkdirAll("/tmp/docs", 0755)
		doc.GenMarkdownTree(rootCmd, "/tmp/docs")
	},
}
